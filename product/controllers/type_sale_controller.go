package controllers

import (
	"product/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ShowSale(c *fiber.Ctx) error {
	var typesale []models.TypeSale

	if err := models.DB.Find(&typesale).Error; err != nil {
		return jsonResponse(c, fiber.StatusInternalServerError, "Failed to load data", err.Error())
	}

	if len(typesale) == 0 {
		return jsonResponse(c, fiber.StatusNotFound, "No data found", nil)
	}

	return c.JSON(typesale)
}

func IndexSale(c *fiber.Ctx) error {
	id := c.Params("id_sale")
	var typesale models.TypeSale

	if err := models.DB.First(&typesale, "id_sale = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return jsonResponse(c, fiber.StatusNotFound, "No data found", nil)
		}
		return jsonResponse(c, fiber.StatusInternalServerError, "Failed to load data", err.Error())
	}

	return c.JSON(typesale)
}

func CreateSale(c *fiber.Ctx) error {
	var typesale models.TypeSale

	if err := c.BodyParser(&typesale); err != nil {
		return jsonResponse(c, fiber.StatusBadRequest, "Invalid input", err.Error())
	}

	if err := models.DB.Create(&typesale).Error; err != nil {
		return jsonResponse(c, fiber.StatusInternalServerError, "Failed to save data", err.Error())
	}

	return jsonResponse(c, fiber.StatusCreated, "Data successfully added", typesale)
}

func UpdateSale(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id_sale"))
	if err != nil {
		return jsonResponse(c, fiber.StatusBadRequest, "Invalid ID format", nil)
	}

	var typesale models.TypeSale
	if err := models.DB.First(&typesale, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return jsonResponse(c, fiber.StatusNotFound, "No data found", nil)
		}
		return jsonResponse(c, fiber.StatusInternalServerError, "Failed to load data", err.Error())
	}

	var updateSale models.TypeSale
	if err := c.BodyParser(&updateSale); err != nil {
		return jsonResponse(c, fiber.StatusBadRequest, "Invalid input", err.Error())
	}

	if updateSale.IdSale != 0 && updateSale.IdSale != id {
		if err := models.DB.First(&models.TypeSale{}, updateSale.IdSale).Error; err == nil {
			return jsonResponse(c, fiber.StatusBadRequest, "The updated ID is already in use", nil)
		}
	}

	if err := models.DB.Model(&typesale).Updates(updateSale).Error; err != nil {
		return jsonResponse(c, fiber.StatusInternalServerError, "Failed to update data", err.Error())
	}

	return jsonResponse(c, fiber.StatusOK, "Data successfully updated", nil)
}

func DeleteSale(c *fiber.Ctx) error {
	id := c.Params("id_sale")

	if models.DB.Delete(&models.TypeSale{}, id).RowsAffected == 0 {
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
