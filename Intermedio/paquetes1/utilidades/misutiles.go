package utilidades

import "fmt"

func Getnumero() int {
	var numero = 0
	fmt.Println("Ingresa un numero: ")
	fmt.Scanf("%d", &numero)
	return numero
}

func Bucles1() {
	for i := 0; i < 3; i++ {
		fmt.Println("Bluce For: ", i)
	}
}

func Bulces2() {
	j := 50
	for j > 0 {
		fmt.Println("Ciclo While", j)
		j -= 10
	}
}

func Bucles3() {
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
