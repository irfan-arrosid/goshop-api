package main

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Customer struct {
	Id       uuid.UUID `gorm:"primaryKey"`
	Name     string
	Email    string
	Password string
	// gorm.Model
}

func (customer *Customer) BeforeCreate(tx *gorm.DB) (err error) {
	customer.Id = uuid.New()
	return nil
}

func main() {

	dsn := "root:my04sql04@tcp(localhost:3306)/dummy-project"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	db.AutoMigrate(&Customer{})

	// CREATE
	db.Create(&Customer{Name: "Test 01", Email: "test01@email.com", Password: "password01"})
	db.Create(&Customer{Name: "Test 02", Email: "test02@email.com", Password: "password02"})
	db.Create(&Customer{Name: "Test 03", Email: "test03@email.com", Password: "password03"})

	// READ ONE
	var customer Customer
	findOne := db.First(&customer)
	if findOne.Error != nil {
		log.Fatalf("Error: %v", findOne.Error)
	}

	fmt.Printf("Customer: %+v\n", customer)

	// READ ALL
	var customers []Customer
	findAll := db.Find(&customers)
	if findAll.Error != nil {
		log.Fatalf("Error: %v", findAll.Error)
	}

	fmt.Printf("Customer: %+v\n", customers)

	// r := gin.Default()
	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "Hello, World!",
	// 	})
	// })
	// r.Run()
}
