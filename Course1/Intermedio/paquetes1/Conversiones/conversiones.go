package Conversiones
import "fmt"

func init(){
	fmt.Println("Paquete Conversiones")
}
/// Metodo Exportado -> iniciando con mayuscula (Publica)
func MetterToIn(valor float32) float32{
	c := valor*1.2
	metodo1()
	return c
}
///Metodo no exportado -> minuscula (privada)
func mettertokm(valor float32) float32{
	c := valor*100
	
	return c
}