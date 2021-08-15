package main

import "fmt"

/*
Programa que simula una calculadora sencilla
*/

func suma(a int, b int) int {
	return a + b
}

func resta(a int, b int) int {
	return a - b
}
func division(a int, b int) int {
	if b > 0 {
		return a / b
	} else {
		return 0
	}
}

func multiplicacion(a int, b int) int {
	return a * b
}
func main() {
	var sumas = suma(1, 2)
	var restas = resta(1, 2)
	var divs = division(1, 2)
	var mult = multiplicacion(1, 5)
	fmt.Println("Suma: ", sumas)
	fmt.Println("Restas: ", restas)
	fmt.Println("Division: ", divs)
	fmt.Println("Multiplicacion: ", mult)

}
