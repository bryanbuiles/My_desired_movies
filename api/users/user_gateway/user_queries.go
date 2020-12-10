package webusergateway

// CreateUserQuery query to create an user
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

// GetUsersQuery querry to get all users
func GetUsersQuery() string {
	return "SELECT id, username, password FROM users"
}

func getUserQuery() string {
	return "SELECT id, username, password FROM users WHERE username = $1"
}

// DeleteMovieQuery To delete a Movie
func deleteUserQuery() string {
	return "DELETE FROM users WHERE id = $1"
}

func updateUserQuery() string {
	return "UPDATE users SET username = $2, password = $3 WHERE id = $1"
}

// GetUsersQuery querry to get user by id
func getUsersQuerybyID() string {
	return "SELECT username, password FROM users WHERE id = $1"
}

// getWhishMoviesQuery query to display the movies included in the wishes list movie of an user
func getWhishMoviesQuery() string {
	return "SELECT wish_list.movie_id, title, caste, release_date, genre, director, wish_list.comment " +
		"FROM movie INNER JOIN wish_list ON movie.id = wish_list.movie_id WHERE wish_list.user_id = $1"
}

// DeleteMovieQuery To delete a Movie
func deleteWishMovieQuery() string {
	return "DELETE FROM wish_list WHERE user_id = $1 and movie_id = $2"
}
