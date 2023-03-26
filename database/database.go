package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DbConnect() {
	db, err := gorm.Open(mysql.Open(os.Getenv("DSN_MYSQL")), &gorm.Config{})

	if err != nil {
		fmt.Println("Connection database is failed.")
	} else {
		fmt.Println("Database connected....")
	}

	DB = db
}
