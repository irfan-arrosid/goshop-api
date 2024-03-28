package product

import (
	"goshop-api/internal/app/user"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	Id          string `gorm:"type:varchar(36);primaryKey"`
	UserId      string
	Name        string
	Description string
	CategoryId  string
	Price       float64
	Quantity    int
	ImageURL    string
	User        user.User // Product can only belong to single User
	Category    Category  // Product can only belong to single Category
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Category struct {
	Id        string `gorm:"type:varchar(36);primaryKey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (product *Product) BeforeCreate(tx *gorm.DB) (err error) {
	if product.Id == "" {
		product.Id = uuid.New().String() // should be string in mysql
	}
	return nil
}

func (category *Category) BeforeCreate(tx *gorm.DB) (err error) {
	if category.Id == "" {
		category.Id = uuid.New().String() // should be string in mysql
	}
	return nil
}
