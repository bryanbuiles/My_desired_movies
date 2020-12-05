package api

import (
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// CreateUserHandler ...
func (w *WebServices) CreateUserHandler(ctx *fiber.Ctx) error {
	var cmd CreateUserCMD
	err := ctx.BodyParser(&cmd)

	res, err := w.Services.users.SaveUser(cmd)

	if err != nil {
		return fiber.NewError(400, "Create user fail")
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["test-name"] = "bryansito"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, err := token.SignedString([]byte("mysecretkey"))
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	res.JWT = t
	return ctx.JSON(res)
}
