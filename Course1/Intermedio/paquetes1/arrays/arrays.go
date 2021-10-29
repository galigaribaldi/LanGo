package arrays

import "fmt"

/*
	Cuidado con no llamar a los paquetes con los nombres
	de las librerias creadas
*/
// ImprimeArray Muestra como se trabaja con arrays y matrices en GO
func ImprimeArray() {
	var array1 [2]string
	array1[0] = "Hola"
	array1[1] = "Mundo"
	fmt.Println("Tamaño del arreglo: ", len(array1))
	fmt.Println(array1)
	///
	array4 := [4]int{1, 2, 3, 4}
	fmt.Println(array4)
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

///ImprimeSlides Miestra como se trabaja con slides con GO
func ImprimeSlides() {
	var slice1 []string
	slice1 = append(slice1, "Hola")
	fmt.Println(slice1)
	fmt.Println("Tamaño_ ", len(slice1))
	slice1 = append(slice1, "Nuevo elemento", "Nuevo elemento 2")
	fmt.Println(slice1)
	fmt.Println("Tamaño_ ", len(slice1))
}
