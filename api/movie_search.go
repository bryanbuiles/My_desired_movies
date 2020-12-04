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

	rows, err := tx.Query(getMoviesQuery(filter))

	if err != nil {
		logs.Error("No se pueden leer peliculas" + err.Error())
		_ = tx.Rollback() // rollback de la transacion
		return nil, err
	}
	var _movies []Movie
	// recorriendo las filas
	for rows.Next() { // Next() devuelve true si la siguiente fila esta preparada o leida
		var movie Movie
		// Scan() lee las columnas en cada fila y las asigna los valores de la db a las variables de go o una struct
		err := rows.Scan(&movie.ID, &movie.Title, &movie.Caste, &movie.ReleaseDate, &movie.Genre, &movie.Director)
		if err != nil {
			logs.Error("No se pueden leer peliculas" + err.Error())
			return nil, err
		}
		_movies = append(_movies, movie) // lista que se retorna con todas las movies buscadas
	}
	_ = tx.Commit()
	return _movies, nil
}
