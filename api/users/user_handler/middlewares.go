package userhandler

import (
	"time"

	"github.com/bryanbuiles/movie_suggester/internal/logs"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

// JwtMiddleware ...
func JwtMiddleware(secret string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(secret),
	})
}

// signToken firm and return the token
func signToken(tokenkey, id string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims) // Mapclaims is a hash table
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

// extractUserIDFromJWT extrct the id using JWT
func extractUserIDFromJWT(bearer, tokenkey string) string {
	// bearrer is a value of http request HEader Authorization
	// example: Bearer eyJhbGciOiJIUzI1Ni .... the other part is the token
	token := bearer[7:]
	logs.Info(token)
	toke, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(tokenkey), nil
	})
	if err != nil {
		logs.Error("Failed parse token" + err.Error())
		return ""
	}
	if toke.Valid {
		claims := toke.Claims.(jwt.MapClaims)
		return claims["sub"].(string)
	}
	return ""
}
