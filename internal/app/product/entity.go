package product

import (
	"goshop-api/internal/app/user"

	"gorm.io/gorm"
)

type Product struct {
	Id          int
	UserId      int
	Name        string
	Description string
	CategoryId  int
	Price       float64
	Quantity    int
	ImageURL    string
	User        user.User // Product can only belong to single User
	Category    Category  // Product can only belong to single Category
	gorm.Model
}

type Category struct {
	Id   int
	Name string
	gorm.Model
}

// func (product *Product) BeforeCreate(tx *gorm.DB) (err error) {
// 	product.Id = uuid.New()
// 	return nil
// }
