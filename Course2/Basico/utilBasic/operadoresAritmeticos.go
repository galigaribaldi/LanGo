package utilBasic

import "fmt"

func operadoresAritmeticos() {
	// Area Cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area Cuadrado: ", areaCuadrado)
	x := 10
	y := 50
	// Suma
	result := x + y
	fmt.Println("Suma: ", result)
	//Resta
	result = y - x
	fmt.Println("Resta: ", result)
	// Multiplicacion
	result = x * y
	fmt.Println("Multiplicacion: ", result)
	//Division (1 parte)
	result = y / x
	fmt.Println("Division: ", result)
	//Modulo
	result = y % x
	fmt.Println("Modulo: ", result)
	// Incremental
	fmt.Println("Valor de X original:", x)
	x++
	fmt.Println("Incremental", x)
	// Decremental
	x--
	fmt.Println("Decremental: ", x)
}

//Area del rectangulo (No recibe parámetros)
func areaRectangulo() {
	x := 10
	y := 50
	area := x * y
	fmt.Println("Area Rectangulo: ", area)
}

//Area del Trapecio (No recibe parámetros)
func areaTrapecio() {
	BM := 10
	Bm := 50
	h := 20
	area := (BM * Bm * h) / 2
	fmt.Println("Area Trapecio: ", area)
}

//Area del Circulo (No recibe parámetros)
func areaCirculo() {
	r := 10
	const baseCuadrado = 10
	area := (r * r * baseCuadrado)
	fmt.Println("Area Circulo: ", area)
}
func main() {
	//operadoresAritmeticos()
	areaRectangulo()
	areaTrapecio()
	areaCirculo()
}
