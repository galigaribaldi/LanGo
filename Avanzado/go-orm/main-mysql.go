package main

import (
	"fmt"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

type Casa struct {
	gorm.Model
	Direccion string
	Numero    string
}

func main() {
	dsn := "test:test@tcp(127.0.0.1:8889)/go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error!")
	}
	fmt.Println("Conectado")
	db.AutoMigrate(&Casa{})
	db.Create(&Casa{Direccion: "Mi direccion", Numero: "25 - 1A"})
	var casa Casa
	var casa2 Casa
	db.First(&casa, 1)
	db.First(&casa2, "direccion = ?", "Mi direccion")
	fmt.Println(casa)
	fmt.Println(casa2)
	///Update
	//db.Model(&casa).Update("Direccion", "Mi nueva direccion")
	//Delete
	//db.Delete(&casa)
	//Consultas con Where
	var casas []Casa
	db.Where("numero LIKE ?", "%25%").Find(&casas)
	fmt.Println("")
	fmt.Println(casas)
}
