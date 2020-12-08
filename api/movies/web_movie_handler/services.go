package webhandler

import (
	moviesgateway "github.com/bryanbuiles/movie_suggester/api/movies/movie_gateway"
	"github.com/bryanbuiles/movie_suggester/internal/database"
)

// Services struct que lista los diferentes servicios
// son los servicios que va a tener el programa
type Services struct {
	search moviesgateway.MovieSearch // get movies
}

// WebServices servicios web
type WebServices struct {
	Services
}

// NewServices Nuevo servicio
func NewServices() Services {
	client := database.NewPostgresSQLClient()
	return Services{
		search: &moviesgateway.MovieService{S: client}, // Search() es un metodo de Movie service
	}
}

// Start comienza un nuevo servicio
func Start() *WebServices { // comieza el servicio
	return &WebServices{NewServices()}
}
