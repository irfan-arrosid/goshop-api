package config

import (
	"fmt"
	"goshop-api/internal/app/product"
	"goshop-api/internal/app/user"
)

func Migration() {
	err := DB.AutoMigrate(&user.User{}, &product.Product{}, &product.Category{})

	if err != nil {
		panic("Failed to migrate schema")
	} else {
		fmt.Println("Migration success")
	}
}
