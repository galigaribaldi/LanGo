package main

/// Before execute this file, you need execute = go mod init
import (
	"fmt"

	paq1 "github.com/galigaribaldi/LanGo/Intermedio/paquetes1"
)

func main() {
	fmt.Println("Paquete 1 Intermedio")
	paq1.ActionArrays()
	fmt.Println("")
	paq1.ActionCalculadora()
	fmt.Println("")
	paq1.ActionMaps()
	fmt.Println("")
	paq1.ActionStructs()
}
