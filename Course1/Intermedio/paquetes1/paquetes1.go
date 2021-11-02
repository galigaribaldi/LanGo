package paquetes1

import (
	"fmt"
)

func ActionArrays() {
	fmt.Println("Nuevo")
	imprimeArray()
	imprimeSlides()
}

func ActionCalculadora() {
	var valor1 int
	calc := Calculadora{OperandoA: 23, OperandoB: 34, OperandoC: 3}
	fmt.Printf("Valor de Calculadora, operando a: %d, operando b: %d, operando c: %d", calc.OperandoA, calc.OperandoB, calc.OperandoC)
	//Suma
	valor1 = calc.Suma()
	fmt.Printf("Valor de la suma: %d\n", valor1)
	//Resta
	valor1 = calc.Resta()
	fmt.Printf("Valor de la resta: %d\n", valor1)
	//Multiplicacion
	valor1 = calc.Multiplicacion()
	fmt.Printf("Valor de la multiplicación: %d\n", valor1)
	//División
	valor1 = calc.Division()
	fmt.Printf("Valor de la División: %d\n", valor1)
}

func ActionMaps() {
	c := GetMap()
	fmt.Println("Mapa: ", c)
	c = GetMap2()
	fmt.Println("Mapa: ", c)
	d := GetKeyMap("Maria")
	fmt.Println("Retorna el valor de la llave Maria:", d)
}
func ActionStructs() {

	antonio := Persona{
		Nombre:    "Antonio Jose",
		Apellido:  "Galisteo Cantero",
		Telefono:  []string{"111", "22"},
		Direccion: "YYY",
		Edad:      30,
	}
	fmt.Println(antonio)
	fmt.Println(antonio.Apellido)
	maria := Persona{
		Nombre:             "Maria",
		Apellido:           "Alvarado",
		DocumentoIdentidad: "YYY",
		Telefono:           []string{"333", "444"},
		Direccion:          "ZZZ",
		Edad:               31,
	}
	fmt.Println("Persona Maria: ", maria)
	////-------------- * = Puntero al valor
	jorge := new(Persona)
	jorge.Nombre = "Jorge"
	jorge.Apellido = "Apellidos"
	jorge.DocumentoIdentidad = "MMM"
	jorge.Telefono = []string{"ttt", "sssa"}
	jorge.Direccion = "AAA"
	jorge.Edad = 40
	fmt.Println(jorge)
	fmt.Println(*jorge)
	/// Arreglo de Structs
	casa := Casa{
		NumeroCasa: 1,
		Persona:    []Persona{antonio, *jorge, maria},
	}
	fmt.Println("Objeto de tipo Casa: ", casa)
	casa.GetnumeroCasa()
	casa.GetPersonasCasa()
}
