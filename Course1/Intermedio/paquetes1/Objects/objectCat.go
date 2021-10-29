package Objects

type Cat struct{
	Patas	int
	Color	string
	Edad	int
	Tam		float32
}

func (c Cat) Caminar() string{
	return "El gato camina con " +string( c.Patas)
}