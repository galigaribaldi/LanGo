package gouroutines

import (
	"fmt"
	"sync"
)

func Say1(text string, c chan<- string) {
	c <- text
}

func Say(text string, wg *sync.WaitGroup) {
	fmt.Println(text)
	defer wg.Done()
}
func Channel1() {
	c := make(chan string, 1)
	println("Hello")
	go Say1("Bye", c)
	fmt.Println(<-c)
}

func Channel2() {
	C := make(chan string, 2)
	C <- "Mensaje1"
	C <- "Mensaje2"
	fmt.Println(len(C), cap(C))
	//Cerrar los canales una vez que dejas de usarlo
	close(C)
	//C <- "Mensaje3"
	///Iterar entre mensajes
	for message := range C {
		fmt.Println(message)
	}
	//Select
	email1 := make(chan string)
	email2 := make(chan string)
	go Say1("Mensaje1", email1)
	go Say1("Mensaje2", email2)
	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email Recibido de email", m1)
		case m2 := <-email2:
			fmt.Println("Email Recibido de email2", m2)
		}
	}
}
