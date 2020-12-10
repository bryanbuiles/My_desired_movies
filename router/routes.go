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
	grp.Post("/", star.CreateMovieHandler)
	grp.Delete("/:id", star.DeleteMovieHandler)
	grp.Patch("/:id", star.UpdateMovieHandler)
}

// SetupUserRoutes ...
func SetupUserRoutes(app *fiber.App, tokenKey string) {
	star := userhandler.Start(tokenKey)
	group := app.Group("/users")

	group.Get("/", star.GetUsersHandler)
	group.Get("/videos", star.SetupVideo) // dont move to above, conflicts wih /:username
	group.Get("/:username", star.GetUserByIDHandler)
	group.Post("/", star.CreateUserHandler)
	group.Post("/login", star.LoginHandler)

	group.Use(userhandler.JwtMiddleware(tokenKey)).Patch("/", star.UpdateUserHandler)
	group.Use(userhandler.JwtMiddleware(tokenKey)).Delete("/", star.DeleteUserHandler)
}

// SetupWishMoviesRoutes ...
func SetupWishMoviesRoutes(app *fiber.App, tokenKey string) {
	star := userhandler.Start(tokenKey)
	group := app.Group("/wishlist")

	group.Use(userhandler.JwtMiddleware(tokenKey)).Get("/", star.GetwishListHandler)
	group.Use(userhandler.JwtMiddleware(tokenKey)).Post("/", star.WhishMoviesHandler)
}
