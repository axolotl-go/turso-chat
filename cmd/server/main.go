package main

import (
	"log"

	"github.com/axolotl-go/turso-chat/internal/config"
	"github.com/axolotl-go/turso-chat/internal/db"
	"github.com/axolotl-go/turso-chat/internal/http"
	"github.com/axolotl-go/turso-chat/internal/messages"
	"github.com/axolotl-go/turso-chat/internal/room"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	if err := db.DB.AutoMigrate(&room.Room{}, &messages.Message{}); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Use(cors.New(config.CorsConfig()))
	http.SetupRouter(app)

	app.Listen(":8080")
}
