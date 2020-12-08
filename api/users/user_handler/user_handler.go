package userhandler

import (
	"github.com/bryanbuiles/movie_suggester/api/users/models"
	"github.com/bryanbuiles/movie_suggester/internal/logs"
	"github.com/gofiber/fiber/v2"
)

// CreateUserHandler ...
func (w *WebServices) CreateUserHandler(ctx *fiber.Ctx) error {
	var cmd models.CreateUserCMD
	err := ctx.BodyParser(&cmd) // toma los elemento que llegan en el endpoint

	res, err := w.Services.users.SaveUser(cmd)

	if err != nil {
		return fiber.NewError(400, "Create user fail")
	}
	t := signToken(w.tokenKey, res.ID)
	res.JWT = t // pasandolo a la struct userinfo
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
	return ctx.JSON(struct { // struct anonima
		Result string `json:"result"`
	}{
		Result: "Movie added to wish list",
	})
}

// SetupVideo sube un video de manera estatica al endpoint movies
func (w *WebServices) SetupVideo(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "video/mp4")            // setea al header un contenido
	err := ctx.SendFile("../../../test.mp4", false) // se maneja de manera estatica
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
		return fiber.NewError(400, "user not found")
	}

	type resLogHandler struct {
		Token string `json:"token"`
	}

	return ctx.JSON(resLogHandler{
		Token: signToken(w.tokenKey, id),
	})
}
