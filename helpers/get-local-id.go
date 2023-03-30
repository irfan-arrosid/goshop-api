package helpers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GetUserIDFromLocals(c *fiber.Ctx) uint {
	return uint(c.Locals("claims").(jwt.MapClaims)["id"].(float64))
}
