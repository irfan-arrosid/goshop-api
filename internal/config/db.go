package config

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(mysql.Open(os.Getenv("DSN_MYSQL")), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	} else {
		fmt.Println("Connection has been established successfully")
	}

	DB = db
}
