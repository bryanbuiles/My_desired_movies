package api

import (
	"fmt"

	"github.com/bryanbuiles/movie_suggester/internal/database"
	"github.com/bryanbuiles/movie_suggester/internal/logs"
	"github.com/gofiber/fiber/v2/utils"
)

// CreateUserCMD struct para la creacion de usuarios
type CreateUserCMD struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeat_password"`
}

// UserGateway ...
type UserGateway interface {
	SaveUser(cmd CreateUserCMD) (*UserInfo, error) // guardar usuario
	Login()
}

// UserService conection to datebase
type UserService struct {
	*database.PostgresSQL
}

// UserInfo ...
type UserInfo struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	JWT      string `json:"token"`
}

// Login ...
func (usuario *UserService) Login() {

}

// SaveUser metodo para guardar usuarios
func (usuario *UserService) SaveUser(cmd CreateUserCMD) (*UserInfo, error) {
	id := utils.UUID()

	if cmd.RepeatPassword != cmd.Password { // si las contraseñas no coiciden
		logs.Error("contraseñas no coiciden")
		return nil, fmt.Errorf("contraseñas no coiciden")
	}
	_, err := usuario.Exec(CreateUserQuery(), id, cmd.Username, cmd.Password) // cmd.Username, cmd.Password son los values del exec

	if err != nil {
		logs.Error("Cannot insert user " + err.Error())
		return nil, err
	}
	return &UserInfo{
		ID:       id,
		Username: cmd.Username,
		JWT:      "",
	}, nil
}
