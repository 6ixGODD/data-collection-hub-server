package mods

import (
	"data-collection-hub-server/internal/pkg/schema"
	commonservice "data-collection-hub-server/internal/pkg/service/common/mods"
	"data-collection-hub-server/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

type ProfileApi struct {
	commonservice.ProfileService
}

func NewProfileApi(profileService commonservice.ProfileService) ProfileApi {
	return ProfileApi{profileService}
}

func (api *ProfileApi) GetProfile(c *fiber.Ctx) error {
	resp, err := api.ProfileService.GetProfile(c.Context())
	if err != nil {
		return err
	}

	return c.JSON(
		schema.Response{
			Code:    errors.CodeSuccess,
			Message: errors.MessageSuccess,
			Data:    resp,
		},
	)
}
