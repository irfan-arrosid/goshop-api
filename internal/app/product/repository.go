package product

import "gorm.io/gorm"

type Repository interface {
	CreateCategory(category Category) (Category, error)
	FindAllCategory() ([]Category, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateCategory(category Category) (Category, error) {
	err := r.db.Create(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *repository) FindAllCategory() ([]Category, error) {
	var categories []Category
	err := r.db.Find(&categories).Error
	if err != nil {
		return categories, err
	}

	return categories, nil
}
