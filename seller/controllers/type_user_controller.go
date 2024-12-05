package controllers

import (
	"seller/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ShowTy(c *fiber.Ctx) error {
	var typuser []models.TypeUser

	if err := models.DB.Find(&typuser).Error; err != nil {
		return jsonResponse(c, fiber.StatusInternalServerError, "Failed to load data", err.Error())
	}

	if len(typuser) == 0 {
		return jsonResponse(c, fiber.StatusNotFound, "No data found", nil)
	}

	return c.JSON(typuser)
}

func IndexTy(c *fiber.Ctx) error {
	id := c.Params("id_type")
	var typuser models.TypeUser

	if err := models.DB.First(&typuser, "id_type = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return jsonResponse(c, fiber.StatusNotFound, "No data found", nil)
		}
		return jsonResponse(c, fiber.StatusInternalServerError, "Failed to load data", err.Error())
	}

	return c.JSON(typuser)
}

func CreateTy(c *fiber.Ctx) error {
	var typuser models.TypeUser

	if err := c.BodyParser(&typuser); err != nil {
		return jsonResponse(c, fiber.StatusBadRequest, "Invalid input", err.Error())
	}

	if err := models.DB.Create(&typuser).Error; err != nil {
		return jsonResponse(c, fiber.StatusInternalServerError, "Failed to save data", err.Error())
	}

	return jsonResponse(c, fiber.StatusCreated, "Data successfully added", typuser)
}

func UpdateTy(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id_type"))
	if err != nil {
		return jsonResponse(c, fiber.StatusBadRequest, "Invalid ID format", nil)
	}

	var typuser models.TypeUser
	if err := models.DB.First(&typuser, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return jsonResponse(c, fiber.StatusNotFound, "No data found", nil)
		}
		return jsonResponse(c, fiber.StatusInternalServerError, "Failed to load data", err.Error())
	}

	var updateType models.TypeUser
	if err := c.BodyParser(&updateType); err != nil {
		return jsonResponse(c, fiber.StatusBadRequest, "Invalid input", err.Error())
	}

	if updateType.IdType != 0 && updateType.IdType != id {
		if err := models.DB.First(&models.TypeUser{}, updateType.IdType).Error; err == nil {
			return jsonResponse(c, fiber.StatusBadRequest, "The updated ID is already in use", nil)
		}
	}

	if err := models.DB.Model(&typuser).Updates(updateType).Error; err != nil {
		return jsonResponse(c, fiber.StatusInternalServerError, "Failed to update data", err.Error())
	}

	return jsonResponse(c, fiber.StatusOK, "Data successfully updated", nil)
}

func DeleteTy(c *fiber.Ctx) error {
	id := c.Params("id_type")

	if models.DB.Delete(&models.TypeUser{}, id).RowsAffected == 0 {
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
