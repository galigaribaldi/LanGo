package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Concurrency with GoRoutines")
	go compute(5)
	go compute(5)
	fmt.Scanln()
}

func compute(value int) {
	for i := 0; i < value; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}
