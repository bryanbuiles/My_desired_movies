package api

import (
	"github.com/gofiber/fiber/v2"
)

// SetupMoviesRoutes ruta para todas las movies
func SetupMoviesRoutes(app *fiber.App) {
	star := start()
	grp := app.Group("/movies")

	grp.Get("/", star.SearchMovieHandler)
}
