package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var x int
	x = 8
	y := 7
	fmt.Println(x, y)
	myValue, err := strconv.ParseInt("", 0, 64)
	//Manejo de errores
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(myValue)
	}
	///Maps
	m := make(map[string]int)
	m["Key"] = 6
	fmt.Println(m["Key"])
	//Slices
	s := []int{1, 2, 3}
	for index, value := range s {
		fmt.Println(index, value)
	}
	fmt.Println("Agregando")
	s = append(s, 16)
	for index, value := range s {
		fmt.Println(index, value)
	}
	//Go Routine
	c := make(chan int)
	go doSomething(c)
	<-c
	//Apuntadores
	g := 25
	fmt.Println(g)
	h := &g
	fmt.Println(h)
	fmt.Println(*h)
}

//Go Routines
func doSomething(c chan int) {
	time.Sleep(3 * time.Second)
	println("Done 1")
	c <- 1
}
