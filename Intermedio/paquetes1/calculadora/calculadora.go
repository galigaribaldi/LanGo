package calculadora

// Calculadora = son los Atributos
type Calculadora struct {
	operandoA int
	operandoB int
	operandoC int
}

//Metodos Obtener Numero de casa
func (c Calculadora) Suma() {

	c.operandoC := c.operandoA + c.operandoB
	return c.operandoC

}
