package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type Persona struct {
	Nombre    string
	Apellidos string
}

func handlerPersonas(response http.ResponseWriter, request *http.Request) {
	persona := Persona{"Antonio", "Galisteo"}
	jsResponse, err := json.Marshal(persona)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	response.Write(jsResponse)
}

func handlerRaiz(response http.ResponseWriter, request *http.Request) {
	io.WriteString(response, "Hola mundo API")
}
func handlerUsuarios(response http.ResponseWriter, request *http.Request) {
	io.WriteString(response, "Hola Usuarios")
}
func main() {
	//Route
	http.HandleFunc("/", handlerRaiz)
	http.HandleFunc("/usuarios", handlerUsuarios)
	//Json
	http.HandleFunc("/personas", handlerPersonas)
	http.ListenAndServe(":8000", nil)
}
