package main

import (
	"fmt"

	conversion "github.com/galigaribaldi/LanGo/Conversions"
)

func main() {
	fmt.Println("Hola mundo")
	fmt.Println("Yardas en metros:", conversion.GetYardMetter())
	fmt.Println("Yardas en Cm", conversion.YardsInCentimeter(1))
}
