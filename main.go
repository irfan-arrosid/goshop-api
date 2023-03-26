package main

import (
	"goshop-api/database"
	"goshop-api/models/entities/migration"
)

func main() {
	database.DbConnect()
	migration.RunMigration()
}
