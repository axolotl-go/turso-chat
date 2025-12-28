package messages

import (
	"github.com/axolotl-go/turso-chat/internal/db"
	"github.com/gofiber/fiber/v2"
)

func SendMessage(c *fiber.Ctx) error {
	var message Message

	if err := c.BodyParser(&message); err != nil {
		return err
	}

	if err := db.DB.Create(&message).Error; err != nil {
		return err
	}

	return c.JSON(message)
}
