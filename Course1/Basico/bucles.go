package main

import "fmt"

func bucles() {
	for i := 0; i < 3; i++ {
		fmt.Println("Bluce For: ", i)
	}
}

func bulces2() {
	j := 50
	for j > 0 {
		fmt.Println("Ciclo While", j)
		j -= 10
	}
}

func bulces3() {
	k := 500
	for {
		k -= 1
		fmt.Printf("%d ", k)
		if k == 0 {
			fmt.Println("Termino el bucle infinito")
			break
		}
	}
}
func main() {
	bucles()
	bulces2()
	bulces3()
}
