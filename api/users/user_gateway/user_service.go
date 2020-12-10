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
	AddNextMovie(userID, movieID, comment string) error
	AllUsers() ([]models.User, error)
	GetUser(id string) (*models.User, error)
	DeleteUser(userID string) error
	UpdateUser(cmd models.User) (*models.User, error)
	GetWishMovies(userID string) ([]models.WishMovie, error)
	DeleteWishMovie(userID, movieID string) error
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
		logs.Error("Fail Begin() at Login " + err.Error())
		return ""
	}
	row := tx.QueryRow(GetLoginQuerry(), cmd.Username, cmd.Password) // solo retorna una fila

	err = row.Scan(&id)
	if err != nil {
		tx.Rollback()
		logs.Error("Cannot querry login Password or username incorrect " + err.Error())
		return ""
	}
	tx.Commit()
	return id
}

// AddNextMovie Add a new movie to wish list
func (usuario *UserService) AddNextMovie(userID, movieID, comment string) error {
	tx, err := usuario.DB.Begin()
	if err != nil {
		logs.Error("Begin fail at whishmovies " + err.Error())
		return err
	}
	_, err = tx.Exec(SetWhishMovieQuery(), userID, movieID, comment)
	if err != nil {
		tx.Rollback()
		logs.Error("Cannot insert a movie in wishlist " + err.Error())
		return err
	}
	tx.Commit()
	return nil
}

// AllUsers Get all user registred
func (usuario *UserService) AllUsers() ([]models.User, error) {
	tx, err := usuario.DB.Begin()
	if err != nil {
		logs.Error("Begin() fail at allUsers " + err.Error())
		return nil, err
	}
	rows, err := tx.Query(GetUsersQuery())
	if err != nil {
		logs.Error("Cannot read users " + err.Error())
		tx.Rollback() // rollback de la transacion
		return nil, err
	}
	var _users []models.User

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.Password)
		if err != nil {
			logs.Error("Cannot asign user to struct " + err.Error())
			return nil, err
		}
		_users = append(_users, user)
	}
	tx.Commit()
	return _users, nil

}

// GetUser get user by id
func (usuario *UserService) GetUser(userName string) (*models.User, error) {
	tx, err := usuario.DB.Begin()
	if err != nil {
		logs.Error("Begin() fail at GetUser " + err.Error())
		return nil, err
	}
	row := tx.QueryRow(getUserQuery(), userName)
	var user models.User
	err = row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		logs.Error("Fail to get user " + err.Error())
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &user, nil
}

// DeleteUser delete a user
func (usuario *UserService) DeleteUser(userID string) error {
	tx, err := usuario.DB.Begin()
	if err != nil {
		logs.Error("Begin fail at Delete user " + err.Error())
		return err
	}
	_, err = tx.Exec(deleteUserQuery(), userID)
	if err != nil {
		logs.Error("error deleting user " + err.Error())
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// UpdateUser update pass o username to User
func (usuario *UserService) UpdateUser(cmd models.User) (*models.User, error) {
	tx, err := usuario.DB.Begin()
	if err != nil {
		logs.Error("Begin fail at Update User " + err.Error())
		return nil, err
	}
	row := tx.QueryRow(getUsersQuerybyID(), cmd.ID)
	var user models.User
	err = row.Scan(&user.Username, &user.Password)
	if err != nil {
		logs.Error("Fail to get user " + err.Error())
		tx.Rollback()
		return nil, err
	}
	user.ID = cmd.ID
	if cmd.Username != "" {
		user.Username = cmd.Username
	}
	if cmd.Password != "" {
		user.Password = cmd.Password
	}
	_, err = tx.Exec(updateUserQuery(), user.ID, user.Username, user.Password)
	if err != nil {
		logs.Error("Error updating movie " + err.Error())
		logs.Info(updateUserQuery())
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &user, nil
}

// GetWishMovies Get a list of whish Movies
func (usuario *UserService) GetWishMovies(userID string) ([]models.WishMovie, error) {
	tx, err := usuario.DB.Begin()
	if err != nil {
		logs.Error("Begin() fail at GetWishMovies " + err.Error())
		return nil, err
	}
	rows, err := tx.Query(getWhishMoviesQuery(), userID)
	if err != nil {
		logs.Error("Cannot read Wish Movies " + err.Error())
		tx.Rollback() // rollback de la transacion
		return nil, err
	}
	var _wishMovies []models.WishMovie
	var wishMovies models.WishMovie
	for rows.Next() {
		err := rows.Scan(&wishMovies.MovieID, &wishMovies.Title, &wishMovies.Caste,
			&wishMovies.ReleaseDate, &wishMovies.Genre, &wishMovies.Director, &wishMovies.Comment)
		if err != nil {
			logs.Error("Cannot asign wish movie to struct " + err.Error())
			return nil, err
		}
		_wishMovies = append(_wishMovies, wishMovies)
	}
	tx.Commit()
	return _wishMovies, nil
}

// DeleteWishMovie to delete a wish Movie
func (usuario *UserService) DeleteWishMovie(userID, movieID string) error {
	tx, err := usuario.DB.Begin()
	if err != nil {
		logs.Error("Begin() fail at Delete wish movie " + err.Error())
		return err
	}
	_, err = tx.Exec(deleteWishMovieQuery(), userID, movieID)
	if err != nil {
		logs.Error("Error deleting wish movie " + err.Error())
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
