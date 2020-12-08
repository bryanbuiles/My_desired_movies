package userhandler

import (
	webusergateway "github.com/bryanbuiles/movie_suggester/api/users/user_gateway"
	"github.com/bryanbuiles/movie_suggester/internal/database"
)

// Services struct que lista los diferentes servicios
// son los servicios que va a tener el programa
type Services struct { // get movies
	users webusergateway.UserGateway // interfas save user
}

// WebServices servicios web
type WebServices struct {
	Services
	tokenKey string // se le pasa el token desde start() en las rutas
}

// NewServices Nuevo servicio
func NewServices() Services {
	client := database.NewPostgresSQLClient()
	return Services{ // Search() es un metodo de Movie service
		users: &webusergateway.UserService{DB: client},
	}
}

// Start comienza un nuevo servicio
func Start(tokenKey string) *WebServices { // comieza el servicio
	return &WebServices{NewServices(), tokenKey}
}
