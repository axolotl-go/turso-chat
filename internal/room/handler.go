package room

import (
	"github.com/axolotl-go/turso-chat/internal/db"
	"github.com/axolotl-go/turso-chat/internal/messages"
	"github.com/gofiber/fiber/v2"
)

func CreateRoom(c *fiber.Ctx) error {
	var room Room

	if err := c.BodyParser(&room); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := db.DB.Create(&room).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create room",
		})
	}

	return c.Status(fiber.StatusOK).JSON(room)
}

func GetRooms(c *fiber.Ctx) error {
	var rooms []Room

	if err := db.DB.Find(&rooms).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get rooms",
		})
	}

	return c.Status(fiber.StatusOK).JSON(rooms)
}

func GetRoomById(c *fiber.Ctx) error {
	var room Room
	roomId := c.Params("id")

	if err := db.DB.Where("id = ?", roomId).First(&room).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Room not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(room)
}

func GetMessagesByRoom(c *fiber.Ctx) error {
	var messages []messages.Message
	roomId := c.Params("id") // Cambia esto de "roomName" a "id"

	if err := db.DB.Where("room_id = ?", roomId).Find(&messages).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get messages",
		})
	}

	return c.Status(fiber.StatusOK).JSON(messages)
}
