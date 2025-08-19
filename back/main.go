package main

import (
	"app/internal/adapters/controllers"
	cmiddleware "app/internal/adapters/controllers/middleware"
	"app/internal/adapters/database"
	groupchat "app/internal/core/models/group-chat"
	groupmessage "app/internal/core/models/group-message"
	groupparticipant "app/internal/core/models/group-participant"
	"app/internal/core/models/private-chat"
	"app/internal/core/models/private-message"
	"app/internal/core/models/session"
	"app/internal/core/models/user"
	userstate "app/internal/core/models/user-state"

	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func onProduction() bool {
	return os.Getenv("ON_PRODUCTION") == "true"
}

func main() {
	if err := godotenv.Load(); err != nil && !onProduction() {
		log.Fatalln("error while loading .env: " + err.Error())
	}

	db, err := database.GetDbConnection()
	if err != nil {
		log.Fatalln(err)
	}

	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Error; err != nil {
		log.Fatalln(err)
	}

	if err := db.AutoMigrate(
		&user.User{},
		&privatechat.PrivateChat{},
		&privatemessage.PrivateMessage{},
		&groupchat.GroupChat{},
		&groupparticipant.GroupParticipant{},
		&groupmessage.GroupMessage{},
		&userstate.UserState{},
		&session.Session{},
	); err != nil {
		log.Fatalln(err)
	}

	app := echo.New()

	app.Use(middleware.Logger())
	app.Use(cmiddleware.Authenticate())

	app.IPExtractor = echo.ExtractIPFromXFFHeader()

	controllers.SetupRoutes(app)

	app.Logger.Fatal(app.Start(net.JoinHostPort(os.Getenv("HOST"), os.Getenv("PORT"))))
}
