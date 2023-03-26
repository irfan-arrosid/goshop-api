package migration

import (
	"fmt"
	"goshop-api/database"
	"goshop-api/models/entities"
	"log"
)

func RunMigration() {
	err := database.DB.AutoMigrate(
		&entities.User{},
	)

	if err != nil {
		log.Println("Database migration is failed")
	} else {
		fmt.Println("Database migrated")
	}
}
