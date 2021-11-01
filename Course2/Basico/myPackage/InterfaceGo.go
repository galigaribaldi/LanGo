package mypackage

import "fmt"

type Cuadrado struct {
	Base float64
}

type Rectangulo struct {
	Base   float64
	Altura float64
}

type figuras2D interface {
	Area() float64
}

func (c Cuadrado) Area() float64 {
	return c.Base * c.Base
}

func (r Rectangulo) Area() float64 {
	return r.Base * r.Altura
}

func Calcular(f figuras2D) {
	fmt.Println("Area: ", f.Area())
}

/*
Lista de interfaces = [1,"",True] = Lista de diferentes valores
*/
func InterfaceList() {
	myInterface := []interface{}{"Hola", 12, 9.90}
	fmt.Println(myInterface...)
}
