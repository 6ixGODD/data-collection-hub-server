package modules

import (
	"data-collection-hub-server/internal/pkg/router"
	"github.com/gofiber/fiber/v2"
)

type CommonRouter struct {
	Router *router.Router
}

// RegisterCommonRouter registers the common_service router.
func RegisterCommonRouter(app fiber.Router) { // TODO: Implement
	common := app.Group("/common_service")

	common.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})
	common.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})
}
