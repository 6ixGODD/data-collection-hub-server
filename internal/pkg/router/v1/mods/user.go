package mods

import (
	"data-collection-hub-server/internal/pkg/api/v1/user"
	"data-collection-hub-server/internal/pkg/models"
	"github.com/gofiber/contrib/casbin"
	"github.com/gofiber/fiber/v2"
)

const userPrefix = "/user"

type UserRouter struct{}

func NewUserRouter() *UserRouter {
	return &UserRouter{}
}

// RegisterUserRouter registers the user router.
func (u *UserRouter) RegisterUserRouter(app fiber.Router, api *user.User, rbac *casbin.Middleware) { // TODO: Implement
	group := app.Group(userPrefix)

	group.Get(
		"/dataStatistic", rbac.RequiresRoles([]string{models.UserRoleUser}), api.GetDataStatistic,
	)

	group.Get(
		"/instructionData", rbac.RequiresRoles([]string{models.UserRoleUser}), api.GetInstructionData,
	)
	group.Get(
		"/instructionData/list", rbac.RequiresRoles([]string{models.UserRoleUser}), api.GetInstructionDataList,
	)
	group.Post(
		"/instructionData", rbac.RequiresRoles([]string{models.UserRoleUser}), api.InsertInstructionData,
	)
	group.Put(
		"/instructionData", rbac.RequiresRoles([]string{models.UserRoleUser}), api.UpdateInstructionData,
	)
	group.Delete(
		"/instructionData", rbac.RequiresRoles([]string{models.UserRoleUser}), api.DeleteInstructionData,
	)
}
