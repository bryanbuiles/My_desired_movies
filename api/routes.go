package api

import (
	"github.com/gofiber/fiber/v2"
)

// SetupMoviesRoutes ruta para todas las movies
func SetupMoviesRoutes(app *fiber.App, tokenKey string) { // app de fiber como parametro
	star := start(tokenKey)
	grp := app.Group("/movies")

	grp.Get("/", star.SearchMovieHandler)
}

// SetupUserRoutes ...
func SetupUserRoutes(app *fiber.App, tokenKey string) {
	star := start(tokenKey)
	group := app.Group("/users")

	group.Get("/videos", star.SetupVideo)
	group.Get("/wishlist", star.WhishMoviesHandler)
	group.Post("/", star.CreateUserHandler)
	group.Post("/login", star.LoginHandler)
}
