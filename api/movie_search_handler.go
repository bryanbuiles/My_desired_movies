// Here is the file for handlers

package api

import "github.com/gofiber/fiber/v2"

// SearchMovieHandler  este es el handler
func (w *WebServices) SearchMovieHandler(c *fiber.Ctx) error {
	// esta llamando a weeb services que es una struct que contiene S Services que es otra struct
	// que contiene la interfas MovieSearch que tiene el metodo Search que recibe la struct Moviefilter que tiene los filtros
	// para buscar, si no tiene filtros busca todos
	res, err := w.search.Search(MovieFilter{}) // asi se pasa una struct

	if err != nil {
		return fiber.NewError(400, "No movies found") // se le esta enviando el error a fiber CTX
	}
	return c.JSON(res)
}
