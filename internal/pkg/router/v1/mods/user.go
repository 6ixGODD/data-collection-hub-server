package mods

import (
	"data-collection-hub-server/internal/pkg/api/v1/user"
	"data-collection-hub-server/internal/pkg/config"
	"github.com/gofiber/contrib/casbin"
	"github.com/gofiber/fiber/v2"
)

const userPrefix = "/user"

type UserRouter struct{}

// RegisterUserRouter registers the user router.
func (u *UserRouter) RegisterUserRouter(
	app fiber.Router, api *user.User, rbac *casbin.Middleware, idempotencyMiddleware fiber.Handler,
) {
	group := app.Group(userPrefix)

	// Statistic API
	group.Get(
		"/data-statistic",
		rbac.RequiresRoles([]string{config.UserRoleUser}),
		api.StatisticApi.GetDataStatistic,
	)

	// Dataset API
	group.Get(
		"/instruction-data",
		rbac.RequiresRoles([]string{config.UserRoleUser}),
		idempotencyMiddleware,
		api.DatasetApi.GetInstructionData,
	)
	group.Get(
		"/instruction-data/list",
		rbac.RequiresRoles([]string{config.UserRoleUser}),
		api.DatasetApi.GetInstructionDataList,
	)
	group.Post(
		"/instruction-data",
		rbac.RequiresRoles([]string{config.UserRoleUser}),
		api.DatasetApi.InsertInstructionData,
	)
	group.Put(
		"/instruction-data",
		rbac.RequiresRoles([]string{config.UserRoleUser}),
		api.DatasetApi.UpdateInstructionData,
	)
	group.Delete(
		"/instruction-data",
		rbac.RequiresRoles([]string{config.UserRoleUser}),
		api.DatasetApi.DeleteInstructionData,
	)
}
