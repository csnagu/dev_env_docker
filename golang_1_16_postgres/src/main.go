package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	dsn := "host=postgres user= password= dbname= port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	new_product := Product{Price: 200, Code: "D42"}
	db.Create(&new_product)
	log.Println("Created product")

	// Read
	var products []Product
	db.Find(&products)
	log.Println("Read products")
	log.Println("products: ", products)

	// Update
	var product Product
	db.Find(&product, "code = ?", "D42")
	db.Model(&product).Update("Price", 300)
	log.Println("Updated product")

	db.Find(&products)
	log.Println("products: ", products)

	// Delete
	db.Delete(&product)
	log.Println("Deleted product")

	db.Find(&products)
	log.Println("products: ", products)
}
