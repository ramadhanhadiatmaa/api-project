package controllers

import (
	"auth/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ShowUs(c *fiber.Ctx) error {
	var user []models.User

	if err := models.DB.Preload("TypeInfo").Find(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch users",
		})
	}

	return c.JSON(user)
}

func IndexUs(c *fiber.Ctx) error {
	id := c.Params("username")
	var user models.User

	// Menggunakan Preload untuk memuat data relasi Type
	if err := models.DB.Preload("Type").First(&user, "username = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "No data found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to load data",
			"error":   err.Error(),
		})
	}

	// Format data response sesuai kebutuhan
	response := map[string]interface{}{
		"username": user.Username,
		"password": user.Password,
		"type":     user.TypeInfo,
	}

	return c.JSON(response)
}

func CreateUs(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return jsonResponse(c, fiber.StatusBadRequest, "Invalid input", err.Error())
	}

	if err := models.DB.Create(&user).Error; err != nil {
		return jsonResponse(c, fiber.StatusInternalServerError, "Failed to save data", err.Error())
	}

	return jsonResponse(c, fiber.StatusCreated, "Data successfully added", user)
}

func UpdateUs(c *fiber.Ctx) error {
	id := c.Params("username")

	var user models.User
	if err := models.DB.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return jsonResponse(c, fiber.StatusNotFound, "No data found", nil)
		}
		return jsonResponse(c, fiber.StatusInternalServerError, "Failed to load data", err.Error())
	}

	var update models.User
	if err := c.BodyParser(&update); err != nil {
		return jsonResponse(c, fiber.StatusBadRequest, "Invalid input", err.Error())
	}

	if update.Username != id {
		if err := models.DB.First(&models.User{}, update.Username).Error; err == nil {
			return jsonResponse(c, fiber.StatusBadRequest, "The updated ID is already in use", nil)
		}
	}

	if err := models.DB.Model(&user).Updates(update).Error; err != nil {
		return jsonResponse(c, fiber.StatusInternalServerError, "Failed to update data", err.Error())
	}

	return jsonResponse(c, fiber.StatusOK, "Data successfully updated", nil)
}

func DeleteUs(c *fiber.Ctx) error {
	id := c.Params("username")

	if models.DB.Delete(&models.User{}, id).RowsAffected == 0 {
		return jsonResponse(c, fiber.StatusNotFound, "Data not found or already deleted", nil)
	}

	return jsonResponse(c, fiber.StatusOK, "Successfully deleted data", nil)
}

func jsonResponse(c *fiber.Ctx, status int, message string, data interface{}) error {
	return c.Status(status).JSON(fiber.Map{
		"message": message,
		"data":    data,
	})
}
