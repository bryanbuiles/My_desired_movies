package main

import (
	"github.com/bryanbuiles/movie_suggester/api"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	api.SetupMoviesRoutes(app)

	app.Listen(":3001")

}
