package main

import (
	"github.com/gofiber/fiber/v2"
	"github.coom/tavo1987/api-project-manager/api/handlers/organizations"
	"github.coom/tavo1987/api-project-manager/api/handlers/projects"
)

func main() {
	app := fiber.New()

	app.Post("/organizations/store", organizations.Create)
	app.Post("/projects/store", projects.Create)

	app.Listen(":3000")
}
