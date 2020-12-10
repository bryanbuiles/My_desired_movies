package moviesgateway

import (
	"github.com/bryanbuiles/movie_suggester/api/movies/models"
	"github.com/bryanbuiles/movie_suggester/internal/database"
	"github.com/bryanbuiles/movie_suggester/internal/logs"
	"github.com/gofiber/fiber/v2/utils"
)

// MovieSearch is an interface that takes methods of struct MovieService
type MovieSearch interface {
	Search(filter models.MovieFilter) ([]models.Movie, error)
	CreateMovie(cmd models.Movie) (*models.Movie, error)
	DeleteMovie(movieID string) error
	UpdateMovie(cmd models.Movie) error
}

// MovieService .create a Db client
type MovieService struct {
	DB *database.PostgresSQL // client postgres
}

// Search movies to get all movies or search movies with filters
func (s *MovieService) Search(filter models.MovieFilter) ([]models.Movie, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		logs.Error("Begin() fail at Search() " + err.Error())
		return nil, err
	}

	rows, err := tx.Query(getMoviesQuery(filter))

	if err != nil {
		logs.Error("Cannot display movies " + err.Error())
		tx.Rollback() // rollback the query
		return nil, err
	}
	var _movies []models.Movie
	for rows.Next() { // Next() return true if the next row is readed
		var movie models.Movie
		// Scan() read the columns in each row and assigns the values of columns to variables or structs parameters of go
		err := rows.Scan(&movie.ID, &movie.Title, &movie.Caste, &movie.ReleaseDate, &movie.Genre, &movie.Director)
		if err != nil {
			logs.Error("Cannot display movies " + err.Error())
			return nil, err
		}
		_movies = append(_movies, movie)
	}
	tx.Commit()
	return _movies, nil
}

// CreateMovie add a new movie to DB
func (s *MovieService) CreateMovie(cmd models.Movie) (*models.Movie, error) {
	id := utils.UUID()

	tx, err := s.DB.Begin()

	if err != nil {
		logs.Error("Begin() fail at Create movie " + err.Error())
		return nil, err
	}
	_, err = tx.Exec(CreateMovieQuery(), id, cmd.Title, cmd.Caste, cmd.ReleaseDate, cmd.Genre, cmd.Director)

	if err != nil {
		logs.Error("Cannot create a movie " + err.Error())
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &models.Movie{
		ID:          id,
		Title:       cmd.Title,
		Caste:       cmd.Caste,
		ReleaseDate: cmd.ReleaseDate,
		Genre:       cmd.Genre,
		Director:    cmd.Director,
	}, nil
}

// DeleteMovie to delete a movie
func (s *MovieService) DeleteMovie(movieID string) error {
	tx, err := s.DB.Begin()

	if err != nil {
		logs.Error("Begin() fail at Delete movie " + err.Error())
		return err
	}
	_, err = tx.Exec(DeleteMovieQuery(), movieID)
	if err != nil {
		logs.Error("error to delete movie " + err.Error())
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// UpdateMovie update a movie value
func (s *MovieService) UpdateMovie(cmd models.Movie) error {
	tx, err := s.DB.Begin()
	if err != nil {
		logs.Error("Begin() fail at Update movie " + err.Error())
		return err
	}
	_, err = tx.Exec(UpdateMovieQuery(cmd))
	if err != nil {
		logs.Error("Error updating movie " + err.Error())
		logs.Info(UpdateMovieQuery(cmd))
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
