package v1

import (
	"data-collection-hub-server/internal/pkg/router"
	"github.com/gofiber/fiber/v2"
)

type AdminRouter struct {
	Router *router.Router
}

// RegisterAdminRouter registers the admin_service router.
func RegisterAdminRouter(router fiber.Router) { // TODO: Implement
	admin := router.Group("/admin_service")

	admin.Get("/login", func(c *fiber.Ctx) error {
		return c.SendString("admin_service login")
	})
	admin.Get("/logout", func(c *fiber.Ctx) error {
		return c.SendString("admin_service logout")
	})
	admin.Get("/profile", func(c *fiber.Ctx) error {
		return c.SendString("admin_service profile")
	})
	admin.Get("/dashboard", func(c *fiber.Ctx) error {
		return c.SendString("admin_service dashboard")
	})
}
