package models

// UserInfo ...
type UserInfo struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	JWT      string `json:"token"`
}

// CreateUserCMD struct para la creacion de usuarios
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

// WishMovieCMD struct for wishlist
type WishMovieCMD struct {
	MovieID string `json:"movie_id"`
	Comment string `json:"comment"`
}

// User table
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
