package main

import (
	"fmt"
	"strconv"
)

//main Ejecución de diferentes hilos
func main() {
	fmt.Println("Hola")
	canal := make(chan string)
	lanzaHilos(200, canal)
	for valor := range canal {
		fmt.Println(valor)
	}
	fmt.Println(<-canal)
}

func holaMundo(i int, canal chan<- string) {
	//fmt.Println("Hola mundo: ", i)
	canal <- "Hola mundo Numero: " + strconv.Itoa(i)
}
func lanzaHilos(numHilos int, canal chan<- string) {
	// canal chan string (entrada y escritura)
	// canal <-chan string = Salida
	// canal chan<- string = Entrada
	for i := 0; i < numHilos; i++ {
		go holaMundo(i, canal)
	}
}
