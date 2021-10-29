package main

import (
	"fmt"
	"time"
)

func SendValue(c chan string) {
	fmt.Println("Execute Gourutine")
	time.Sleep(1 * time.Second)
	c <- "Hola mundo"
	fmt.Println("Finished Executing GoRoutine")
}

func main() {
	fmt.Println("Go Channels Tutorial")
	values := make(chan string, 6)
	defer close(values)

	go SendValue(values)
	go SendValue(values)
	go SendValue(values)
	go SendValue(values)
	go SendValue(values)
	go SendValue(values)

	value := <-values
	fmt.Println(value)
	time.Sleep(1 * time.Second)

}
