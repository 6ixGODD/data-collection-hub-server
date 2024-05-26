package mods

import (
	"fmt"

	"data-collection-hub-server/internal/pkg/config"
	"data-collection-hub-server/internal/pkg/service/common/mods"
	"data-collection-hub-server/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

type IdempotencyMiddleware struct {
	IdempotencyService mods.IdempotencyService
	Config             *config.Config
}

func (m *IdempotencyMiddleware) IdempotencyMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.UserContext()
		token := c.Get(m.Config.IdempotencyConfig.IdempotencyTokenHeader)
		if token != "" {
			ok, err := m.IdempotencyService.CheckIdempotencyToken(ctx, token)
			if err != nil {
				return err
			}
			if ok {
				return c.Next()
			}
		}
		return errors.IdempotencyCheckFailed(fmt.Errorf("idempotency check failed"))
	}
}
