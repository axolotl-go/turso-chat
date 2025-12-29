package messages

import (
	"github.com/axolotl-go/turso-chat/internal/db"
	"github.com/axolotl-go/turso-chat/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func SendMessage(c *fiber.Ctx) error {
	var message Message

	if err := c.BodyParser(&message); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	message.Sender = utils.RandomName()

	if err := db.DB.Create(&message).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to send message",
		})
	}

	return c.Status(fiber.StatusOK).JSON(message)
}

func GetMessages(c *fiber.Ctx) error {
	var messages []Message

	if err := db.DB.Find(&messages).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get messages",
		})
	}

	return c.Status(fiber.StatusOK).JSON(messages)
}
