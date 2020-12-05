package api

import "github.com/bryanbuiles/movie_suggester/internal/database"

// Services struct que lista los diferentes servicios
// son los servicios que va a tener el programa
type Services struct {
	search MovieSearch // get movies
	users  UserGateway // interfas save user
}

// WebServices servicios web
type WebServices struct {
	Services
	tokenKey string // se le pasa el token desde start() en las rutas
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
func start(tokenKey string) *WebServices { // comieza el servicio
	return &WebServices{NewServices(), tokenKey}
}
