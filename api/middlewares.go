package api

import (
	"time"

	"github.com/bryanbuiles/movie_suggester/internal/logs"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

func jwtMiddleware(secret string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(secret),
	})
}

func signToken(tokenkey, id string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	claims["sub"] = id
	t, err := token.SignedString([]byte(tokenkey))

	if err != nil {
		logs.Error("Create token fail" + err.Error())
		return ""
	}
	return t // token
}
