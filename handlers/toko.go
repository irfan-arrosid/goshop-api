package handlers

import (
	"fmt"
	"goshop-api/database"
	"goshop-api/helpers"
	"goshop-api/models/entities"
	"goshop-api/models/requests"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func TokoGetAll(c *fiber.Ctx) error {
	var toko []entities.Toko

	pagination := helpers.GetPagination
	if pagination != nil {
		return c.Status(http.StatusNotAcceptable).JSON(fiber.Map{
			"message": "Failed to get pagination",
		})
	}

	err := database.DB.Find(&toko).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
			"data":    nil,
		})
	}

	// pagination.Data = toko

	return c.JSON(fiber.Map{
		"data": pagination,
	})
}

func TokoGetById(c *fiber.Ctx) error {
	var toko entities.Toko
	tokoId := c.Params("id")

	if tokoId == "0" {
		return c.Status(http.StatusNotAcceptable).JSON(fiber.Map{
			"message": "Param must be a number greater than 0",
		})
	}

	err := database.DB.First(&toko, "id = ?", tokoId).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Toko does not exist",
		})
	}

	return c.JSON(toko)
}

func GetMyToko(c *fiber.Ctx) error {
	idUser := helpers.GetUserIDFromLocals(c)
	var toko entities.Toko

	err := database.DB.Where("id_user", idUser).First(toko).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Toko does not exist",
		})
	}

	if err := database.DB.Find(&toko).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
			"data":    nil,
		})
	}

	return c.JSON(toko)
}

func TokoUpdate(c *fiber.Ctx) error {
	var toko entities.Toko
	tokoId := c.Params("id")

	if tokoId == "0" {
		return c.Status(http.StatusNotAcceptable).JSON(fiber.Map{
			"message": "Param must be a number greater than 0",
		})
	}

	tokoRequest := new(requests.TokoUpdateRequest)
	if err := c.BodyParser(tokoRequest); err != nil {
		return err
	}

	file, err := c.FormFile("photo")
	if err != nil {
		return c.Status(http.StatusNotAcceptable).JSON(fiber.Map{
			"message": "Something error",
		})
	}

	url, err := helpers.ProcessImage(c, file, "toko")
	if err != nil {
		return c.Status(http.StatusNotAcceptable).JSON(fiber.Map{
			"message": "Failed to processing image",
		})
	}

	idUser := helpers.GetUserIDFromLocals(c)
	condition := fmt.Sprintf("id_user = %d AND id = %d", idUser, tokoId)

	if err := database.DB.Where(condition).First(toko).Updates(url).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Toko does not exist",
		})
	}

	if err := database.DB.Find(&toko).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
			"data":    nil,
		})
	}

	return c.JSON(toko)
}
