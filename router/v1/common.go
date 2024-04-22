package v1

import (
	"github.com/gofiber/fiber/v2"
)

// RegisterCommonRouter registers the common router.
func RegisterCommonRouter(app fiber.Router) { // TODO: Implement
	common := app.Group("/common")

	common.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})
	common.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})
}
