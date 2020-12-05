package api

import (
	"github.com/gofiber/fiber/v2"
)

// SetupMoviesRoutes ruta para todas las movies
func SetupMoviesRoutes(app *fiber.App) { // app de fiber como parametro
	star := start()
	grp := app.Group("/movies")

	grp.Get("/", star.SearchMovieHandler)
}

// SetupUserRoutes ...
func SetupUserRoutes(app *fiber.App) {
	star := start()
	group := app.Group("/users")

	group.Post("/", star.CreateUserHandler)
}
