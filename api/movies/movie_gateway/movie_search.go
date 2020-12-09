package moviesgateway

import (
	"github.com/bryanbuiles/movie_suggester/api/movies/models"
	"github.com/bryanbuiles/movie_suggester/internal/database"
	"github.com/bryanbuiles/movie_suggester/internal/logs"
	"github.com/gofiber/fiber/v2/utils"
)

// MovieSearch es una interfas que toma el metodo de la struct MovieService
type MovieSearch interface {
	Search(filter models.MovieFilter) ([]models.Movie, error) // retorn slice de estructuras
	CreateMovie(cmd models.Movie) (*models.Movie, error)
	DeleteMovie(movieID string) error
	UpdateMovie(cmd models.Movie) error
}

// MovieService ...
type MovieService struct {
	S *database.PostgresSQL // cliente postgres
}

// Search Busca pelicuas, si tiene filtros los aplica si no arroja todas
func (s *MovieService) Search(filter models.MovieFilter) ([]models.Movie, error) { // metodo de MovieService
	// transaciones vas a ser segura en multitrading
	// rollback and commit cuando queramos, nos da la versatilidad de ir para atras y adelante
	tx, err := s.S.Begin()
	if err != nil {
		logs.Error("No se pudo crear transacion" + err.Error())
		return nil, err
	}

	rows, err := tx.Query(getMoviesQuery(filter))

	if err != nil {
		logs.Error("No se pueden leer peliculas" + err.Error())
		tx.Rollback() // rollback de la transacion
		return nil, err
	}
	var _movies []models.Movie
	// recorriendo las filas
	for rows.Next() { // Next() devuelve true si la siguiente fila esta preparada o leida
		var movie models.Movie
		// Scan() lee las columnas en cada fila y las asigna los valores de la db a las variables de go o una struct
		err := rows.Scan(&movie.ID, &movie.Title, &movie.Caste, &movie.ReleaseDate, &movie.Genre, &movie.Director)
		if err != nil {
			logs.Error("No se pueden leer peliculas" + err.Error())
			return nil, err
		}
		_movies = append(_movies, movie) // lista que se retorna con todas las movies buscadas
	}
	tx.Commit()
	return _movies, nil
}

// CreateMovie add a new movie to DB
func (s *MovieService) CreateMovie(cmd models.Movie) (*models.Movie, error) {
	id := utils.UUID()

	tx, err := s.S.Begin()

	if err != nil {
		logs.Error("Begin fail at Create movie" + err.Error())
		return nil, err
	}
	_, err = tx.Exec(CreateMovieQuery(), id, cmd.Title, cmd.Caste, cmd.ReleaseDate, cmd.Genre, cmd.Director)

	if err != nil {
		logs.Error("Cannot create a movie" + err.Error())
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
	tx, err := s.S.Begin()

	if err != nil {
		logs.Error("Begin fail at Delete movie" + err.Error())
		return err
	}
	_, err = tx.Exec(DeleteMovieQuery(), movieID)
	if err != nil {
		logs.Error("error to delete movie" + err.Error())
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// UpdateMovie update a movie value
func (s *MovieService) UpdateMovie(cmd models.Movie) error {
	tx, err := s.S.Begin()
	if err != nil {
		logs.Error("Begin fail at Update movie" + err.Error())
		return err
	}
	_, err = tx.Exec(UpdateMovieQuery(cmd))
	if err != nil {
		logs.Error("Error updating movie" + err.Error())
		logs.Info(UpdateMovieQuery(cmd))
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
