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

func (r *Router) RegisterRouter(app *fiber.App, casbin *casbin.Middleware, idempotencyMiddleware fiber.Handler) {
	group := app.Group(prefix)
	r.RouterV1.RegisterRouter(&group, casbin, idempotencyMiddleware)
}
