package router

import (
	webhandler "github.com/bryanbuiles/movie_suggester/api/movies/web_movie_handler"
	userhandler "github.com/bryanbuiles/movie_suggester/api/users/user_handler"
	"github.com/gofiber/fiber/v2"
)

// SetupMoviesRoutes ruta para todas las movies
func SetupMoviesRoutes(app *fiber.App, tokenKey string) { // app de fiber como parametro
	star := webhandler.Start()
	grp := app.Group("/movies")

	grp.Get("/", star.SearchMovieHandler)
}

// SetupUserRoutes ...
func SetupUserRoutes(app *fiber.App, tokenKey string) {
	star := userhandler.Start(tokenKey)
	group := app.Group("/users")

	group.Get("/videos", star.SetupVideo)
	group.Post("/", star.CreateUserHandler)
	group.Post("/login", star.LoginHandler)

	group.Use(userhandler.JwtMiddleware(tokenKey)).Post("/wishlist", star.WhishMoviesHandler)
}
