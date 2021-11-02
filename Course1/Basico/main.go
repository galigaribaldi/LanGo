package main

import (
	"fmt"

	ejerc "github.com/galigaribaldi/LanGo/Course1/Basico/Ejercicios"
	utilKit "github.com/galigaribaldi/LanGo/Course1/Basico/UtilKit"
)

func main() {
	fmt.Printf("1.- Bucles\n2.- Condicionales\n3.- Tipos de dato\n4.-Calculadora\n5.-Recorrer\nSeleccicon opcion>>")
	var opcion int
	var opcion2 int
	fmt.Scanf("%d", opcion)
	switch opcion {
	case 1:
		fmt.Printf("1.- Bulce For\n2.- Bucle While\n3.- Bucle Do While\nSeleccicon opcion>>")
		fmt.Scanf("%d", opcion2)
		utilKit.ActionBucles(opcion2, 10, 1)
	case 2:
		fmt.Printf("1.- If \n2.- Switch Normal\n3.- Switch con Condicional\nSeleccicon opcion>>")
		fmt.Scanf("%d", opcion2)
		utilKit.ActionConditionals(opcion2)
	case 3:
		fmt.Printf("1.- String\n2.- Numerico\n3.- Array\n4.-Funciones\nSeleccicon opcion>>")
		fmt.Scanf("%d", opcion2)
		utilKit.ActionTipoDato(opcion2)
	case 4:
		fmt.Println("Action Calculadora")
		ejerc.ActionCalculadora()
	case 5:
		fmt.Println("Recorrer")
		ejerc.ActionRecorrer()
	}

}

//go run main.go = Sin geenrar un ejecutable
