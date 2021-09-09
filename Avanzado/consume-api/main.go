package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Tarea struct {
	UserId    int    `json:userId`
	Id        int    `json:id`
	Title     string `json:title`
	Completed bool   `json:completed`
}

func main() {
	var urlApi = "https://jsonplaceholder.typicode.com/todos"
	var cliente = &http.Client{Timeout: 10 * time.Second}
	response, err := cliente.Get(urlApi)
	if err != nil {
		panic(err.Error())
	}
	var tareas []Tarea
	//Se pasa el puntero a las tareas
	err = json.NewDecoder(response.Body).Decode(&tareas)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(tareas)
}
