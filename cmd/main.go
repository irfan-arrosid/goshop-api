package main

import (
	"goshop-api/internal/app/product"
	"goshop-api/internal/config"
	"goshop-api/internal/handler"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	config.Connect()
	config.Migration()

	// import repository
	productRepository := product.NewRepository(config.DB)

	// import service
	productService := product.NewService(productRepository)

	// import handler
	productHandler := handler.NewProductHandler(productService)

	r := gin.Default()
	r.Use(cors.Default())
	api := r.Group("/api/v1")

	api.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world",
		})
	})

	// PRODUCT endpoints
	api.POST("/categories", productHandler.NewCategory)
	api.GET("/categories", productHandler.GetCategories)

	r.Run(os.Getenv("PORT"))
}
