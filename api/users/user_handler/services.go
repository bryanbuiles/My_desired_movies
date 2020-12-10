package userhandler

import (
	webusergateway "github.com/bryanbuiles/movie_suggester/api/users/user_gateway"
	"github.com/bryanbuiles/movie_suggester/internal/database"
)

// Services ...
type Services struct {
	users webusergateway.UserGateway
}

// WebServices .for the users
type WebServices struct {
	Services
	tokenKey string // The token is passed through the start() function
}

// NewServices New user servicio
func NewServices() Services {
	client := database.NewPostgresSQLClient()
	return Services{
		users: &webusergateway.UserService{DB: client},
	}
}

// Start a new user service
func Start(tokenKey string) *WebServices {
	return &WebServices{NewServices(), tokenKey}
}
