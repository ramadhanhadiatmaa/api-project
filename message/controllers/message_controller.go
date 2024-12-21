package controllers

import (
	"time"

	"message/models"

	"github.com/gofiber/fiber/v2"
)

func CreateMessage(c *fiber.Ctx) error {
	var data models.Message

	if err := c.BodyParser(&data); err != nil {
		return jsonResponse(c, fiber.StatusBadRequest, "Invalid input", err.Error())
	}

	var conversation models.Conversation
	err := models.DB.Where("cust_user = ? AND seller_user = ?", data.Sender, data.Receiver).First(&conversation).Error

	if err != nil {

		conversation = models.Conversation{
			CustUser:   data.Sender,
			SellerUser: data.Receiver,
			LastId:     0,
			CreatedAt:  time.Now(),
		}

		if err := models.DB.Create(&conversation).Error; err != nil {
			return jsonResponse(c, fiber.StatusInternalServerError, "Failed to create conversation", err.Error())
		}
	}

	data.ConversationID = conversation.ID

	if data.Sender != conversation.CustUser && data.Sender != conversation.SellerUser {
		return jsonResponse(c, fiber.StatusForbidden, "Sender is not part of the conversation", nil)
	}

	// Set additional fields for the message
	data.CreatedAt = time.Now()
	data.IsRead = false

	// Save the new message to the database
	if err := models.DB.Create(&data).Error; err != nil {
		return jsonResponse(c, fiber.StatusInternalServerError, "Failed to save message", err.Error())
	}

	// Update the conversation's last_id field
	if err := models.DB.Model(&conversation).Update("last_id", data.ID).Error; err != nil {
		return jsonResponse(c, fiber.StatusInternalServerError, "Failed to update conversation", err.Error())
	}

	return jsonResponse(c, fiber.StatusCreated, "Message successfully added", data)
}

// jsonResponse is a helper function to send a JSON response.
func jsonResponse(c *fiber.Ctx, status int, message string, data interface{}) error {
	return c.Status(status).JSON(fiber.Map{
		"message": message,
		"data":    data,
	})
}
