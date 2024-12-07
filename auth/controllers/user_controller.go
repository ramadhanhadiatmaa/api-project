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
	username := c.Params("username")
	var user models.User

	// Preload TypeInfo untuk memuat data dari tabel type_user
	if err := models.DB.Preload("TypeInfo").First(&user, "username = ?", username).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "No data found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Failed to load data",
		})
	}

	// Format data sesuai yang diinginkan
	response := map[string]interface{}{
		"username": user.Username,
		"password": user.Password,
		"type":     user.TypeInfo.Type, // Mengambil type dari TypeInfo
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
	username := c.Params("username") // Ambil parameter username dari URL
	var user models.User

	// Cari data user berdasarkan username
	if err := models.DB.First(&user, "username = ?", username).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return jsonResponse(c, fiber.StatusNotFound, "User not found", nil)
		}
		return jsonResponse(c, fiber.StatusInternalServerError, "Failed to load user", err.Error())
	}

	// Parse body request ke struct user
	if err := c.BodyParser(&user); err != nil {
		return jsonResponse(c, fiber.StatusBadRequest, "Invalid input data", nil)
	}

	// Jika data valid, lakukan update
	if err := models.DB.Save(&user).Error; err != nil {
		return jsonResponse(c, fiber.StatusInternalServerError, "Failed to update user", err.Error())
	}

	return jsonResponse(c, fiber.StatusOK, "User updated successfully", user)
}

func DeleteUs(c *fiber.Ctx) error {
	username := c.Params("username")

	var user models.User
	if err := models.DB.First(&user, "username = ?", username).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return jsonResponse(c, fiber.StatusNotFound, "Data not found", nil)
		}
		return jsonResponse(c, fiber.StatusInternalServerError, "Failed to load data", err.Error())
	}

	if err := models.DB.Delete(&user).Error; err != nil {
		return jsonResponse(c, fiber.StatusInternalServerError, "Failed to delete data", err.Error())
	}

	return jsonResponse(c, fiber.StatusOK, "Successfully deleted data", nil)
}

func jsonResponse(c *fiber.Ctx, status int, message string, data interface{}) error {
	return c.Status(status).JSON(fiber.Map{
		"message": message,
		"data":    data,
	})
}
