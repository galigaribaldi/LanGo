package utilBasic

import "fmt"

func constantes() {
	//Declaracion de constantes
	const pi = 3.14 //En Automaticos
	const pi2 float64 = 3.14

	fmt.Println("Hola mundo")
	fmt.Println("Pi:", pi)
	fmt.Println("Pi 2:", pi2)
	// Declaración de Variables enteras
	base := 12          //Variable no creada anteriormente
	var altura int = 14 //Indicar el tipo de variable para omitir los ":="
	var area int        //Declarar el tipo de dato, pero no inicializamos
	fmt.Println("Variable 1 'Base':", base)
	fmt.Println("Variable 2 'altura':", altura)
	fmt.Println("Variable 1 'area':", area)
	// Zero Values Valores por defecto = y ""
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println(a, b, c, d)
}
