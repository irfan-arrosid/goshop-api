package handlers

import (
	"goshop-api/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func ProvinceGetAll(c *fiber.Ctx) error {
	data, err := utils.ProvinceGetAll()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"data": data,
	})
}

func ProvinceGetById(c *fiber.Ctx) error {
	id := c.Params("id")

	data, err := utils.ProvinceGetById(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
			"data":    nil,
		})
	}

	if data.Id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Params could not be zero or empty",
		})
	}

	return c.JSON(fiber.Map{
		"data": data,
	})
}

func CityGetAll(c *fiber.Ctx) error {
	// id := c.Params("id")

	data, err := utils.CityGetAll()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"data": data,
	})
}

func CityGetById(c *fiber.Ctx) error {
	id := c.Params("id")

	data, err := utils.CityGetById(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
			"data":    nil,
		})
	}

	if data.Id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Params could not be zero or empty",
		})
	}

	return c.JSON(fiber.Map{
		"data": data,
	})
}
