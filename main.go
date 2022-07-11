package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/litriggy/go-serve/database"
	"github.com/litriggy/go-serve/routes"
)

func main() {
	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":8000")
}
