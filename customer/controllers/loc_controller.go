package controllers

import (
	"customer/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ShowLoc(c *fiber.Ctx) error {
	var loc []models.Loc

	if err := models.DB.Find(&loc).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch users",
		})
	}

	return c.JSON(loc)
}

func IndexLoc(c *fiber.Ctx) error {
	idloc := c.Params("id_loc")
	var loc models.Loc

	if err := models.DB.First(&loc, "id_loc = ?", idloc).Error; err != nil {
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

	return c.JSON(loc)
}

func CreateLoc(c *fiber.Ctx) error {
	var loc models.Loc

	if err := c.BodyParser(&loc); err != nil {
		return jsonResponse(c, fiber.StatusBadRequest, "Invalid input", err.Error())
	}

	if err := models.DB.Create(&loc).Error; err != nil {
		return jsonResponse(c, fiber.StatusInternalServerError, "Failed to save data", err.Error())
	}

	return jsonResponse(c, fiber.StatusCreated, "Data successfully added", loc)
}

func UpdateLoc(c *fiber.Ctx) error {
	idloc := c.Params("id_loc") // Ambil parameter username dari URL
	var loc models.Loc

	if err := models.DB.First(&loc, "id_loc = ?", idloc).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return jsonResponse(c, fiber.StatusNotFound, "User not found", nil)
		}
		return jsonResponse(c, fiber.StatusInternalServerError, "Failed to load user", err.Error())
	}

	if err := c.BodyParser(&loc); err != nil {
		return jsonResponse(c, fiber.StatusBadRequest, "Invalid input data", nil)
	}

	if err := models.DB.Save(&loc).Error; err != nil {
		return jsonResponse(c, fiber.StatusInternalServerError, "Failed to update user", err.Error())
	}

	return jsonResponse(c, fiber.StatusOK, "User updated successfully", loc)
}

func DeleteLoc(c *fiber.Ctx) error {
	idloc := c.Params("id_loc")

	var loc models.Loc
	if err := models.DB.First(&loc, "id_loc = ?", idloc).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return jsonResponse(c, fiber.StatusNotFound, "Data not found", nil)
		}
		return jsonResponse(c, fiber.StatusInternalServerError, "Failed to load data", err.Error())
	}

	if err := models.DB.Delete(&loc).Error; err != nil {
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
