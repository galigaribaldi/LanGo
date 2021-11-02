package ejercicios

import "fmt"

func ActionCalculadora() {
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
func ActionRecorrer() {
	arregloCadena("Cadena2")
	listasArrayStrings()
	listasArraysInt()
}
