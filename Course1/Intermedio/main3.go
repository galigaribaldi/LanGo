package main
import (
	"fmt"
	"github.com/galigaribaldi/LanGo/Intermedio/paquetes1/Conversiones"
)

func main(){
	fmt.Println("Main 3")
	valor := Conversiones.MetterToIn(15.1)
	fmt.Println(valor)
	Conversiones.Metodo2()
	//valor1 := Conversiones.mettertokm(13.2)
	//fmt.Println(valor1)
}