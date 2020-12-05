package api

import "github.com/bryanbuiles/movie_suggester/internal/database"

// Services struct que lista los diferentes servicios
// son los servicios que va a tener el programa
type Services struct {
	search MovieSearch
	users  UserGateway
}

// WebServices servicios web
type WebServices struct {
	Services
}

// NewServices Nuevo servicio
func NewServices() Services {
	client := database.NewPostgresSQLClient()
	return Services{
		search: &MovieService{client}, // Search() es un metodo de Movie service
		users:  &UserService{client},
	}
}

// Start comienza un nuevo servicio
func start() *WebServices { // comieza el servicio
	return &WebServices{NewServices()}
}
