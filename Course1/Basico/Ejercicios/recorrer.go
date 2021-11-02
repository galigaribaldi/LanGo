package ejercicios

import "fmt"

func arregloCadena(cadena string) {
	fmt.Println("Tener cuidado de agregar la palabra 'string()' para recorrer")
	for i := 0; i < 3; i++ {
		fmt.Println(cadena[i])
	}
	fmt.Println("Recorriendo una cadena hasta 3 lugares")
	for i := 0; i < 3; i++ {
		fmt.Println(string(cadena[i]))
	}
	fmt.Println("Recorriendo una cadena con su extensión")
	for i := 0; i < len(cadena); i++ {
		fmt.Println(string(cadena[i]))
	}
}

func listasArrayStrings() {
	var array [5]string
	array[0] = "Arreglo 0"
	array[1] = "Arreglo 1"
	array[2] = "Arreglo 2"
	fmt.Println(array)
	for i := 0; i < len(array); i++ {
		if len(array[i]) == 0 {
			fmt.Printf("El lugar %d Esta vacio!\n", i)
		} else {
			fmt.Println(array[i])
		}
	}
}

//En Arreglos de tipo int no es necesario poner un Int
//para la impresión ya que los inicializa con 0
func listasArraysInt() {
	var array [10]int
	array[0] = 10
	array[1] = 15
	array[9] = 30
	for i := 0; i < len(array); i++ {
		if array[i] == 0 {
			fmt.Println("El arreglo en la posición: ", i, "esta vacío")
		} else {
			fmt.Println("Arreglo en la posición: ", i, "es: ", array[i])
		}
	}
}
