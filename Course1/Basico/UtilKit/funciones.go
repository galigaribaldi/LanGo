package utilKit

import "fmt"

func getName() string {
	var name string = "Nombre por Defecto"
	fmt.Print("Ingresa tu nombre: ")
	fmt.Scanf("%s", &name)
	return name
}

func getMultipleValues() (int, int, int) {
	return 1, 2, 3

}
func suma(a int, b int, c int) int {
	return a + b + c
}
func suma2(a int, b int) int {
	return a + b
}
func ActionFunction() {
	name := getName()
	fmt.Println(name)
	x, y, z := getMultipleValues()
	fmt.Println(x, y, z)
	var b = suma(1, 2, 3)
	fmt.Println(b)
	var minumero = suma2(1, 5)
	fmt.Println(minumero)
}
