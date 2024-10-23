package controllers

import (
	"github.com/bancodebogota/bbog-dig-pl-go-mngr-template/src/app/services"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app fiber.Router) {
	r := app.Group("/auth")

	r.Post("/signup", services.Signup)
	r.Post("/login", services.Login)
}
