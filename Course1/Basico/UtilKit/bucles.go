package utilKit

import "fmt"

// bucles Ejemplo de Bucles con ciclo normal
func bucles(max int) {
	for i := 0; i < max; i++ {
		fmt.Printf("%d\t", i)
	}
}

//bulces2 Ejemplo de implementación de While Usando For
func bulces2(max, decremento int) {
	for max > 0 {
		fmt.Printf("%d\t", max)
		max -= decremento
	}
}

// bucles3 Ejemplo de como funciona un Do-While
func bucles3(max, decremento int) {
	for {
		max -= decremento
		fmt.Printf("%d\t", max)
		if max == 0 {
			fmt.Println("Termino el bucle infinito")
			break
		}
	}
}
