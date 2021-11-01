package utilBasic

import "fmt"

type car struct {
	brand   string
	year    int
	color   string
	creador string
	nuevo   bool
}

func newStrcut() {
	myCar := car{
		brand: "BMW",
		year:  2020,
	}
	fmt.Println(myCar)
	///Segunda forma de Instanciar
	var otherCar car
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar)
}
