package router

import (
	"data-collection-hub-server/internal/pkg/router/v1"
	"github.com/gofiber/contrib/casbin"
	"github.com/gofiber/fiber/v2"
)

const prefix = "/api"

type Router struct {
	RouterV1 *router.Router
}

func New(routerV1 *router.Router) *Router {
	return &Router{
		RouterV1: routerV1,
	}
}

func (r *Router) RegisterRouter(app *fiber.App, rbac *casbin.Middleware) {
	group := app.Group(prefix) // TODO: Implement
	r.RouterV1.RegisterRouter(&group, rbac)
}
