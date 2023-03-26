package main

import (
	"goshop-api/database"
	"goshop-api/models/entities/migration"
	"goshop-api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	database.DbConnect()
	migration.RunMigration()

	app := fiber.New()

	routes.UserRoutes(app)

	app.Listen(":8080")
}
