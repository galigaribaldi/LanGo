package utilKit

import "fmt"

func getUnicode() string {
	return "弥彦の痛み"
}

func tipoDatoString() {
	hola := "Hola"
	fmt.Println(getUnicode())
	fmt.Println("Tamaño de la cadena Hola: ", len(hola))
	fmt.Println(string(hola[:2]))
}
