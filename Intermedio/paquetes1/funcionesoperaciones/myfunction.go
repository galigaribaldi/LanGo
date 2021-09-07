package funcionesoperaciones

import "errors"

func GetMultipleValues() (int, int, int) {
	return 1, 2, 3

}
func Addition(a interface{}, b interface{}, c interface{}) (int, error) {
	//enum y clojure en swift
	switch a.(type) {
	case string:
		return 0, errors.New("El valor es unString y no numerico")
	case bool:
		return 0, errors.New("El valor es unString y no numerico")
	}
	switch b.(type) {
	case string:
		return 0, errors.New("El valor no es correcto")
	}
	switch c.(type) {
	case string:
		return 0, errors.New("El valor no es correcto")
	}
	return a.(int) + b.(int) + c.(int), nil
}
func Addition2(a int, b int) int {
	return a + b
}
