package models

// UserInfo display userinfo token
type UserInfo struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	JWT      string `json:"token"`
}

// CreateUserCMD struct to create new users
type CreateUserCMD struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeat_password"`
}

// LoginCMD struct to login
type LoginCMD struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// WishMovieCMD struct for add a wish movie
type WishMovieCMD struct {
	MovieID string `json:"movie_id"`
	Comment string `json:"comment"`
}

// User table to get user info
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// WishMovie class for display the wishes movies for a user
type WishMovie struct {
	MovieID     string `json:"movie_id"`
	Title       string `json:"title"`
	Caste       string `json:"caste"`
	ReleaseDate string `json:"release_date"`
	Genre       string `json:"genre"`
	Director    string `json:"director"`
	Comment     string `json:"comment"`
}
