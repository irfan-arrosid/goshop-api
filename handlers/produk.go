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

func ProdukGetAll(c *fiber.Ctx) error {
	var produk []entities.Produk

	pagination := helpers.GetPagination
	if pagination != nil {
		return c.Status(http.StatusNotAcceptable).JSON(fiber.Map{
			"message": "Failed to get pagination",
		})
	}

	err := database.DB.Find(&produk).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
			"data":    nil,
		})
	}

	// pagination.Data = produk

	return c.JSON(fiber.Map{
		"data": pagination,
	})
}

func ProdukGetById(c *fiber.Ctx) error {
	var produk entities.Produk
	produkId := c.Params("id")

	if produkId == "0" {
		return c.Status(http.StatusNotAcceptable).JSON(fiber.Map{
			"message": "Param must be a number greater than 0",
		})
	}

	err := database.DB.First(&produk, "id = ?", produkId).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Produk does not exist",
		})
	}

	return c.JSON(produk)
}

func ProdukCreate(c *fiber.Ctx) error {
	produkRequest := new(requests.Produk)
	if err := c.BodyParser(produkRequest); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad request",
		})
	}

	idUser := helpers.GetUserIDFromLocals(c)
	idToko := 0
	condition := fmt.Sprintf("SELECT id FROM toko WHERE id_user=%d", idUser)

	if err := database.DB.Raw(condition).Scan(&idToko).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	newProduk := entities.Produk{
		Nama_produk:    produkRequest.Nama_produk,
		Harga_reseller: produkRequest.Harga_reseller,
		Harga_konsumen: produkRequest.Harga_konsumen,
		Stok:           produkRequest.Stok,
		Deskripsi:      produkRequest.Deskripsi,
		Id_category:    produkRequest.Id_category,
	}

	errCreateProduk := database.DB.Create(&newProduk).Error
	if errCreateProduk != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Produk successfully stored",
	})
}

func ProdukUpdate(c *fiber.Ctx) error {
	var produk entities.Produk
	produkId := c.Params("id")

	if produkId == "0" {
		return c.Status(http.StatusNotAcceptable).JSON(fiber.Map{
			"message": "Param must be a number greater than 0",
		})
	}

	produkRequest := new(requests.Produk)
	if err := c.BodyParser(produkRequest); err != nil {
		return err
	}

	if err := database.DB.Where("id", produkId).Save(&produk).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
			"data":    produk,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Produk update success",
		"data":    produk,
	})
}

func ProdukDelete(c *fiber.Ctx) error {
	var produk entities.Produk
	produkId := c.Params("id")

	if err := database.DB.First(&produk, "id = ?", produkId).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Produk not found",
		})
	}

	if err := database.DB.Debug().Delete(&produk).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Produk was deleted",
	})
}
