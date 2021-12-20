package conversions

func YardsInCentimeter(centimeter float32) float32 {
	centimeter = centimeter * GetYardMetter() * 100
	return centimeter
}

/*
func EnglishToInternacional(opcion string, value int) float32 {
	if opcion == "metterToIn" {
		c := metterToIn(value)
		return c
	} else {
		return 0
	}
}
*/
