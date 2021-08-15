package main

import "fmt"

/*
Entero = int
Unsigned = uint (valores mas grandes) de 32 bits
int64 = int64
*/

func getMultipleValues() (int, int32, int64) {
	return 1, 23113111, 3
}

/*
	Floatnates, hay 2 tipos de datos:
	Flotantes normales
	Flotantes con mas decimales
*/
func getDecimal() (float32, float64) {
	return float32(0.1), float64((float32(0.1)))
}
func main() {
	var (
		a, b, c = getMultipleValues()
	)
	fmt.Printf("Int normal: %s, Int de 32: %s, Int de 64: %s", a, b, c)
	fmt.Println("")
	decimal32, deccimal64 := getDecimal()
	fmt.Println(decimal32, deccimal64)

}
