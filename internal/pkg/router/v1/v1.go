package router

import (
	"data-collection-hub-server/internal/pkg/api/v1"
	"data-collection-hub-server/internal/pkg/router/v1/mods"
	"github.com/gofiber/contrib/casbin"
	"github.com/gofiber/fiber/v2"
)

const v1Prefix = "/v1"

type Router struct {
	ApiV1        *api.Api
	AdminRouter  *mods.AdminRouter
	CommonRouter *mods.CommonRouter
	UserRouter   *mods.UserRouter
}

func New(
	apiV1 *api.Api, adminRouter *mods.AdminRouter, commonRouter *mods.CommonRouter, userRouter *mods.UserRouter,
) *Router {
	return &Router{
		ApiV1:        apiV1,
		AdminRouter:  adminRouter,
		CommonRouter: commonRouter,
		UserRouter:   userRouter,
	}
}

func (a *Router) RegisterRouter(router *fiber.Router, rbac *casbin.Middleware) {
	a.registerV1Router(router, rbac)
}

func (a *Router) registerV1Router(router *fiber.Router, rbac *casbin.Middleware) {
	v1Router := (*router).Group(v1Prefix)
	a.AdminRouter.RegisterAdminRouter(v1Router, a.ApiV1.AdminApi, rbac)
	a.CommonRouter.RegisterCommonRouter(v1Router, a.ApiV1.CommonApi, rbac)
	a.UserRouter.RegisterUserRouter(v1Router, a.ApiV1.UserApi, rbac)
}
