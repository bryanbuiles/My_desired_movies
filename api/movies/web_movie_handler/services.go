package webhandler

import (
	moviesgateway "github.com/bryanbuiles/movie_suggester/api/movies/movie_gateway"
	"github.com/bryanbuiles/movie_suggester/internal/database"
)

// Services struct list all services taht will have movies
// is a struct of movieSearch interface
type Services struct {
	search moviesgateway.MovieSearch
}

// WebServices ...
type WebServices struct {
	Services
}

// NewServices New service
func NewServices() Services {
	client := database.NewPostgresSQLClient()
	return Services{
		search: &moviesgateway.MovieService{DB: client}, // Search() is a method of Movie service
	}
}

// Start starts a new service
func Start() *WebServices {
	return &WebServices{NewServices()}
}
