package main

import "fmt"

/*
Resume: Programa que simula una calculadora sencilla
*/

//Función que realiza una suma
func suma(a int, b int) int {
	return a + b
}

//Función que realiza una resta
func resta(a int, b int) int {
	return a - b
}

//Función que realiza una división
func division(a int, b int) int {
	if b > 0 {
		return a / b
	} else {
		return 0
	}
}

//Función que realiza una multiplicación
func multiplicacion(a int, b int) int {
	return a * b
}

//Eleccion de operación
func eleccion(opcion int, valor1 int, valor2 int) int {
	var res int
	switch opcion {
	case 1:
		res = suma(valor1, valor2)
	case 2:
		res = resta(valor1, valor2)
	case 3:
		res = multiplicacion(valor1, valor2)
	case 4:
		res = division(valor1, valor2)
	default:
		fmt.Println("Elección no procesada")
		res = 0
	}
	return res
}

func main() {
	var opcion int
	var valor1 int
	var valor2 int
	var res int
	for true {
		fmt.Println("Opción 1: Suma\nOpcion 2: Resta\nOpcion 3: Multiplicación\nOpcion 4: División\nOpcion 5: Salir")
		fmt.Scanln(&opcion)
		if opcion >= 5 {
			println("Eleccion diferente")
			break
		} else {
			fmt.Println("Valor 1: ")
			fmt.Scanln(&valor1)
			fmt.Println("Valor 2: ")
			fmt.Scanln(&valor2)
			res = eleccion(opcion, valor1, valor2)
			fmt.Println("El resultado es: ", res)
		}
	}

}
