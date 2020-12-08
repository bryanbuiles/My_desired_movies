package webusergateway

// esta es la funcion para filtrar la busqueda de las peliculas por director, title or genre

// CreateUserQuery ...
func CreateUserQuery() string {
	return "insert into users (id, username, password) values ($1, $2, $3)"
}

// GetLoginQuerry query to bring the id
func GetLoginQuerry() string {
	return "SELECT id FROM users WHERE username = $1 and password = $2"
}

// SetWhishMovieQuery query to wishlists
func SetWhishMovieQuery() string {
	return "INSERT INTO wish_list (user_id, movie_id, comment) VALUES ($1, $2, $3)"
}
