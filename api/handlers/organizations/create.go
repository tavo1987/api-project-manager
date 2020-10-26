package organizations

import "github.com/gofiber/fiber/v2"

// Create this is a handler to create a new organization
func Create(c *fiber.Ctx) error {

	return c.Send(c.Body())
}
