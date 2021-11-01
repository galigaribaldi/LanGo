# LanGo
Repositorio para aprender a usar GO Golang, éste se encuentra dividido en 3 seccciones:  

* Básico:
    - Variables
    - Funciones
    - Ciclos
    - Condicionales
    - Estructuras
* Intermedio:
    - Paquetes
* Avanzado:
    - Hilos
    - Apuntadores
    - Conecciones a DB (Simples)
    - Creación de API's
    - ORM: conecciones con GORM (https://gorm.io/es_ES/docs/connecting_to_the_database.html)

En Go no existe Pulbico, privado, existe exportado o no exportado
https://qastack.mx/programming/17891226/difference-between-and-operators-in-go  
**Notes**
go run *file.go*
go build *file.go*
Para Documentacion: https://pkg.go.dev
Instalar Tour: go get golang.org/x/website/tour 
Instalar Echo: go get github.com/labstack/echo/v4
Quitar librerías ya no necesarias: go mod tidy
**Seguir*
Gophers Slack: gophers.slack.com
Go Tour: https://tour.golang.com/welcome/1
Golang Weekly: golangweekly.com
Play With Go: https://play-with-go.dev
Go By Example: https://gobyexample.com
Go Time: Spotify



// Inicializar un proyecto
go mod init path_del_proyecto

// Verificar que el código externo no esté corrupto
go mod verify

// Reemplazar fuente del código
go mod edit -replace path_del_repo_online=path_del_repo_en_local

// Quitar el replace
go mod edit -dropreplace path_del_repo_online

// Empaquetar todo el código de terceros que usa nuestro código
go mod vendor

// Eliminar todos los paquetes externos que no estemos usando
go mod tidy

// Aprender más de go modules
go help mod