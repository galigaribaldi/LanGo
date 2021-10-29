package calculadora

// Calculadora
type Calculadora struct {
	OperandoA int
	OperandoB int
	OperandoC int
}

//Suma Metodo para sumar los elementos A y B, retorna el operando C
func (c Calculadora) Suma() int {
	return c.OperandoA + c.OperandoB
}

//Resta Metodo para restar los elementos A y B, retorna el operando C
func (c Calculadora) Resta() int {
	return c.OperandoA - c.OperandoB
}

//Multiplicacion Metodo para mulriplicar los elementos A y B, retorna el operando C
func (c Calculadora) Multiplicacion() int {
	return c.OperandoA * c.OperandoB
}

//Division Metodo para Dividir los elementos A y B, retorna el operando C, si marca, error este lo avisa
func (c Calculadora) Division() int {
	if c.OperandoB <= 0 {
		return 0
	} else {
		return c.OperandoA / c.OperandoB
	}
}
