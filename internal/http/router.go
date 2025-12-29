package http

import (
	"github.com/axolotl-go/turso-chat/internal/messages"
	"github.com/axolotl-go/turso-chat/internal/room"
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {

	api := app.Group("/api")

	api.Get("/messages", messages.GetMessages)
	api.Post("/messages", messages.SendMessage)

	api.Get("/rooms", room.GetRooms)
	api.Post("/rooms", room.CreateRoom)
	api.Get("/rooms/:id/messages", room.GetMessagesByRoom)

}
