package router

import (
	"data-collection-hub-server/internal/pkg/api/v1"
	"data-collection-hub-server/internal/pkg/router/v1/modules"
	"github.com/gofiber/fiber/v2"
)

const v1Prefix = "/v1"

type Router struct {
	ApiV1        *api.Api
	AdminRouter  *modules.AdminRouter
	CommonRouter *modules.CommonRouter
	UserRouter   *modules.UserRouter
}

// RegisterRouter registers the router.
func RegisterRouter(app *fiber.App) {
	// TODO: Implement
}
