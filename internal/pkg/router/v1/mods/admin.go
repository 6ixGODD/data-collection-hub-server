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
func (a *AdminRouter) RegisterAdminRouter(app fiber.Router, api *admin.Admin, rbac *casbin.Middleware) {
	group := app.Group(adminPrefix)

	group.Get("/data-statistic", rbac.RequiresRoles([]string{config.UserRoleAdmin}), api.StatisticApi.GetDataStatistic)
	group.Get(
		"/user-statistic", rbac.RequiresRoles([]string{config.UserRoleAdmin}), api.StatisticApi.GetUserStatistic,
	)
	group.Get(
		"/user-statistic/list", rbac.RequiresRoles([]string{config.UserRoleAdmin}),
		api.StatisticApi.GetUserStatisticList,
	)

	group.Get(
		"/instruction-data", rbac.RequiresRoles([]string{config.UserRoleAdmin}), api.DataAuditApi.GetInstructionData,
	)
	group.Get(
		"/instruction-data/list", rbac.RequiresRoles([]string{config.UserRoleAdmin}),
		api.DataAuditApi.GetInstructionDataList,
	)
	group.Get(
		"instruction-data/approve", rbac.RequiresRoles([]string{config.UserRoleAdmin}),
		api.DataAuditApi.ApproveInstructionData,
	)
	group.Get(
		"/instruction-data/reject", rbac.RequiresRoles([]string{config.UserRoleAdmin}),
		api.DataAuditApi.RejectInstructionData,
	)
	group.Get(
		"/instruction-data/update", rbac.RequiresRoles([]string{config.UserRoleAdmin}),
		api.DataAuditApi.UpdateInstructionData,
	)
	group.Delete(
		"/instruction-data/", rbac.RequiresRoles([]string{config.UserRoleAdmin}),
		api.DataAuditApi.DeleteInstructionData,
	)

	group.Post(
		"/notice", rbac.RequiresRoles([]string{config.UserRoleAdmin}), api.NoticeApi.InsertNotice,
	)
	group.Put(
		"/notice", rbac.RequiresRoles([]string{config.UserRoleAdmin}), api.NoticeApi.UpdateNotice,
	)
	group.Delete(
		"/notice", rbac.RequiresRoles([]string{config.UserRoleAdmin}), api.NoticeApi.DeleteNotice,
	)

	group.Post(
		"/user", rbac.RequiresRoles([]string{config.UserRoleAdmin}), api.UserApi.InsertUser,
	)
	group.Get(
		"/user", rbac.RequiresRoles([]string{config.UserRoleAdmin}), api.UserApi.GetUser,
	)
	group.Get(
		"/user/list", rbac.RequiresRoles([]string{config.UserRoleAdmin}), api.UserApi.GetUserList,
	)
	group.Put(
		"/user", rbac.RequiresRoles([]string{config.UserRoleAdmin}), api.UserApi.UpdateUser,
	)
	group.Delete(
		"/user", rbac.RequiresRoles([]string{config.UserRoleAdmin}), api.UserApi.DeleteUser,
	)
	group.Post(
		"/user/password", rbac.RequiresRoles([]string{config.UserRoleAdmin}), api.UserApi.ChangeUserPassword,
	)

	group.Post(
		"/documentation", rbac.RequiresRoles([]string{config.UserRoleAdmin}), api.DocumentationApi.InsertDocumentation,
	)
	group.Put(
		"/documentation", rbac.RequiresRoles([]string{config.UserRoleAdmin}), api.DocumentationApi.UpdateDocumentation,
	)
	group.Delete(
		"/documentation", rbac.RequiresRoles([]string{config.UserRoleAdmin}), api.DocumentationApi.DeleteDocumentation,
	)

	group.Get(
		"/login-log", rbac.RequiresRoles([]string{config.UserRoleAdmin}), api.LogsApi.GetLoginLog,
	)
	group.Get(
		"/login-log/list", rbac.RequiresRoles([]string{config.UserRoleAdmin}), api.LogsApi.GetLoginLogList,
	)
	group.Get(
		"/operation-log", rbac.RequiresRoles([]string{config.UserRoleAdmin}), api.LogsApi.GetOperationLog,
	)
	group.Get(
		"/operation-log/list", rbac.RequiresRoles([]string{config.UserRoleAdmin}), api.LogsApi.GetOperationLogList,
	)
}
