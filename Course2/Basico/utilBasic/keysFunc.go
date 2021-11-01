package utilBasic

import "fmt"

//Ejecuta la función antes de que muera (defer)
func defers() {
	//Defer
	defer fmt.Println("Ultimo 1")
	defer fmt.Println("Ultimo 2")
	defer fmt.Println("Ultimo 3")
	fmt.Println("Nuevo texto de un archivo")
}

//Continue = Cuando una condición dada dentro de un ciclo For se continue
//Ejemplo: Error
func continues() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		//Continue
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}
		//Break
		if i == 7 {
			fmt.Println("Break")
			break
		}
	}
}
func main() {
	defers()
	continues()
}
