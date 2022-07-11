package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/litriggy/go-serve/controllers"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Hello)
}
