// Remember, before this execute *go mod init*
package mypackage

// CarPublic Car con acceso publico
type CarPublic struct {
	Brand string
	Year  int
}

// CarPruvate Car con acceso privado
type carPrivate struct {
	Brand string
	Year  int
}

// PrintMessage Imprimir un mensaje
func PrintMessage() string {
	return "hola"
}
