package handlers

import (
	"goshop-api/database"
	"goshop-api/models/entities"
	"goshop-api/models/requests"
	"log"

	"github.com/gofiber/fiber/v2"
)

func CategoryGetAll(c *fiber.Ctx) error {
	var categories []entities.Category
	result := database.DB.Find(&categories)
	if result.Error != nil {
		log.Println(result.Error)
	}

	return c.JSON(categories)
}

func CategoryGetById(c *fiber.Ctx) error {
	var categories entities.Category
	categoryId := c.Params("id")
	err := database.DB.First(&categories, "id = ?", categoryId).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Category does not exist",
		})
	}

	return c.JSON(categories)
}

func CategoryCreate(c *fiber.Ctx) error {
	categoryRequest := new(requests.CategoryCreateRequest)
	if err := c.BodyParser(categoryRequest); err != nil {
		return err
	}

	newCategory := entities.Category{
		Nama_category: categoryRequest.Nama_category,
	}

	errCreateCategory := database.DB.Create(&newCategory).Error
	if errCreateCategory != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Stroring data is failed",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Category successfully stored",
	})
}

func CategoryUpdate(c *fiber.Ctx) error {
	categoryRequest := new(requests.CategoryCreateRequest)
	if err := c.BodyParser(categoryRequest); err != nil {
		return err
	}

	var categories entities.Category
	categoryId := c.Params("id")

	err := database.DB.First(&categories, "id = ?", categoryId).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Category does not exist",
		})
	}

	if categoryRequest.Nama_category != "" {
		categories.Nama_category = categoryRequest.Nama_category
	}

	errUpdate := database.DB.Save(&categories).Error
	if errUpdate != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
			"data":    categories,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Category update success",
		"data":    categories,
	})
}

func CategoryDelete(c *fiber.Ctx) error {
	var categories entities.Category
	categoryId := c.Params("id")
	err := database.DB.First(&categories, "id = ?", categoryId).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Category does not exist",
		})
	}

	errDelete := database.DB.Debug().Delete(&categories).Error
	if errDelete != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Category was deleted",
	})
}
