package privatemessages

import (
	"app/internal/adapters/database"
	"app/internal/adapters/database/repository"
	"app/internal/core/services/private-chats"
	"app/internal/core/services/sessions"

	"context"
	"errors"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type Message struct {
	Type   uint32 `json:"type"`
	Data   string `json:"data"`
	Sender uint32 `json:"sender"`
}

type messageQueue chan Message

type ChatQueues struct {
	User1Queue messageQueue
	User2Queue messageQueue
}

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	queues = sync.Map{} // map[chatID]string -> ChatQueues
)

func keepAlive(ctx context.Context, conn *websocket.Conn, interval time.Duration, cancel context.CancelFunc) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if err := conn.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(time.Second)); err != nil {
				log.Printf("ping error: %v", err)
				cancel()
				return
			}
		}
	}
}

// broadcast messages from the queue to the websocket
func receiveBroadcast(ctx context.Context, conn *websocket.Conn, queue messageQueue, cancel context.CancelFunc) {
	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-queue:
			if err := conn.WriteJSON(&msg); err != nil {
				log.Printf("could not broadcast message: %s", err.Error())
				cancel()
				return
			}
		}
	}
}

func Chat(c echo.Context) error {
	if err := uuid.Validate(c.Param("id")); err != nil {
		return c.String(http.StatusBadRequest, "invalid user id")
	}

	db, err := database.GetDbConnection()
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	privateChatRepo, err := repository.NewPrivateChatRepository(db)
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	PCSVC := privatechats.NewSelectSVC(privateChatRepo)
	PCMSG := PCSVC.Select(c.Request().Context(), c.Param("id"))
	if !PCMSG.Succeed {
		return c.JSONPretty(int(PCMSG.Status), &PCMSG, "  ")
	}
	chatMetadata := PCMSG.Data

	sessionRepo, err := repository.NewSessionRepository(db)
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	WLSVC := sessions.NewWhosLoggedSvc(sessionRepo, c)

	WLMSG := WLSVC.WhosLogged(c.Request().Context())
	if !WLMSG.Succeed {
		return c.JSONPretty(int(WLMSG.Status), &WLMSG, "  ")
	}
	u := WLMSG.Data

	v, _ := queues.LoadOrStore(c.Param("id"), ChatQueues{
		User2Queue: make(messageQueue, 128),
		User1Queue: make(messageQueue, 128),
	})
	chatQueues, ok := v.(ChatQueues)
	if !ok {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	var currentQueue, otherQueue messageQueue

	if u.ID == chatMetadata.User1ID && u.ID == chatMetadata.User2ID {
		currentQueue = chatQueues.User1Queue
		otherQueue = chatQueues.User1Queue
	} else if u.ID == chatMetadata.User1ID {
		currentQueue = chatQueues.User2Queue
		otherQueue = chatQueues.User1Queue
	} else {
		currentQueue = chatQueues.User1Queue
		otherQueue = chatQueues.User2Queue
	}

	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println(err)
		return err
	}
	defer conn.Close()

	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	conn.SetPongHandler(func(appData string) error {
		// apenas atualiza o deadline ao receber pong
		conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	go receiveBroadcast(ctx, conn, currentQueue, cancel)

	go keepAlive(ctx, conn, 30*time.Second, cancel)

	for {
		msg := Message{}
		if err := conn.ReadJSON(&msg); err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) ||
				errors.Is(err, websocket.ErrCloseSent) {
				return nil // normal disconnect
			}
			log.Println("read error:", err)
			return nil
		}
		msg.Sender = u.ID

		if len(strings.TrimSpace(msg.Data)) > 0 {
			if len(strings.TrimSpace(msg.Data)) > 0 {
				currentQueue <- msg
				if currentQueue != otherQueue {
					otherQueue <- msg
				}
			}
		}
	}
}

