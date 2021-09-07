package operacionesstrings

import (
	"fmt"
	"strings"
)

func ReemplazoStrings() {
	var texto = "Hola Go, Hola mundo, hola sr"
	fmt.Println(texto)
	///Mayuscula
	fmt.Println(strings.ToUpper(texto))
	///Minuscula
	fmt.Println(strings.ToLower(texto))
	///Replace 1 vez
	var texto2 = (strings.Replace(texto, "Hola", "Adios", 1))
	fmt.Println(texto2)
	///Replace Todas
	var texto3 = (strings.Replace(texto, "Hola", "Adios", -1))
	fmt.Println(texto3)
}
func ComparacionStrings() {
	var texto12 = "Hola Go, Hola mundo, hola sr"
	fmt.Println(texto12)
	//Devuelve -1 Si son diferentes
	fmt.Println(strings.Compare(texto12, "Nueva cadea"))
	//Devuelve 0 Si son iguales
	fmt.Println(strings.Compare(texto12, texto12))
	///Separado
	fmt.Println(strings.Split(texto12, ","))
	///Encnotrar
	fmt.Println(strings.Contains(texto12, "Hola"))
	fmt.Println(strings.Contains(texto12, "Coche"))
}
