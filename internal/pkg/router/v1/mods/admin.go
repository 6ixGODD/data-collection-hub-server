package mods

import (
	"data-collection-hub-server/internal/pkg/api/v1/admin"
	"data-collection-hub-server/internal/pkg/config"
	"github.com/gofiber/contrib/casbin"
	"github.com/gofiber/fiber/v2"
)

const adminPrefix = "/admin"

type AdminRouter struct{}

// RegisterAdminRouter registers the admin router.
func (a *AdminRouter) RegisterAdminRouter(
	app fiber.Router, api *admin.Admin, casbin *casbin.Middleware, idempotencyMiddleware fiber.Handler,
) {
	group := app.Group(adminPrefix)

	group.Get(
		"/data-statistic",
		casbin.RequiresRoles([]string{config.UserRoleAdmin}),
		api.StatisticApi.GetDataStatistic,
	)
	group.Get(
		"/user-statistic",
		casbin.RequiresRoles([]string{config.UserRoleAdmin}),
		api.StatisticApi.GetUserStatistic,
	)
	group.Get(
		"/user-statistic/list",
		casbin.RequiresRoles([]string{config.UserRoleAdmin}),
		api.StatisticApi.GetUserStatisticList,
	)

	group.Get(
		"/instruction-data",
		casbin.RequiresRoles([]string{config.UserRoleAdmin}),
		api.DataAuditApi.GetInstructionData,
	)
	group.Get(
		"/instruction-data/list",
		casbin.RequiresRoles([]string{config.UserRoleAdmin}),
		api.DataAuditApi.GetInstructionDataList,
	)
	group.Put(
		"instruction-data/approve",
		casbin.RequiresRoles([]string{config.UserRoleAdmin}),
		api.DataAuditApi.ApproveInstructionData,
	)
	group.Put(
		"/instruction-data/reject",
		casbin.RequiresRoles([]string{config.UserRoleAdmin}),
		api.DataAuditApi.RejectInstructionData,
	)
	group.Get(
		"/instruction-data/export",
		casbin.RequiresRoles([]string{config.UserRoleAdmin}),
		api.DataAuditApi.ExportInstructionData,
	)
	group.Get(
		"/instruction-data/export/alpaca",
		casbin.RequiresRoles([]string{config.UserRoleAdmin}),
		api.DataAuditApi.ExportInstructionDataAsAlpaca,
	)
	group.Put(
		"/instruction-data",
		casbin.RequiresRoles([]string{config.UserRoleAdmin}),
		idempotencyMiddleware,
		api.DataAuditApi.UpdateInstructionData,
	)
	group.Delete(
		"/instruction-data",
		casbin.RequiresRoles([]string{config.UserRoleAdmin}),
		api.DataAuditApi.DeleteInstructionData,
	)

	group.Post(
		"/notice",
		casbin.RequiresRoles([]string{config.UserRoleAdmin}),
		api.NoticeApi.InsertNotice,
	)
	group.Put(
		"/notice",
		casbin.RequiresRoles([]string{config.UserRoleAdmin}),
		api.NoticeApi.UpdateNotice,
	)
	group.Delete(
		"/notice",
		casbin.RequiresRoles([]string{config.UserRoleAdmin}),
		api.NoticeApi.DeleteNotice,
	)

	group.Post(
		"/user",
		casbin.RequiresRoles([]string{config.UserRoleAdmin}),
		api.UserApi.InsertUser,
	)
	group.Get(
		"/user",
		casbin.RequiresRoles([]string{config.UserRoleAdmin}),
		api.UserApi.GetUser,
	)
	group.Get(
		"/user/list",
		casbin.RequiresRoles([]string{config.UserRoleAdmin}),
		api.UserApi.GetUserList,
	)

	group.Put(
		"/user",
		casbin.RequiresRoles([]string{config.UserRoleAdmin}),
		api.UserApi.UpdateUser,
	)
	group.Delete(
		"/user",
		casbin.RequiresRoles([]string{config.UserRoleAdmin}),
		api.UserApi.DeleteUser,
	)
	group.Put(
		"/user/password",
		casbin.RequiresRoles([]string{config.UserRoleAdmin}),
		api.UserApi.ChangeUserPassword,
	)

	group.Post(
		"/documentation",
		casbin.RequiresRoles([]string{config.UserRoleAdmin}),
		api.DocumentationApi.InsertDocumentation,
	)
	group.Put(
		"/documentation",
		casbin.RequiresRoles([]string{config.UserRoleAdmin}),
		api.DocumentationApi.UpdateDocumentation,
	)
	group.Delete(
		"/documentation",
		casbin.RequiresRoles([]string{config.UserRoleAdmin}),
		api.DocumentationApi.DeleteDocumentation,
	)

	group.Get(
		"/login-log/list",
		casbin.RequiresRoles([]string{config.UserRoleAdmin}),
		api.LogsApi.GetLoginLogList,
	)
	group.Get(
		"/operation-log/list",
		casbin.RequiresRoles([]string{config.UserRoleAdmin}),
		api.LogsApi.GetOperationLogList,
	)
}
