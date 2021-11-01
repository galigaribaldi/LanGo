package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	routines "github.com/galigaribaldi/LanGo/Course2/Basico/GouRoutines"
	mypackage "github.com/galigaribaldi/LanGo/Course2/Basico/myPackage"
	"github.com/labstack/echo"
)

func main() {
	//Instanciar Echo
	//Ruta

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))

	fmt.Println("Nueco")

}
func funcion2() {
	var myCar mypackage.CarPublic
	myCar.Brand = "Ferrari"
	myCar.Year = 2021
	fmt.Println(myCar)
	c := mypackage.PrintMessage()
	fmt.Println(c)
	//mypackage.StructsPunteros()
	myPc1 := mypackage.MyPc{Ram: 20, Disk: 20, Band: "ARM"}
	fmt.Println(myPc1)
	//interface go
	myCuadrado := mypackage.Cuadrado{Base: 10.1}               //Cuadrado
	myRectangulo := mypackage.Rectangulo{Base: 20, Altura: 10} //Rectángulo
	fmt.Println(myCuadrado, myRectangulo)
	mypackage.Calcular(myCuadrado)
	mypackage.Calcular(myRectangulo)
	mypackage.InterfaceList()
	///Rutinesas
	var wg sync.WaitGroup
	fmt.Println("Hello 1")
	wg.Add(1)

	go routines.Say("Worlf", &wg)
	wg.Wait()

	go func(text string) {
		fmt.Println(text)
	}("Adios")
	time.Sleep(time.Second * 1)

	//Channels
	routines.Channel1()
	routines.Channel2()
}
