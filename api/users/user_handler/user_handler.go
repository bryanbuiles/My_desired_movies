package userhandler

import (
	"github.com/bryanbuiles/movie_suggester/api/users/models"
	"github.com/bryanbuiles/movie_suggester/internal/logs"
	"github.com/gofiber/fiber/v2"
)

// CreateUserHandler ...
func (w *WebServices) CreateUserHandler(ctx *fiber.Ctx) error {
	var cmd models.CreateUserCMD
	err := ctx.BodyParser(&cmd) // take the elements of body request

	res, err := w.Services.users.SaveUser(cmd)

	if err != nil {
		return fiber.NewError(400, "Create user fail")
	}
	t := signToken(w.tokenKey, res.ID)
	res.JWT = t //asigning the token to userinfo struct
	return ctx.JSON(res)
}

// WhishMoviesHandler handler to wish movie list
func (w *WebServices) WhishMoviesHandler(ctx *fiber.Ctx) error {
	var cmd models.WishMovieCMD
	ctx.BodyParser(&cmd)
	bearer := ctx.Get("Authorization")
	userID := extractUserIDFromJWT(bearer, w.tokenKey)
	err := w.users.AddNextMovie(userID, cmd.MovieID, cmd.Comment)
	if err != nil {
		return fiber.NewError(400, "cannot add movie to whish list")
	}
	return ctx.JSON(struct { // Anonymous struct
		Result string `json:"result"`
	}{
		Result: "Movie added to wish list",
	})
}

// SetupVideo upload a video in static way in the endpoint
func (w *WebServices) SetupVideo(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "video/mp4") // set the videos in the response header
	err := ctx.SendFile("/home/kinaret/movie_suggester/test.mp4", false)
	if err != nil {
		logs.Error("display video fail" + err.Error())
		return fiber.NewError(400, "Display video fail")
	}
	return nil
}

// LoginHandler handler to login
func (w *WebServices) LoginHandler(ctx *fiber.Ctx) error {
	var cmdLogin models.LoginCMD
	err := ctx.BodyParser(&cmdLogin)
	if err != nil {
		return fiber.NewError(400, "Login fail, cannot parse params")
	}
	id := w.users.Login(cmdLogin)
	if id == "" {
		return fiber.NewError(400, "user not found or id wrong")
	}

	type resLogHandler struct {
		Token string `json:"token"`
	}

	return ctx.JSON(resLogHandler{
		Token: signToken(w.tokenKey, id),
	})
}

// GetUsersHandler habdler for gey all users
func (w *WebServices) GetUsersHandler(ctx *fiber.Ctx) error {
	res, err := w.users.AllUsers()
	if err != nil {
		return fiber.NewError(400, "cannot display users")
	}
	if len(res) == 0 {
		return ctx.JSON([]interface{}{})
	}
	return ctx.JSON(res)
}

// GetUserByIDHandler get user by id
func (w *WebServices) GetUserByIDHandler(ctx *fiber.Ctx) error {
	userName := ctx.Params("username")
	res, err := w.users.GetUser(userName)
	if err != nil {
		return fiber.NewError(400, "User not exist")
	}
	return ctx.JSON(res)
}

// DeleteUserHandler by token
func (w *WebServices) DeleteUserHandler(ctx *fiber.Ctx) error {
	bearer := ctx.Get("Authorization")
	userID := extractUserIDFromJWT(bearer, w.tokenKey)
	err := w.users.DeleteUser(userID)
	if err != nil {
		return fiber.NewError(400, "Delete a user fail")
	}
	return ctx.JSON(fiber.Map{"status": "success", "message": "User successfully deleted", "data": nil})
}

// UpdateUserHandler handler to update user
func (w *WebServices) UpdateUserHandler(ctx *fiber.Ctx) error {
	bearer := ctx.Get("Authorization")
	userID := extractUserIDFromJWT(bearer, w.tokenKey)
	var user models.User
	err := ctx.BodyParser(&user)
	user.ID = userID
	if err != nil {
		return fiber.NewError(400, "BodyParser Fail at UpdateUserHandler")
	}
	res, err := w.users.UpdateUser(user)
	if err != nil {
		return fiber.NewError(400, "Update a user fail")
	}
	return ctx.JSON(res)
}

// GetwishedListHandler handler to get wish movies
func (w *WebServices) GetwishedListHandler(ctx *fiber.Ctx) error {
	bearer := ctx.Get("Authorization")
	userID := extractUserIDFromJWT(bearer, w.tokenKey)
	res, err := w.users.GetWishedMovies(userID)
	if err != nil {
		return fiber.NewError(400, "Get a list of wish movies fail")
	}
	if len(res) == 0 {
		return ctx.JSON([]interface{}{})
	}
	return ctx.JSON(res)
}

// DeleteWishedMovieHandler handler to delete wish movie
func (w *WebServices) DeleteWishedMovieHandler(ctx *fiber.Ctx) error {
	bearer := ctx.Get("Authorization")
	movieID := ctx.Params("movieID")
	userID := extractUserIDFromJWT(bearer, w.tokenKey)
	err := w.users.DeleteWishedMovie(userID, movieID)
	if err != nil {
		return fiber.NewError(400, "Delete a wish movie fail")
	}
	return ctx.JSON(fiber.Map{"status": "success", "message": "Wish movie successfully deleted", "data": nil})
}
