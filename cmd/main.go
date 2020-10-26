package main

import (
	"github.com/gofiber/fiber/v2"
	"github.coom/tavo1987/api-project-manager/api/handlers/organizations"
)

func main() {
	app := fiber.New()

	app.Post("/organization/store", organizations.Create)

	app.Listen(":3000")
}
