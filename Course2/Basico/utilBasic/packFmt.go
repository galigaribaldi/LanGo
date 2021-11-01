package utilBasic

import "fmt"

//Pequeños ejercicios con el paquete inicial FMT
func packFmt() {
	//Declaración de Variables
	helloMessage := "Hello"
	worldMessage := "World"
	var3 := 40
	//Println
	fmt.Println(helloMessage, worldMessage)

	//Printf
	fmt.Printf("El mensaje 1 es: %s y el mensaje 2 es: %d\n", helloMessage, var3)
	//En caso de no saber el valor de la variable, le podnremos "%v"
	fmt.Printf("Buena Practica es con %v'")
	fmt.Printf("El mensaje 1 es: %v y el mensaje 2 es: %v\n", helloMessage, var3)

	// Sprintf = Guardar pero no imprimir
	message := fmt.Sprintf("El mensaje 1 es: %s y el mensaje 2 es: %d\n", helloMessage, var3)
	fmt.Println(message)

	//Tipo de dato
	fmt.Printf("%s es de tipo: %t\n", helloMessage, helloMessage)
	fmt.Printf("%s es de tipo: %T\n", helloMessage, helloMessage)
}

func main() {
	packFmt()
}
