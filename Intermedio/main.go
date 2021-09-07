package main

/// Before execute this file, you need execute = go mod init
import (
	/*
		"github.com/galigaribaldi/LanGo/Intermedio/paquetes1/arrays"
		"github.com/galigaribaldi/LanGo/Intermedio/paquetes1/operacionesstrings"
		"github.com/galigaribaldi/LanGo/Intermedio/paquetes1/maps"
		"github.com/galigaribaldi/LanGo/Intermedio/paquetes1/utilidades"
	*/
	"fmt"

	"github.com/galigaribaldi/LanGo/Intermedio/paquetes1/funcionesoperaciones"
)

func main() {
	miNumero, error := funcionesoperaciones.Addition("1", 2, 4)
	if error != nil {
		panic(error)
	} else {
		fmt.Println(miNumero)
	}

	/*
		//First Package
		arrays.ImprimeArray()
		arrays.ImprimeSlides()
		//Second Package
		operacionesstrings.ReemplazoStrings()
		//Third Package
		utilidades.Getnumero()
		utilidades.Bucles1()
		//foirth Package
		c := funcionesoperaciones.Addition(1, 2, 3)
		fmt.Println("El valor de la suma es: ", c)
	*/
	//fmt.Println(maps.GetMap())
	//fmt.Println(maps.GetMap2())
	//fmt.Println(maps.GetKeyMap("Be2atriz"))
}
