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
	app fiber.Router, api *user.User, casbin *casbin.Middleware, idempotencyMiddleware fiber.Handler,
) {
	group := app.Group(userPrefix)

	// Statistic API
	group.Get(
		"/data-statistic",
		casbin.RequiresRoles([]string{config.UserRoleUser}),
		api.StatisticApi.GetDataStatistic,
	)

	// Dataset API
	group.Get(
		"/instruction-data",
		casbin.RequiresRoles([]string{config.UserRoleUser}),
		api.DatasetApi.GetInstructionData,
	)
	group.Get(
		"/instruction-data/list",
		casbin.RequiresRoles([]string{config.UserRoleUser}),
		api.DatasetApi.GetInstructionDataList,
	)
	group.Post(
		"/instruction-data",
		casbin.RequiresRoles([]string{config.UserRoleUser}),
		idempotencyMiddleware,
		api.DatasetApi.InsertInstructionData,
	)
	group.Put(
		"/instruction-data",
		casbin.RequiresRoles([]string{config.UserRoleUser}),
		api.DatasetApi.UpdateInstructionData,
	)
	group.Delete(
		"/instruction-data",
		casbin.RequiresRoles([]string{config.UserRoleUser}),
		api.DatasetApi.DeleteInstructionData,
	)
}
