package Objects

// Perro
type Perro struct {
	Raza  string
	Patas int
	Color string
}

// Ejemplo de método constructor
func NewPerro(nombre string, patas int, color string) *Perro{
	return &Perro{
		Raza: nombre,
		Patas: patas,
		Color: color,
	}
}
/*
	1.- Métodos Getters de Perro
*/

//Getter de Raza
func (p *Perro) GetRaza() string{
	return p.Raza
}
//Getter de Patas
func (p *Perro) GetPata() int{
	return p.Patas
}
//Getter de Color
func (p *Perro) GetColor() string{
	return p.Color
}
/*
	2.- Metodos Setter de Perro
*/

//Setter de Raza
func (p *Perro) SetterRaza(r string){
	p.Raza = r
}
//Setter de Pata
func (p *Perro) SetterPata(pata int){
	p.Patas = pata
}

//Setter de Pata
func (p *Perro) SetterColor(color string){
	p.Color = color
}
/*
	Métodos - Acciones de el objeto
*/

// Métodos
func (p *Perro) Ladrar(l string) string {
	return "El perro " + p.Raza + " está ladrando"
}

func (p *Perro) Mover() int {
	c := p.Patas * 4
	return  c
}