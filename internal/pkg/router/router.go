package router

import (
	"data-collection-hub-server/internal/pkg/router/v1"
	"github.com/gofiber/fiber/v2"
)

const prefix = "/api"

type Router struct {
	RouterV1 *router.Router
}

func New() *Router {
	return &Router{}
}

func (r *Router) RegisterRouter(app *fiber.App) {
	_ = app.Group(prefix) // TODO: Implement
}
