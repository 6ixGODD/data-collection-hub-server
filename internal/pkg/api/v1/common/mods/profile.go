package mods

import (
	"data-collection-hub-server/internal/pkg/domain/vo"
	commonservice "data-collection-hub-server/internal/pkg/service/common/mods"
	"data-collection-hub-server/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

type ProfileApi struct {
	ProfileService commonservice.ProfileService
}

func (api *ProfileApi) GetProfile(c *fiber.Ctx) error {
	resp, err := api.ProfileService.GetProfile(c.UserContext())
	if err != nil {
		return err
	}

	return c.JSON(
		vo.Response{
			Code:    errors.CodeSuccess,
			Message: errors.MessageSuccess,
			Data:    resp,
		},
	)
}
