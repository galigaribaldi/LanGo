package main

import "fmt"

func imprimeArray() {
	var array1 [2]string
	array1[0] = "Hola"
	array1[1] = "Mundo"
	//array1[2] = "Error"
	fmt.Println("Tamaño del arreglo: ", len(array1))
	fmt.Println(array1)
	///
	array4 := [4]int{1, 2, 3, 4}
	fmt.Println(array4)
	///
	//var arrayVacio []int
	var matriz [2][2]string
	matriz[0][0] = "Hola"
	matriz[0][1] = "Mimdo"
	matriz[1][0] = "nuevo"
	matriz[1][1] = "Nuevo3"
	fmt.Println(matriz)
	////
	matriz22 := [2][2]int{
		{1, 2},
		{3, 4}}
	fmt.Println(matriz22)
}

///Slices o bien las listas
func imprimeSlides() {
	var slice1 []string
	slice1 = append(slice1, "Hola")
	fmt.Println(slice1)
	fmt.Println("Tamaño_ ", len(slice1))
	slice1 = append(slice1, "Nuevo elemento", "Nuevo elemento 2")
	fmt.Println(slice1)
	fmt.Println("Tamaño_ ", len(slice1))
}

//// Matrices
func imprimeMatrices() {
	matriz := [][]string{
		{"Hola", "Mundo"},
		{"Curso", "Go"}}
	fmt.Println(matriz)
	var matriz2 [][]string
	row1 := []string{"Hola2", "SliceMatriz2"}
	row2 := []string{"Hola2", "Go2"}
	matriz2 = append(matriz2, row1)
	matriz2 = append(matriz2, row2)
	fmt.Println(matriz2)
}
func main() {
	fmt.Println("Arrays")
	imprimeArray()
	fmt.Println("Slides (lsitas)")
	fmt.Println("Slides (lsitas)")
	imprimeSlides()
	imprimeMatrices()
}
