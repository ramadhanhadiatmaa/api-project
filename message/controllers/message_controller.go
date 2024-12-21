package controllers

import (
	"time"

	"message/models"

	"github.com/gofiber/fiber/v2"
)

// CreateMessage handles creating a new message in a conversation.
func CreateMessage(c *fiber.Ctx) error {
	var data models.Message

	// Parse the request body into the Message struct
	if err := c.BodyParser(&data); err != nil {
		return jsonResponse(c, fiber.StatusBadRequest, "Invalid input", err.Error())
	}

	// Validate if the ConversationID exists
	var conversation models.Conversation
	err := models.DB.First(&conversation, data.ConversationID).Error

	if err != nil {
		// If conversation does not exist, create a new conversation
		// Create a new conversation between sender and receiver
		conversation = models.Conversation{
			CustUser:   data.Sender,
			SellerUser: conversation.SellerUser, // assuming Receiver is the seller
			LastId:     0,             // No messages yet
			CreatedAt:  time.Now(),
		}

		// Save the new conversation in the database
		if err := models.DB.Create(&conversation).Error; err != nil {
			return jsonResponse(c, fiber.StatusInternalServerError, "Failed to create conversation", err.Error())
		}
		// Update conversation_id in message
		data.ConversationID = conversation.ID
	}

	// Ensure the sender is part of the conversation (either customer or seller)
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