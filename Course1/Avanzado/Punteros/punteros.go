package main

import "fmt"

func main() {
	x := 100
	var y *int
	y = &x
	fmt.Println("Valor de x", x, " Direccion de memoria: ", &x)
	fmt.Println("Valor de y", *y, " Direccion de memoria:", y)
	*y = 500
	fmt.Println("Cambio de valor")
	fmt.Println("Valor de x", x, " Direccion de memoria: ", &x)
	fmt.Println("Valor de y", *y, " Direccion de memoria:", y)

}
