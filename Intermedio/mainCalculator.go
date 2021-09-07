package main

/// Before execute this file, you need execute = go mod init
import (
	"fmt"

	"github.com/galigaribaldi/LanGo/Intermedio/paquetes1/calculadora"
)

func main() {
	calcu := calculadora.Calculadora{
		OperandoA: 1,
		OperandoB: 0,
	}
	fmt.Println(calcu)
	fmt.Println("El resultado de la suma es: ", calcu.Suma())
	fmt.Println("El resultado de la Resta es: ", calcu.Resta())
	fmt.Println("El resultado de la Resta es: ", calcu.Division())
}
