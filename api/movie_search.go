package api

import "time"

// MovieFilter son los filtros de busqueda que se van a usar
type MovieFilter struct {
	Title    string `json:"title"`
	Genre    string `json:"genre"`
	Director string `json:"director"`
}

// Movie class for Movie
type Movie struct {
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
}

// Search Busca pelicuas, si tiene filtros los aplica si no arroja todas
func (s *MovieService) Search(filter MovieFilter) ([]Movie, error) { // metodo de MovieService
	movie1 := Movie{
		Title:       "Blade Runner",
		Caste:       "Harrison Ford",
		ReleaseDate: time.Now(),
		Genre:       "Cs Fiction",
		Director:    "",
	}
	movie2 := Movie{
		Title:       "Driver",
		Caste:       "Ryan Gosling",
		ReleaseDate: time.Now(),
		Genre:       "Drama",
		Director:    "",
	}
	var _movies []Movie
	_movies = append(_movies, movie1)
	_movies = append(_movies, movie2)
	return _movies, nil
}
