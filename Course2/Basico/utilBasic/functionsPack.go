package utilBasic

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}
func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
	fmt.Printf("%v es: %T y %v es %T\n", a, a, b, b, c, c)

}
func returnValue(a int) int {
	return a * 2
}
func return2Value(a int) (c, d int) {
	return a, a * 2
}
func main() {
	normalFunction("Hola")
	tripleArgument(1, 2, "Hola")
	value := returnValue(2)
	fmt.Println("Value: ", value)
	value1, value2 := return2Value(2)
	fmt.Println("Value1: ", value1, " Value 2", value2)
	//Retornar solo 1 valor
	value3, _ := return2Value(2)
	fmt.Println("Escapando los valores: ", value3)

}
