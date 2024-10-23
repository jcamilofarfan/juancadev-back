package main

import (
	"fmt"

	"github.com/bancodebogota/bbog-dig-pl-go-mngr-template/src/app/models"
	"github.com/bancodebogota/bbog-dig-pl-go-mngr-template/src/app/controllers"
	"github.com/bancodebogota/bbog-dig-pl-go-mngr-template/src/config/database"
	"github.com/bancodebogota/bbog-dig-pl-go-mngr-template/src/config"
	"github.com/bancodebogota/bbog-dig-pl-go-mngr-template/src/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := Setup()

	app.Listen(fmt.Sprintf(":%v", config.PORT))
}

func Setup() *fiber.App {
	database.Connect()
	database.Migrate(&models.User{}, &models.Todo{})

	app := fiber.New(fiber.Config{
		ErrorHandler: utils.ErrorHandler,
	})

	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("healthy")
	})

	controllers.AuthRoutes(app)
	controllers.TodoRoutes(app)

	return app
}