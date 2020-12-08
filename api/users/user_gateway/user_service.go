package webusergateway

import (
	"fmt"

	"github.com/bryanbuiles/movie_suggester/api/users/models"
	"github.com/bryanbuiles/movie_suggester/internal/database"
	"github.com/bryanbuiles/movie_suggester/internal/logs"
	"github.com/gofiber/fiber/v2/utils"
)

// UserGateway ...
type UserGateway interface {
	SaveUser(cmd models.CreateUserCMD) (*models.UserInfo, error) // guardar usuario
	Login(cmd models.LoginCMD) string
	AddNextMovie(userID, movieID, comment string) error // add
}

// UserService conection to datebase
type UserService struct {
	DB *database.PostgresSQL
}

// SaveUser metodo para guardar usuarios
func (usuario *UserService) SaveUser(cmd models.CreateUserCMD) (*models.UserInfo, error) {
	if cmd.RepeatPassword != cmd.Password { // si las contraseñas no coiciden
		logs.Error("contraseñas no coiciden")
		return nil, fmt.Errorf("contraseñas no coiciden")
	}

	id := utils.UUID()
	tx, err := usuario.DB.Begin()

	if err != nil {
		logs.Error("Begin() fail at SaveUser " + err.Error())
		return nil, err
	}

	_, err = tx.Exec(CreateUserQuery(), id, cmd.Username, cmd.Password) // cmd.Username, cmd.Password son los values del exec

	if err != nil {
		logs.Error("Cannot insert user " + err.Error())
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &models.UserInfo{
		ID:       id,
		Username: cmd.Username,
		JWT:      "",
	}, nil
}

// Login return string for JWT
func (usuario *UserService) Login(cmd models.LoginCMD) string {
	var id string
	tx, err := usuario.DB.Begin()
	if err != nil {
		logs.Error("Fail Begin() at Login" + err.Error())
		return ""
	}
	row := tx.QueryRow(GetLoginQuerry(), cmd.Username, cmd.Password) // solo retorna una fila

	err = row.Scan(&id)
	if err != nil {
		tx.Rollback()
		logs.Error("Cannot querry login " + err.Error())
		return ""
	}
	tx.Commit()
	return id
}

// AddNextMovie Add a new movie to wish list
func (usuario *UserService) AddNextMovie(userID, movieID, comment string) error {
	tx, err := usuario.DB.Begin()
	if err != nil {
		logs.Error("Begin fail at whishmovies" + err.Error())
		return err
	}
	_, err = tx.Exec(SetWhishMovieQuery(), userID, movieID, comment)
	if err != nil {
		tx.Rollback()
		logs.Error("Cannot inssert a movie in wishlist" + err.Error())
		return err
	}
	tx.Commit()
	return nil
}
