package main

import (
	"app/internal/adapters/controllers"
	"app/internal/adapters/database"
	"app/internal/core/models/user"

	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error while loading .env: " + err.Error())
	}

	db, err := database.GetDbConnection()
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&user.User{})

	app := echo.New()

	controllers.SetupRoutes(app)

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	app.Logger.Fatal(app.Start(net.JoinHostPort(host, port)))
}
