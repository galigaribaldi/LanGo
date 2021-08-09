package main

import "fmt"

const holamundo string = "Hola %s %s, bienvenido al curso de GO"

func main() {
	var name string = "Nombre por defecto" //Asigna el valor a una variable previamente creada
	lasName := "Variable"                  //infiere el valor de la variable
	var numero = 100
	var (
		x = 1
		y = 2
		z = 3
	)
	fmt.Println("")
	fmt.Print("Ingresa tu nombre: ")
	fmt.Scanf("%s", &name)
	fmt.Println("Hola ", name)
	fmt.Printf("Hola %s, Bienvenido al curso de go", name)
	fmt.Println("")
	fmt.Printf(holamundo, name, lasName)
	fmt.Println("")
	fmt.Println(numero, x, y, z)
	fmt.Print("Direccion de memoria: ", &name)

}

//go run main.go = Sin geenrar un ejecutable
