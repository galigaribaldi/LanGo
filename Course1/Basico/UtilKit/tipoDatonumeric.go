package utilKit

import "fmt"

/*
Entero = int
Unsigned = uint (valores mas grandes) de 32 bits
int64 = int64
*/

func getMultipleValues2() (int, int32, int64) {
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
func tipoDatoNumeric() {
	var (
		a, b, c = getMultipleValues2()
	)
	fmt.Printf("Int normal: %v, Int de 32: %v, Int de 64: %v", a, b, c)
	fmt.Println("")
	decimal32, deccimal64 := getDecimal()
	fmt.Println(decimal32, deccimal64)

}
