package middleware

import (
	"goshop-api/helpers"
	"goshop-api/utils"

	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) error {
	token := c.Get("token")
	if token == "" {
		return helpers.BuildResponse(c, false, "UNAUTHORIZED", "UNAUTHORIZED", nil, fiber.StatusUnauthorized)
	}

	claims, err := utils.DecodeToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}

	// role := claims["role"].(string)
	// if role != "admin" {
	// 	return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
	// 		"message": "Forbidden access",
	// 	})
	// }

	// c.Locals("userInfo", claims)

	c.Locals("userid", claims["id"])
	c.Locals("useremail", claims["email"])

	return c.Next()
}
