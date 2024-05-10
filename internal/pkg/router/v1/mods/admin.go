package mods

import (
	"data-collection-hub-server/internal/pkg/api/v1/admin"
	"github.com/gofiber/contrib/casbin"
	"github.com/gofiber/fiber/v2"
)

const adminPrefix = "/admin"

type AdminRouter struct{}

func NewAdminRouter() *AdminRouter {
	return &AdminRouter{}
}

// RegisterAdminRouter registers the admin router.
func (a *AdminRouter) RegisterAdminRouter(app fiber.Router, api *admin.Admin, rbac *casbin.Middleware) {
	group := app.Group(adminPrefix)

	group.Get(
		"/dataStatistic", rbac.RequiresRoles([]string{"admin"}), api.StatisticApi.GetDataStatistic,
	)
	group.Get(
		"/userStatistic", rbac.RequiresRoles([]string{"admin"}), api.StatisticApi.GetUserStatistic,
	)
	group.Get(
		"/userStatistic/list", rbac.RequiresRoles([]string{"admin"}), api.StatisticApi.GetUserStatisticList,
	)

	group.Get(
		"/instructionData", rbac.RequiresRoles([]string{"admin"}), api.DataAuditApi.GetInstructionData,
	)
	group.Get(
		"/instructionData/list", rbac.RequiresRoles([]string{"admin"}), api.DataAuditApi.GetInstructionDataList,
	)
	group.Get(
		"instructionData/approve", rbac.RequiresRoles([]string{"admin"}), api.DataAuditApi.ApproveInstructionData,
	)
	group.Get(
		"/instructionData/reject", rbac.RequiresRoles([]string{"admin"}), api.DataAuditApi.RejectInstructionData,
	)
	group.Get(
		"/instructionData/update", rbac.RequiresRoles([]string{"admin"}), api.DataAuditApi.UpdateInstructionData,
	)
	group.Get(
		"/instructionData/delete", rbac.RequiresRoles([]string{"admin"}), api.DataAuditApi.DeleteInstructionData,
	)

	group.Post(
		"/notice", rbac.RequiresRoles([]string{"admin"}), api.NoticeApi.InsertNotice,
	)
	group.Put(
		"/notice", rbac.RequiresRoles([]string{"admin"}), api.NoticeApi.UpdateNotice,
	)
	group.Delete(
		"/notice", rbac.RequiresRoles([]string{"admin"}), api.NoticeApi.DeleteNotice,
	)

	group.Post(
		"/user", rbac.RequiresRoles([]string{"admin"}), api.UserApi.InsertUser,
	)
	group.Get(
		"/user", rbac.RequiresRoles([]string{"admin"}), api.UserApi.GetUser,
	)
	group.Get(
		"/user/list", rbac.RequiresRoles([]string{"admin"}), api.UserApi.GetUserList,
	)
	group.Put(
		"/user", rbac.RequiresRoles([]string{"admin"}), api.UserApi.UpdateUser,
	)
	group.Delete(
		"/user", rbac.RequiresRoles([]string{"admin"}), api.UserApi.DeleteUser,
	)
	group.Post(
		"/user/password", rbac.RequiresRoles([]string{"admin"}), api.UserApi.ChangeUserPassword,
	)

	group.Post(
		"/documentation", rbac.RequiresRoles([]string{"admin"}), api.DocumentationApi.InsertDocumentation,
	)
	group.Put(
		"/documentation", rbac.RequiresRoles([]string{"admin"}), api.DocumentationApi.UpdateDocumentation,
	)
	group.Delete(
		"/documentation", rbac.RequiresRoles([]string{"admin"}), api.DocumentationApi.DeleteDocumentation,
	)

	group.Get(
		"/loginLog", rbac.RequiresRoles([]string{"admin"}), api.LogsApi.GetLoginLog,
	)
	group.Get(
		"/loginLog/list", rbac.RequiresRoles([]string{"admin"}), api.LogsApi.GetLoginLogList,
	)
	group.Get(
		"/operationLog", rbac.RequiresRoles([]string{"admin"}), api.LogsApi.GetOperationLog,
	)
	group.Get(
		"/operationLog/list", rbac.RequiresRoles([]string{"admin"}), api.LogsApi.GetOperationLogList,
	)
	group.Get(
		"/errorLog", rbac.RequiresRoles([]string{"admin"}), api.LogsApi.GetErrorLog,
	)
	group.Get(
		"/errorLog/list", rbac.RequiresRoles([]string{"admin"}), api.LogsApi.GetErrorLogList,
	)
}
