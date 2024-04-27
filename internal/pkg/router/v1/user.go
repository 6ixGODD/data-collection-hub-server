package v1

import (
	"github.com/gofiber/fiber/v2"
)

// RegisterUserRouter registers the user router.
func RegisterUserRouter(app fiber.Router) {
	user := app.Group("/user")

	user.Get("/login", func(c *fiber.Ctx) error {
		return c.SendString("user login") // TODO: Implement
	})
	user.Get("/logout", func(c *fiber.Ctx) error {
		return c.SendString("user logout") // TODO: Implement
	})
	user.Get("/profile", func(c *fiber.Ctx) error {
		return c.SendString("user profile") // TODO: Implement
	})
	user.Get("/dashboard", func(c *fiber.Ctx) error {
		return c.SendString("user dashboard") // TODO: Implement
	})
}
