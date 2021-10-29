package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("leeme.txt")
	if err != nil {
		panic(err.Error())
	}
	//Imprime el codigo ASCII de cada caracter
	fmt.Println(data)
	//Decodificando el archivo
	fmt.Println(string(data))
}
