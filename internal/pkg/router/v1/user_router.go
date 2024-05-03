package v1

import (
	"github.com/gofiber/fiber/v2"
)

// RegisterUserRouter registers the user_service router.
func RegisterUserRouter(app fiber.Router) {
	user := app.Group("/user_service")

	user.Get("/login", func(c *fiber.Ctx) error {
		return c.SendString("user_service login") // TODO: Implement
	})
	user.Get("/logout", func(c *fiber.Ctx) error {
		return c.SendString("user_service logout") // TODO: Implement
	})
	user.Get("/profile", func(c *fiber.Ctx) error {
		return c.SendString("user_service profile") // TODO: Implement
	})
	user.Get("/dashboard", func(c *fiber.Ctx) error {
		return c.SendString("user_service dashboard") // TODO: Implement
	})
}
