package routes

import (
	"goshop-api/handlers"
	"goshop-api/middleware"

	"github.com/gofiber/fiber/v2"
)

// func AuthRoutes
// Post /register handlers.UserCreate
// Post /login handlers.UserLogin

func UserRoutes(r *fiber.App) {
	r.Post("/user", handlers.UserCreate)
	r.Get("/user", middleware.Auth, handlers.UserGetAll)
	r.Get("/user/:id", handlers.UserGetById)
	r.Put("/user/:id", handlers.UserUpdate)
	r.Put("/user/:id/update-email", handlers.UserUpdateEmail)
	r.Delete("/user/:id", handlers.UserDelete)
	r.Post("/login", handlers.UserLogin)
}

func CategoryRoutes(r *fiber.App) {
	r.Get("/category", handlers.CategoryGetAll)
	r.Get("/category/:id", handlers.CategoryGetById)
	r.Post("/category", handlers.CategoryCreate)
	r.Put("/category/:id", handlers.CategoryUpdate)
	r.Delete("/category/:id", handlers.CategoryDelete)
}

func AlamatRoutes(r *fiber.App) {
	r.Get("/user/alamat", handlers.AlamatGetAll)
	r.Get("/user/alamat/:id", handlers.AlamatGetById)
	r.Post("/user/alamat", handlers.AlamatCreate)
	r.Put("/user/alamat/:id", handlers.AlamatUpdate)
	r.Delete("user/alamat/:id", handlers.AlamatDelete)
}
