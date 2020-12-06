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
	Login(cmd LoginCMD) string
	AddNextMovie(userID, movieID, comment string) error // add whishlist
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

// Login return string for JWT
func (usuario *UserService) Login(cmd LoginCMD) string {
	var id string
	row := usuario.QueryRow(GetLoginQuerry(), cmd.Username, cmd.Password) // solo retorna una fila
	err := row.Scan(&id)
	if err != nil {
		logs.Error("Cannot querry login " + err.Error())
		return ""
	}
	return id
}

// AddNextMovie Add a new movie to wish list
func (usuario *UserService) AddNextMovie(userID, movieID, comment string) error {
	_, err := usuario.Exec(SetWhishMovieQuery(), userID, movieID, comment)
	if err != nil {
		logs.Error("Cannot inssert a movie in wishlist" + err.Error())
		return err
	}
	return nil
}
