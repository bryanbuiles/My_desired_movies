package models

// MovieFilter are the parameters that works as a filter to search a movie
type MovieFilter struct {
	Title    string `json:"title"` // json format in output
	Genre    string `json:"genre"`
	Director string `json:"director"`
}

// Movie class for Movie
type Movie struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Caste       string `json:"caste"`
	ReleaseDate string `json:"release_date"`
	Genre       string `json:"genre"`
	Director    string `json:"director"`
}
