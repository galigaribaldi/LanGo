package utilBasic

import "fmt"

func main() {
	m := make(map[string]int)
	m["Jose"] = 14
	m["Pepe"] = 20
	m["John"] = 40
	fmt.Println(m)
	//Recorriendo Maps
	for i, v := range m {
		fmt.Println(i, v)
	}
	//Encontrar un valor
	value := m["Jose"]
	fmt.Println("El valor Jose: ", value)
	//Si no se encuntra la llave se retorna un 0 value, para ello se reptrna Error
	value1, ok := m["Joseph"]
	if !ok {
		fmt.Println("No existe el valor Joseph")
	}
	fmt.Println("El valor inexistente de Joseph: ", value1, ok)
}
