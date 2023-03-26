package routes

import (
	"goshop-api/handlers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(r *fiber.App) {
	r.Post("/user", handlers.UserCreate)
	r.Get("/user", handlers.UserGetAll)
	r.Get("/user/:id", handlers.UserGetById)
	r.Put("/user/:id", handlers.UserUpdate)
	r.Put("/user/:id/update-email", handlers.UserUpdateEmail)
	r.Delete("/user/:id", handlers.UserDelete)
}
