package main

/// Before execute this file, you need execute = go mod init
import (
	"fmt"

	"github.com/galigaribaldi/LanGo/Intermedio/paquetes1/structs"
)

func main() {

	antonio := structs.Persona{
		Nombre:    "Antonio Jose",
		Apellido:  "Galisteo Cantero",
		Telefono:  []string{"111", "22"},
		Direccion: "YYY",
		Edad:      30,
	}
	fmt.Println(antonio)
	fmt.Println(antonio.Apellido)
	maria := structs.Persona{
		Nombre:             "Maria",
		Apellido:           "Alvarado",
		DocumentoIdentidad: "YYY",
		Telefono:           []string{"333", "444"},
		Direccion:          "ZZZ",
		Edad:               31,
	}
	fmt.Println("Persona Maria: ", maria)
	////-------------- * = Puntero al valor
	jorge := new(structs.Persona)
	jorge.Nombre = "Jorge"
	jorge.Apellido = "Apellidos"
	jorge.DocumentoIdentidad = "MMM"
	jorge.Telefono = []string{"ttt", "sssa"}
	jorge.Direccion = "AAA"
	jorge.Edad = 40
	fmt.Println(jorge)
	fmt.Println(*jorge)
	/// Arreglo de Structs
	casa := structs.Casa{
		NumeroCasa: 1,
		Persona:    []structs.Persona{antonio, *jorge, maria},
	}
	fmt.Println("Objeto de tipo Casa: ", casa)
	casa.GetnumeroCasa()
	casa.GetPersonasCasa()
}
