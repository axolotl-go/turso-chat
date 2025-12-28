package main

import (
	"log"
	"os/user"

	"github.com/axolotl-go/turso-chat/internal/db"
	"github.com/axolotl-go/turso-chat/internal/http"
	"github.com/axolotl-go/turso-chat/internal/messages"
	"github.com/gofiber/fiber/v2"
)

func main() {

	if err := db.DB.AutoMigrate(&user.User{}, &messages.Message{}); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	http.SetupRouter(app)

	app.Listen(":8080")
}
