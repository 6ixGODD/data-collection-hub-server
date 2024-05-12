package mods

import (
	"data-collection-hub-server/internal/pkg/api/v1/common"
	"data-collection-hub-server/internal/pkg/config"
	"github.com/gofiber/contrib/casbin"
	"github.com/gofiber/fiber/v2"
)

const commonPrefix = "/common"

type CommonRouter struct{}

func NewCommonRouter() *CommonRouter {
	return &CommonRouter{}
}

// RegisterCommonRouter registers the common router.
func (c *CommonRouter) RegisterCommonRouter(
	app fiber.Router, api *common.Common, rbac *casbin.Middleware,
) {
	group := app.Group(commonPrefix)

	group.Post("/login", api.Login)
	group.Post("/logout", api.Logout)
	group.Get("/refreshToken", api.RefreshToken)
	group.Get("/profile", rbac.RequiresRoles([]string{config.UserRoleAdmin, config.UserRoleUser}), api.GetProfile)
	group.Put(
		"/changePassword", rbac.RequiresRoles([]string{config.UserRoleAdmin, config.UserRoleUser}), api.ChangePassword,
	)

	group.Get("/notice", api.GetNotice)
	group.Get("/notice/list", api.GetNoticeList)

	group.Get("/documentation", api.GetDocumentation)
	group.Get("/documentation/list", api.GetDocumentationList)
}
