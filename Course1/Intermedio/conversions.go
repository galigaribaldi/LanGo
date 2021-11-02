package main

import (
	"fmt"

	conv "github.com/galigaribaldi/LanGo/Intermedio/paquetes1/Conversiones"
)

func main() {
	fmt.Println("Main 3")
	valor := conv.MetterToIn(15.1)
	fmt.Println(valor)
	conv.Metodo2()
	//valor1 := Conversiones.mettertokm(13.2)
	//fmt.Println(valor1)
}
