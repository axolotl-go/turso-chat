package http

import (
	"github.com/axolotl-go/turso-chat/internal/messages"
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {

	api := app.Group("/api")

	api.Post("/messages", messages.SendMessage)

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Este!")
	})

}
