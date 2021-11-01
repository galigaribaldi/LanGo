package mypackage

import "fmt"

type pc struct {
	ram  int
	disk int
	band string
}

func (myPc pc) ping() {
	fmt.Println(myPc.band, "Pong")

}

func (myPC *pc) duplicateRAM() {
	myPC.ram = myPC.ram * 2
}

func StructsPunteros() {
	a := 50
	b := &a
	fmt.Println(a)
	fmt.Println(*b)
	*b = 100
	fmt.Println("Valor de a:", a, "Valor de b:", b)
	var myPC pc
	myPC.ram = 10
	myPC.disk = 10
	myPC.band = "ARM"
	myPC.ping()
	///Getters y Setters
	fmt.Println(myPC)
	myPC.duplicateRAM()
	fmt.Println(myPC)
	myPC.duplicateRAM()

}
