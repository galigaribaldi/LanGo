package utilKit

import "fmt"

// condicionalesIf Condicionales con If
func condicionalesIf() {
	if numero := getNumero(); numero%5 == 0 {
		fmt.Println("Es Multiplo de 5")

	} else if numero%3 == 0 {
		fmt.Println("Es multiplo de 3")
	} else {
		fmt.Println("No es multiplo de 3 ni de 5")
	}
}

//getNumero Obtener un numero
func getNumero() int {
	var numero = 0
	fmt.Println("Ingresa un numero: ")
	fmt.Scanf("%d", &numero)
	return numero
}

//Switch1 con solo casos "case"
func switch1() int {
	numero := getNumero()
	switch numero {
	case 1:
		fmt.Println("El numero es 1")
		return 1
	case 2:
		fmt.Println("Eñ numero es 2")
		return 2
	default:
		fmt.Println("Ningno")
		return 0
	}
}

//Switch2 con casos "case" con condicionales lógicas
func switch2() {
	numero := getNumero()
	switch {
	case numero%5 == 0:
		fmt.Println("Es el 5")
	}
}
