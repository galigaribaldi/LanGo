package main

import (
	"fmt"
	"github.com/galigaribaldi/LanGo/Intermedio/paquetes1/Objects"
)

func main() {
	//Objetos sin Constructor ni Getter ni Setter
	charlie := Objects.Cat{
		Patas: 4,
		Color: "Gris",
		Edad: 23,
		Tam: 3.2,
	}
	fmt.Println(charlie)
	//Objetos con Getters y Setters
	jett := Objects.NewPerro("Maltes", 4, "Negro")
	fmt.Println(jett)
	//------------------------
	//Getters
	fmt.Println("La raza del Jett es: ", jett.GetRaza())
	//Setter
	jett.SetterRaza("Pastor Aleman")
	fmt.Println("La raza del Jett es: ",jett.GetRaza())
	//------------------------	
	//Getters
	fmt.Println("El jett tiene ",jett.GetPata(), "Patas", &jett.Patas)
	//Setter
	jett.SetterPata(8)
	fmt.Println("El jett tiene ",jett.GetPata(), "Patas", &jett.Patas)
	//------------------------	
	//Getter
	fmt.Println("El jett es de color: ", jett.GetColor(), &jett.Color)
	jett.SetterColor("Gris Obscuro")
	fmt.Println("El jett es de color: ", jett.GetColor(), &jett.Color)
	//----------Metodos propios--------------
}
