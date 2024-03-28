package handler

import (
	"goshop-api/internal/app/product"
	"goshop-api/internal/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	productService product.Service
}

func NewProductHandler(productService product.Service) *productHandler {
	return &productHandler{productService}
}

func (h *productHandler) NewCategory(c *gin.Context) {
	var input product.CreateCategoryInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Create new category is failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newCategory, err := h.productService.NewCategory(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Create new category is failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := product.FormatCategory(newCategory)
	response := helper.APIResponse("Category has been created", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *productHandler) GetCategories(c *gin.Context) {
	categories, err := h.productService.GetCategories()
	if err != nil {
		response := helper.APIResponse("Error to get categories", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := product.FormatCategories(categories)
	response := helper.APIResponse("Successfully get categories", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}
