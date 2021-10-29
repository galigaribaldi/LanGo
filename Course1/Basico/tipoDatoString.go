package main

import "fmt"

func getUnicode() string {
	return "弥彦の痛み"
}

func main() {
	hola := "Hola"
	fmt.Println(getUnicode())
	fmt.Println("Tamaño de la cadena Hola: ", len(hola))
	fmt.Println(string(hola[:2]))
}
