package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tavo1987/api-project-manager/api/handlers/employees"
	"github.coom/tavo1987/api-project-manager/api/handlers/organizations"
	"github.coom/tavo1987/api-project-manager/api/handlers/projects"
)

func main() {
	app := fiber.New()

	app.Post("/organizations/store", organizations.Create)
	app.Post("/projects/store", projects.Create)
	app.Post("/employees/store", employees.Create)

	app.Listen(":3000")
}
