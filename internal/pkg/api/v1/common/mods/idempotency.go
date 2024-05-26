package mods

import (
	"data-collection-hub-server/internal/pkg/domain/vo"
	commonservice "data-collection-hub-server/internal/pkg/service/common/mods"
	"data-collection-hub-server/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

type IdempotencyApi struct {
	IdempotencyService commonservice.IdempotencyService
}

func (i *IdempotencyApi) GenerateIdempotencyToken(c *fiber.Ctx) error {
	resp, err := i.IdempotencyService.GenerateIdempotencyToken(c.UserContext())
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
