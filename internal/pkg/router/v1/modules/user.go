package modules

import (
	"data-collection-hub-server/internal/pkg/router/v1"
	"github.com/gofiber/fiber/v2"
)

type UserRouter struct {
	Router *router.Router
}

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
