package api

import (
	"time"

	"github.com/bryanbuiles/movie_suggester/internal/database"
	"github.com/bryanbuiles/movie_suggester/internal/logs"
)

// MovieFilter son los filtros de busqueda que se van a usar
type MovieFilter struct {
	Title    string `json:"title"` // asi es como va a quedar en el json
	Genre    string `json:"genre"`
	Director string `json:"director"`
}

// Movie class for Movie
type Movie struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Caste       string    `json:"caste"`
	ReleaseDate time.Time `json:"release_date"` // es una archivo de tiempo
	Genre       string    `json:"genre"`
	Director    string    `json:"director"`
}

// MovieSearch es una interfas que toma el metodo de la struct MovieService
type MovieSearch interface {
	Search(filter MovieFilter) ([]Movie, error) // retorn slice de estructuras
}

// MovieService ...
type MovieService struct {
	*database.PostgresSQL // cliente postgres
}

// Search Busca pelicuas, si tiene filtros los aplica si no arroja todas
func (s *MovieService) Search(filter MovieFilter) ([]Movie, error) { // metodo de MovieService
	// transaciones vas a ser segura en multitrading
	// rollback and commit cuando queramos, nos da la versatilidad de ir para atras y adelante
	tx, err := s.Begin()
	if err != nil {
		logs.Error("No se pudo crear transacion" + err.Error())
		return nil, err
	}

	rows, err := tx.Query(getMoviesQuery())

	if err != nil {
		logs.Error("No se pueden leer peliculas" + err.Error())
		_ = tx.Rollback() // rollback de la transacion
		return nil, err
	}
	var _movies []Movie
	for rows.Next() { // leee las filas de la tabla con scan
		var movie Movie
		err := rows.Scan(&movie.ID, &movie.Title, &movie.Caste, &movie.ReleaseDate, &movie.Genre, &movie.Director) // lee las columnas y las convierte datatype de go
		if err != nil {
			logs.Error("No se pueden leer peliculas" + err.Error())
			return nil, err
		}
		_movies = append(_movies, movie)
	}
	_ = tx.Commit()
	return _movies, nil
}
