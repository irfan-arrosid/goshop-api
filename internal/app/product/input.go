package product

type CreateCategoryInput struct {
	Name string `json:"name" binding:"required"`
}
