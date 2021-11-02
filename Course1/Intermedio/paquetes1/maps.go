package paquetes1

//GetMap Retorna un objeto de tipo Dicccionario o Mapa
func GetMap() map[string]int {
	myMap := make(map[string]int)
	myMap["edad1"] = 18
	myMap["edad2"] = 19
	myMap["edad3"] = 20
	return myMap
}

func GetMap2() map[string]int {
	myMap := map[string]int{
		"Antonio": 30,
		"Maria":   31,
		"Beatriz": 32,
	}
	myMap["edad1"] = 18
	myMap["edad2"] = 19
	delete(myMap, "edad1")
	return myMap

}

func GetKeyMap(key string) int {
	myMap := map[string]int{
		"Antonio": 30,
		"Maria":   31,
		"Beatriz": 32,
	}
	myMap["edad1"] = 18
	myMap["edad2"] = 19
	delete(myMap, "edad1")
	return myMap[key]
}
