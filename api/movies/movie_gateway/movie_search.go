package moviesgateway

import (
	"github.com/bryanbuiles/movie_suggester/api/movies/models"
	"github.com/bryanbuiles/movie_suggester/internal/database"
	"github.com/bryanbuiles/movie_suggester/internal/logs"
)

// MovieSearch es una interfas que toma el metodo de la struct MovieService
type MovieSearch interface {
	Search(filter models.MovieFilter) ([]models.Movie, error) // retorn slice de estructuras
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
