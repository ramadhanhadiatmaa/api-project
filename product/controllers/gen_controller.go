package controllers

import (
	"product/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ShowGen(c *fiber.Ctx) error {
	var gen []models.Gen

	if err := models.DB.Find(&gen).Error; err != nil {
		return jsonResponse(c, fiber.StatusInternalServerError, "Failed to load data", err.Error())
	}

	if len(gen) == 0 {
		return jsonResponse(c, fiber.StatusNotFound, "No data found", nil)
	}

	return c.JSON(gen)
}

func IndexGen(c *fiber.Ctx) error {
	id := c.Params("id")
	var gen models.Gen

	if err := models.DB.First(&gen, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return jsonResponse(c, fiber.StatusNotFound, "No data found", nil)
		}
		return jsonResponse(c, fiber.StatusInternalServerError, "Failed to load data", err.Error())
	}

	return c.JSON(gen)
}

func CreateGen(c *fiber.Ctx) error {
	var gen models.Gen

	if err := c.BodyParser(&gen); err != nil {
		return jsonResponse(c, fiber.StatusBadRequest, "Invalid input", err.Error())
	}

	if err := models.DB.Create(&gen).Error; err != nil {
		return jsonResponse(c, fiber.StatusInternalServerError, "Failed to save data", err.Error())
	}

	return jsonResponse(c, fiber.StatusCreated, "Data successfully added", gen)
}

func UpdateGen(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return jsonResponse(c, fiber.StatusBadRequest, "Invalid ID format", nil)
	}

	var gen models.Gen
	if err := models.DB.First(&gen, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return jsonResponse(c, fiber.StatusNotFound, "No data found", nil)
		}
		return jsonResponse(c, fiber.StatusInternalServerError, "Failed to load data", err.Error())
	}

	var updateGen models.Gen
	if err := c.BodyParser(&updateGen); err != nil {
		return jsonResponse(c, fiber.StatusBadRequest, "Invalid input", err.Error())
	}

	if updateGen.ID != 0 && updateGen.ID != id {
		if err := models.DB.First(&models.Gen{}, updateGen.ID).Error; err == nil {
			return jsonResponse(c, fiber.StatusBadRequest, "The updated ID is already in use", nil)
		}
	}

	if err := models.DB.Model(&gen).Updates(updateGen).Error; err != nil {
		return jsonResponse(c, fiber.StatusInternalServerError, "Failed to update data", err.Error())
	}

	return jsonResponse(c, fiber.StatusOK, "Data successfully updated", nil)
}

func DeleteGen(c *fiber.Ctx) error {
	id := c.Params("id")

	if models.DB.Delete(&models.Gen{}, id).RowsAffected == 0 {
		return jsonResponse(c, fiber.StatusNotFound, "Data not found or already deleted", nil)
	}

	return jsonResponse(c, fiber.StatusOK, "Successfully deleted data", nil)
}
