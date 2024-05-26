package mods

import (
	"data-collection-hub-server/internal/pkg/api/v1/common"
	"data-collection-hub-server/internal/pkg/config"
	"github.com/gofiber/contrib/casbin"
	"github.com/gofiber/fiber/v2"
)

type CommonRouter struct {
	Config *config.Config
}

// RegisterCommonRouter registers the common router.
func (c *CommonRouter) RegisterCommonRouter(
	app fiber.Router, api *common.Common, rbac *casbin.Middleware,
) {
	app.Get(
		"/ping",
		func(c *fiber.Ctx) error {
			return c.SendString("pong")
		},
	)
	app.Get(
		"/profile",
		rbac.RequiresRoles([]string{config.UserRoleAdmin, config.UserRoleUser}),
		api.ProfileApi.GetProfile,
	)
	app.Put(
		"/change-password",
		rbac.RequiresRoles([]string{config.UserRoleAdmin, config.UserRoleUser}),
		api.AuthApi.ChangePassword,
	)

	authGroup := app.Group("/auth")
	authGroup.Post(
		"/login",
		api.AuthApi.Login,
	)
	authGroup.Post(
		"/logout",
		api.AuthApi.Logout,
	)
	authGroup.Get(
		"/refresh",
		api.AuthApi.RefreshToken,
	)

	noticeGroup := app.Group("/notice")
	noticeGroup.Get(
		"/",
		api.NoticeApi.GetNotice,
	)
	noticeGroup.Get(
		"/list",
		api.NoticeApi.GetNoticeList,
	)

	documentationGroup := app.Group("/documentation")
	documentationGroup.Get(
		"/documentation",
		api.DocumentationApi.GetDocumentation,
	)
	documentationGroup.Get(
		"/documentation/list",
		api.DocumentationApi.GetDocumentationList,
	)
}
