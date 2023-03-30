package handlers

import (
	"goshop-api/database"
	"goshop-api/helpers"
	"goshop-api/models/entities"
	"goshop-api/models/requests"
	"log"

	"github.com/gofiber/fiber/v2"
)

func AlamatGetAll(c *fiber.Ctx) error {
	var alamat []entities.Alamat
	result := database.DB.Find(&alamat)
	if result.Error != nil {
		log.Println(result.Error)
	}

	return c.JSON(alamat)
}

func AlamatGetById(c *fiber.Ctx) error {
	var alamat entities.Alamat
	alamatId := c.Params("id")

	err := database.DB.First(&alamat, "id = ?", alamatId).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Alamat does not exist",
		})
	}

	return c.JSON(alamat)
}

func AlamatCreate(c *fiber.Ctx) error {
	alamatRequest := new(requests.AlamatCreateRequest)
	if err := c.BodyParser(alamatRequest); err != nil {
		return err
	}

	idUser := helpers.GetUserIDFromLocals(c)

	newAlamat := entities.Alamat{
		Judul_alamat:  alamatRequest.Judul_alamat,
		Nama_penerima: alamatRequest.Nama_penerima,
		Notelp:        alamatRequest.Notelp,
		Detail_alamat: alamatRequest.Detail_alamat,
	}

	newAlamat.Id_user = idUser

	errCreateAlamat := database.DB.Create(&newAlamat).Error
	if errCreateAlamat != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Stroring data is failed",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Alamat successfully stored",
	})
}

func AlamatUpdate(c *fiber.Ctx) error {
	alamatRequest := new(requests.AlamatCreateRequest)
	if err := c.BodyParser(alamatRequest); err != nil {
		return err
	}

	var alamat entities.Alamat
	alamatId := c.Params("id")

	err := database.DB.First(&alamat, "id = ?", alamatId).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Alamat does not exist",
		})
	}

	if alamatRequest.Judul_alamat != "" {
		alamat.Judul_alamat = alamatRequest.Judul_alamat
	}
	alamat.Nama_penerima = alamatRequest.Nama_penerima
	alamat.Notelp = alamatRequest.Notelp
	alamat.Detail_alamat = alamatRequest.Detail_alamat

	errUpdate := database.DB.Save(&alamat).Error
	if errUpdate != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
			"data":    alamat,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Alamat update success",
		"data":    alamat,
	})
}

func AlamatDelete(c *fiber.Ctx) error {
	var alamat entities.Alamat
	alamatId := c.Params("id")

	err := database.DB.First(&alamat, "id = ?", alamatId).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Alamat does not exist",
		})
	}

	errDelete := database.DB.Debug().Delete(&alamat).Error
	if errDelete != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Alamat was deleted",
	})
}
