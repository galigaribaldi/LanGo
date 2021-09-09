package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Persona struct {
	ID        int
	Nombre    string
	Apellidos string
}

// before execute this, execute the command go get github.com/go-sql-driver/mysql
func main() {
	fmt.Println("Go con Mysql")
	db, err := sql.Open("mysql", "test:test@tcp(127.0.0.1:8889)/go")
	if err != nil {
		panic(err)
	}
	fmt.Println("Conectado")
	results, err := db.Query("SELECT id, nombre, apellidos FROM persona")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var persona Persona
		err = results.Scan(&persona.ID, &persona.Nombre, &persona.Apellidos)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(persona)
	}
	defer db.Close()
}
