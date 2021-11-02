package utilKit

import "fmt"

//ActionBucles Main que selecciona entre For normal, While y Do While
func ActionBucles(opcion, max, decremento int) {
	fmt.Printf("Maximo de Rango:%d y Decremento: %d \n", max, decremento)
	switch opcion {
	case 1:
		fmt.Println("Ejemplo de Bucle For")
		bucles(max)
	case 2:
		fmt.Println("Ejemplo de Ciclo While")
		bulces2(max, decremento)
	case 3:
		fmt.Println("Ejemplo de Do While")
		bucles3(max, decremento)
	}
}

//ActionConditionals Main que selecciona entre los condicionales If,Switch
func ActionConditionals(opcion int) {
	switch opcion {
	case 1:
		fmt.Println("Condicional con If normal")
		condicionalesIf()
	case 2:
		fmt.Println("Condicional de tipo Switch Sencillo\nSólo evalúa Case")
		switch1()
	case 3:
		fmt.Println("Condicional de tipo Switch con condicionales lógicas")
		switch2()
	}
}

func ActionTipoDato(opcion int) {
	switch opcion {
	case 1:
		fmt.Println("Tipo de dato String")
		getUnicode()
		tipoDatoString()
	case 2:
		fmt.Println("Tipo de dato Numeric")
		tipoDatoNumeric()
	case 3:
		fmt.Println("Tipo de dato Array")
		tipoDatoArray()
	case 4:
		fmt.Println("Funciones")
		ActionFunction()
	}
}
