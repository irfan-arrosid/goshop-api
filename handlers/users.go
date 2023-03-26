package handlers

import (
	"goshop-api/database"
	"goshop-api/models/entities"
	"goshop-api/models/requests"
	"goshop-api/utils"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UserCreate(c *fiber.Ctx) error {
	user := new(requests.UserCreateRequest)
	if err := c.BodyParser(user); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(user)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Storing data is failed",
			"error":   errValidate.Error(),
		})
	}

	newUser := entities.User{
		Nama:          user.Nama,
		Notelp:        user.Notelp,
		Tanggal_lahir: user.Tanggal_lahir,
		Jenis_kelamin: user.Jenis_kelamin,
		Tentang:       user.Tentang,
		Pekerjaan:     user.Pekerjaan,
		Email:         user.Email,
		Id_provinsi:   user.Id_provinsi,
		Id_kota:       user.Id_kota,
		IsAdmin:       user.IsAdmin,
	}

	hashedPassword, err := utils.HashingPassword(user.Kata_sandi)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	newUser.Kata_sandi = hashedPassword

	errCreateUser := database.DB.Create(&newUser).Error
	if errCreateUser != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Storing data is failed",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data successfully stored",
	})
}

func UserGetAll(c *fiber.Ctx) error {
	// userInfo := c.Locals("userInfo")
	// log.Println("user info data :: ", userInfo)

	var users []entities.User
	result := database.DB.Find(&users)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return c.JSON(users)
}

func UserGetById(c *fiber.Ctx) error {
	userId := c.Params("id")

	var user entities.User

	err := database.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Request is success",
		"data":    user,
	})
}

func UserUpdate(c *fiber.Ctx) error {
	userRequest := new(requests.UserUpdateRequest)

	if err := c.BodyParser(userRequest); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad request",
		})
	}

	var user entities.User

	userId := c.Params("id")

	err := database.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	if userRequest.Nama != "" {
		user.Nama = userRequest.Nama
	}
	user.Notelp = userRequest.Notelp
	user.Tanggal_lahir = userRequest.Tanggal_lahir
	user.Jenis_kelamin = 
}
