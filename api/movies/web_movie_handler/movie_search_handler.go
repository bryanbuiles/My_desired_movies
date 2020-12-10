// Here is the file for handlers

package webhandler

import (
	"github.com/bryanbuiles/movie_suggester/api/movies/models"
	"github.com/gofiber/fiber/v2"
)

// SearchMovieHandler handler of get all movies
func (w *WebServices) SearchMovieHandler(c *fiber.Ctx) error {
	// esta llamando a weeb services que es una struct que contiene Services que es otra struct
	// que contiene la interfas MovieSearch que tiene el metodo Search que recibe la struct Moviefilter que tiene los filtros
	// para buscar, si no tiene filtros busca todos
	res, err := w.search.Search(models.MovieFilter{ // Look this struct
		Title:    c.Query("title"), // Query() is a method of fiber diferet of sql
		Genre:    c.Query("genre"),
		Director: c.Query("director"),
	})

	if err != nil {
		return fiber.NewError(400, "No movies found") // se le esta enviando el error a fiber CTX
	}

	if len(res) == 0 {
		return c.JSON([]interface{}{})
	}

	return c.JSON(res)
}

// CreateMovieHandler handler to add new movie
func (w *WebServices) CreateMovieHandler(ctx *fiber.Ctx) error {
	var cmd models.Movie
	err := ctx.BodyParser(&cmd)
	if err != nil {
		return fiber.NewError(400, "BodyParser Fail at CreateMovieHandler")
	}
	res, err := w.search.CreateMovie(cmd)
	if err != nil {
		return fiber.NewError(400, "Create Movie fail")
	}
	return ctx.JSON(res)
}

// DeleteMovieHandler delete movie Handler
func (w *WebServices) DeleteMovieHandler(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := w.search.DeleteMovie(id)
	if err != nil {
		return fiber.NewError(400, "Connot Delete Movie")
	}
	return ctx.JSON(fiber.Map{"status": "success", "message": "Movie successfully deleted", "data": nil})
}

// UpdateMovieHandler update movie handler
func (w *WebServices) UpdateMovieHandler(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var cmd models.Movie
	cmd.ID = id
	err := ctx.BodyParser(&cmd)
	if err != nil {
		return fiber.NewError(400, "BodyParser Fail at UpdateMovieHandler")
	}
	err = w.search.UpdateMovie(cmd)
	if err != nil {
		return fiber.NewError(400, "Cannot update Movie")
	}
	return ctx.JSON(fiber.Map{"status": "success", "message": "Movie successfully Update", "data": nil})
}
