package api

import (
	"github.com/bryanbuiles/movie_suggester/internal/database"
	"github.com/bryanbuiles/movie_suggester/internal/logs"
)

// CreateUser struct para la creacion de usuarios
type CreateUser struct {
	Username       string `json: "username"`
	Password       string `json: "password"`
	RepeatPassword string `json: "Repeatpassword"`
}

// UserService ...
type userService interface {
	SaveUser(cmd CreateUser) (string, error) // guardar usuario
	Login()
}

// UserService conection to datebase
type UserService struct {
	*database.PostgresSQL
}

// Login ...
func (usuario *UserService) Login() {

}

// SaveUser metodo para guardar usuarios
func (usuario *UserService) SaveUser(cmd CreateUser) (string, error) {
	_, err := usuario.Exec(createUserQuery(), cmd.Username, cmd.Password) // cmd.Username, cmd.Password son los values del exec

	if err != nil {
		logs.Error("Cannot insert user " + err.Error())
		return "", err
	}
	return cmd.Username, nil
}
