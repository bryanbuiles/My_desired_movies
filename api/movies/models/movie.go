package models

// MovieFilter son los filtros de busqueda que se van a usar
type MovieFilter struct {
	Title    string `json:"title"` // asi es como va a quedar en el json
	Genre    string `json:"genre"`
	Director string `json:"director"`
}

// Movie class for Movie
type Movie struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Caste       string `json:"caste"`
	ReleaseDate string `json:"release_date"` // es una archivo de tiempo
	Genre       string `json:"genre"`
	Director    string `json:"director"`
}
