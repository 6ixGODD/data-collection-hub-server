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
	app.Get("/profile", rbac.RequiresRoles([]string{config.UserRoleAdmin, config.UserRoleUser}), api.GetProfile)
	app.Put(
		"/change-password", rbac.RequiresRoles([]string{config.UserRoleAdmin, config.UserRoleUser}), api.ChangePassword,
	)

	authGroup := app.Group("/auth")
	authGroup.Post("/login", api.Login)
	authGroup.Post("/logout", api.Logout)
	authGroup.Get("/refresh", api.RefreshToken)

	noticeGroup := app.Group("/notice")
	noticeGroup.Get("/", api.GetNotice)
	noticeGroup.Get("/list", api.GetNoticeList)

	documentationGroup := app.Group("/documentation")
	documentationGroup.Get("/documentation", api.GetDocumentation)
	documentationGroup.Get("/documentation/list", api.GetDocumentationList)
}
