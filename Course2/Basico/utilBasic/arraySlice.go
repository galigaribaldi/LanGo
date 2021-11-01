package utilBasic

import (
	"fmt"
	"strings"
)

//Tipso de datos inmutables
func arrays() {
	//Array
	var array [4]int
	fmt.Println(array)
	array[0] = 1
	array[2] = 2
	fmt.Println(array)
	//Len = Elementos en un array
	//cap = Capacidad maxima en un array
	fmt.Println(array, "Elementos del array", len(array), cap(array))
}

//Recorriendo slices y Arrays
func arraysSliceRecorrer() {
	slice := []string{"Hola", "Que", "Hace"}
	for i, valor := range slice {
		fmt.Println(i, valor)
	}
	for _, valor := range slice {
		fmt.Println(valor)
	}
}

//Funcion si es palindromo
func isPalindromo(text string) {
	var textReverse string
	for i := len(text) - 1; i >= 0; i-- {
		textReverse += strings.ToUpper(string(text[i]))
	}
	text = strings.ToUpper(text)
	if textReverse == text {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindormo")
	}
	fmt.Printf("Palabra original es: %s y la palabra alreves es: %s\n", text, textReverse)
}

//Slice = Listas Mutables, arreglos de datos dinamicos
func slicing() {
	//Slice
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))
	/// Metodos en el slicing
	fmt.Println("Pisición 0:", slice[0])
	fmt.Println("Hasta el 3: ", slice[:3])
	fmt.Println("Lugar 2 al 4: ", slice[2:4])
	fmt.Println("De 4 en Adelante: ", slice[4:])
	//Append
	slice = append(slice, 1)
	fmt.Println(slice)
	// Append Nueva list
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)
}
