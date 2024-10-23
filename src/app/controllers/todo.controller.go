package controllers

import (
	"github.com/bancodebogota/bbog-dig-pl-go-mngr-template/src/app/services"
	"github.com/bancodebogota/bbog-dig-pl-go-mngr-template/src/utils/middleware"

	"github.com/gofiber/fiber/v2"
)

func TodoRoutes(app fiber.Router) {
	r := app.Group("/todo").Use(middleware.Auth)

	r.Post("/create", services.CreateTodo)
	r.Get("/list", services.GetTodos)
	r.Get("/:todoID", services.GetTodo)
	r.Patch("/:todoID", services.UpdateTodoTitle)
	r.Patch("/:todoID/check", services.CheckTodo)
	r.Delete("/:todoID", services.DeleteTodo)
}
