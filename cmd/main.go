package main

import (
	"github.com/bryanbuiles/movie_suggester/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type internalError struct {
	Message string `json:"message"`
}

func main() {
	// Error handler va a manejar todos los errores que internamente le llegan a fiber
	app := fiber.New(fiber.Config{
		ErrorHandler: func(context *fiber.Ctx, err error) error { // Error handler es llamada cuando ocurre un error
			code := fiber.StatusInternalServerError // default 500 por si no trae ninguno
			var msg string
			e, ok := err.(*fiber.Error) // // Retreive the custom statuscode if it's an fiber.*Error
			if ok {
				code = e.Code   // codigo error
				msg = e.Message // codigo mensaje
			}
			if msg == "" {
				msg = "Cannot procces the http call"
			}
			// pagina de error custom
			err = context.Status(code).JSON(internalError{ // JSON convierte una struct o string a json
				Message: msg,
			})
			return nil
		},
	})
	key := "tokenKey"
	app.Use(recover.New()) // recover from panic, Permite que siga andando el servidor si hay un panic
	router.SetupMoviesRoutes(app, key)
	router.SetupUserRoutes(app, key)
	router.SetupWishMoviesRoutes(app, key)
	app.Listen(":3001")

}
