package utils

import (
	"errors"
	"fmt"
	"goshop-api/helpers"
	"log"
	"mime/multipart"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

var DefaultPathAssetImage = helpers.ProjectRootPath + "/public/img/"

func HandleSingleFile(c *fiber.Ctx) error {
	file, errFile := c.FormFile("photo")
	if errFile != nil {
		log.Println("Error file : ", errFile)
	}

	var fileName *string
	var newFileName string
	if file != nil {
		errCheckContentType := checkContentType(file, "image/jpg", "image/jpeg", "image/png", "image/gif")
		if errCheckContentType != nil {
			log.Println(errCheckContentType)
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"message": errCheckContentType.Error(),
			})
		}

		fileName = &file.Filename
		newFileName = fmt.Sprintf("%d-%s", time.Now().Unix(), *fileName)
		arrStr := strings.Split(newFileName, " ")
		newFileName = strings.Join(arrStr, "-")

		errSaveFile := c.SaveFile(file, fmt.Sprintf("./public/img/%s", newFileName))
		if errSaveFile != nil {
			log.Println("errSaveFile", errSaveFile)
			log.Println("Failed to store file into public/img/directory ")
		}

		log.Println("Succeed to store file into public/img/directory ")
	} else {
		log.Println("Nothing file uploaded")
	}

	if fileName != nil {
		c.Locals("filename", newFileName)
	} else {
		c.Locals("filename", nil)
	}

	return c.Next()
}

func checkContentType(file *multipart.FileHeader, contentTypes ...string) error {
	if len(contentTypes) > 0 {
		for _, contentType := range contentTypes {
			contentTypeFile := file.Header.Get("Content-Type")
			if contentTypeFile == contentType {
				return nil
			}
		}
		return errors.New("not allowed file type")
	} else {
		return errors.New("not found content type to be checking")
	}
}
