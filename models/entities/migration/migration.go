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
		&entities.Alamat{},
		&entities.Category{},
		&entities.DetailTrx{},
		&entities.FotoProduk{},
		&entities.LogProduk{},
		&entities.Produk{},
		&entities.Toko{},
		&entities.Trx{},
	)

	if err != nil {
		log.Println("Database migration is failed")
	} else {
		fmt.Println("Database migrated")
	}
}
