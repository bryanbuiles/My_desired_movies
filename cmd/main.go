package main

import (
	"github.com/bryanbuiles/movie_suggester/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type internalError struct {
	Message string `json:"message"`
}

func main() {
	// Error handler will handle all the errors that come internally to fiber
	app := fiber.New(fiber.Config{
		ErrorHandler: func(context *fiber.Ctx, err error) error { // Error handler is called when an error occurs
			code := fiber.StatusInternalServerError // default 500
			var msg string
			e, ok := err.(*fiber.Error) // // Retreive the custom statuscode if it's an fiber.*Error
			if ok {
				code = e.Code   // error code
				msg = e.Message // error message
			}
			if msg == "" {
				msg = "Cannot procces the http call"
			}
			// error custom page
			err = context.Status(code).JSON(internalError{
				Message: msg,
			})
			return nil
		},
	})
	key := "tokenKey"
	app.Use(recover.New()) // allows the server to keep running even if occurs a panic
	app.Use(cors.New())
	router.SetupMoviesRoutes(app, key)
	router.SetupUserRoutes(app, key)
	router.SetupWishMoviesRoutes(app, key)
	app.Listen(":3001")
}
