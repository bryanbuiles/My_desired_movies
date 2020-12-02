package api

func getMoviesQuery() string {
	return "select id, title, caste, release_date, genre, director from movie"
}
