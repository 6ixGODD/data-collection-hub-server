package router

import (
	"data-collection-hub-server/internal/pkg/router/v1"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	App        *fiber.App
	RouteGroup *fiber.Router
}

func (r *Router) RegisterRouter() {
	g := r.App.Group("/api")
	r.RouteGroup = &g
	r.registerRouterV1()
}

func (r *Router) registerRouterV1() {
	v1.RegisterAdminRouter(*r.RouteGroup)
	v1.RegisterCommonRouter(*r.RouteGroup)
	v1.RegisterUserRouter(*r.RouteGroup)
}
