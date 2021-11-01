package utilBasic

import "fmt"

//Func par o impar
func parImpar(numero int) string {
	if numero%2 == 0 {
		fmt.Println("Es par")
		return "Es par"
	} else {
		fmt.Println("Es impar")
		return "Es impar"
	}
}

//Ifs
func Conditionals() {
	valor1 := 1
	valor2 := 2
	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}
	// With And
	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Es verdad")
	}
	// With Or
	if valor1 == 0 || valor2 == 2 {
		fmt.Println("Puerta OR Es verdad")
	}
}

//Switches 1
func Switches() {
	modulo := 5 % 2
	switch modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar!")
	}
}

//Switches 2
func Switches2() {

	switch modulo := 5 % 2; modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar!")
	}
}

//Swithces sin Condicion
func Switches3() {
	value := 50
	switch {
	case value > 100:
		fmt.Println("Es mayor a 100")
	case value < 0:
		fmt.Println("Es menor a 0")
	default:
		fmt.Println("Defaulr")
	}
}
func main() {
	//var n int
	//fmt.Printf("Ingresa un numero: ")
	//fmt.Scanf("%d", &n)
	//c := parImpar(n)
	//fmt.Println(c)
	Switches()
	Switches2()
	Switches3()
}
