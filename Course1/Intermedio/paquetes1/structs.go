package paquetes1

import "fmt"

type Galisteo string

// Persona = son los Atributos
type Persona struct {
	Nombre             string
	Apellido           string
	DocumentoIdentidad string
	Telefono           []string
	Direccion          string
	Edad               int
}

type Casa struct {
	NumeroCasa int
	Persona    []Persona
}

//Metodos Obtener Numero de casa
func (c Casa) GetnumeroCasa() {
	fmt.Println("El numero de la casa es: ", c.NumeroCasa)

}

//Metodos Obtener Numero de Personas
func (c Casa) GetPersonasCasa() {
	defer SoyDefer()
	fmt.Println("El numero de la casa es: ", c.Persona)

}
func SoyDefer() {
	fmt.Println("Defer esto que se ejecute al final")
}
