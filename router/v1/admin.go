package v1

import (
	"github.com/gofiber/fiber/v2"
)

// RegisterAdminRouter registers the admin router.
func RegisterAdminRouter(router fiber.Router) { // TODO: Implement
	admin := router.Group("/admin")

	admin.Get("/login", func(c *fiber.Ctx) error {
		return c.SendString("admin login")
	})
	admin.Get("/logout", func(c *fiber.Ctx) error {
		return c.SendString("admin logout")
	})
	admin.Get("/profile", func(c *fiber.Ctx) error {
		return c.SendString("admin profile")
	})
	admin.Get("/dashboard", func(c *fiber.Ctx) error {
		return c.SendString("admin dashboard")
	})
}
