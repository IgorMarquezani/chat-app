package privatemessages

import (
	"strings"

	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type Message struct {
	Data string `json:"data"`
}

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func Chat(c echo.Context) error {
	chatID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid user id")
	}

	log.Println(chatID)

	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println(err)
		return err
	}
	defer conn.Close()

	for {
		msg := Message{}
		if err := conn.ReadJSON(&msg); err != nil {
			if errors.Is(err, websocket.ErrCloseSent) {
				return c.String(websocket.CloseNormalClosure, "connection closed")
			}
			log.Println(err)
		}

		if len(strings.TrimSpace(msg.Data)) < 1 {
			continue
		}

		conn.WriteJSON(&msg)
	}
}
