package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create (inserts)
	db.Create(&Product{Code: "P1", Price: 100})
	/// (Selects)
	var product Product
	db.First(&product, 1)
	db.First(&product, "code = ?", "P1")
}
