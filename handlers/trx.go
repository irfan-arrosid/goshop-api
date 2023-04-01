package handlers

import (
	"goshop-api/database"
	"goshop-api/helpers"
	"goshop-api/models/entities"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func TrxGetAll(c *fiber.Ctx) error {
	var trx []entities.Trx

	pagination := helpers.GetPagination
	if pagination != nil {
		return c.Status(http.StatusNotAcceptable).JSON(fiber.Map{
			"message": "Failed to get pagination",
		})
	}

	if err := database.DB.Preload("Alamat").Preload("DetailTrx").Preload("DetailTrx.Toko").Preload("DetailTrx.LogProduk").Preload("DetailTrx.LogProduk.Toko").Preload("DetailTrx.LogProduk.Category").Preload("DetailTrx.LogProduk.Photos").Find(&trx).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
			"data":    nil,
		})
	}

	return c.JSON(pagination)
}

func TrxGetById(c *fiber.Ctx) error {
	var trx entities.Trx
	id := c.Params("id")

	if id == "0" {
		return c.Status(http.StatusNotAcceptable).JSON(fiber.Map{
			"message": "Param must be a number greater than 0",
		})
	}

	if err := database.DB.Preload("Alamat").Preload("DetailTrx").Preload("DetailTrx.Toko").Preload("DetailTrx.LogProduk").Preload("DetailTrx.LogProduk.Toko").Preload("DetailTrx.LogProduk.Category").Preload("DetailTrx.LogProduk.Photos").First(trx, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "No data transaction",
		})
	}

	return c.JSON(trx)
}

// func TrxCreate(c *fiber.Ctx) error {
// 	trxRequest := new(requests.Trx)
// 	if err := c.BodyParser(trxRequest); err != nil {
// 		return c.Status(400).JSON(fiber.Map{
// 			"message": "Bad request",
// 		})
// 	}

// }
