package main

/*
	Inserciones con el teclado con condicinales
	1.- Condicionales con If
	2.- Condicionales con Switch - Case
*/
import "fmt"

func main() {
	var opcion = 0
	fmt.Print("Ingresa tu Opcion: ")
	fmt.Scanf("%d", &opcion)
	//opcionif(opcion)
	opcionCase(opcion)
}

///Funcion con condicionales con "If"
func opcionif(opcion int) {
	if opcion == 0 {
		fmt.Println("Opcion 0")
	}
	if opcion == 1 {
		fmt.Println("Opcion 1")
	} else {
		fmt.Println("Opcion no seleccionada", opcion)
	}
}

//Funciones con condicional Case
func opcionCase(opcion int) {
	switch opcion {
	case 1:
		fmt.Println("La opcion selecccionaste es 1")
	case 2:
		fmt.Println("La opcion selecccionaste es 2")
	default:
		fmt.Println("Opcion no considerada en el prgorama: ", opcion)
	}
}
