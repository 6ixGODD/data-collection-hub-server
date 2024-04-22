package router

import (
	"data-collection-hub-server/router/v1"
	"github.com/gofiber/fiber/v2"
)

func RegisterRouter(app *fiber.App) {
	router := app.Group("/api")
	registerRouterV1(router)
}

func registerRouterV1(router fiber.Router) {
	v1.RegisterAdminRouter(router)
	v1.RegisterCommonRouter(router)
	v1.RegisterUserRouter(router)
}
