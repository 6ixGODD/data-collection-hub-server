package errors

import (
	"errors"

	"data-collection-hub-server/internal/pkg/schema"
	e "data-collection-hub-server/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	var appErr *e.AppError
	if errors.As(err, &appErr) {
		return c.Status(appErr.Status()).JSON(
			schema.Response{
				Code:    appErr.Code(),
				Message: appErr.Error(),
				Data:    nil,
			},
		)
	}

	return c.Status(fiber.StatusInternalServerError).JSON(
		schema.Response{
			Code:    e.CodeUnknownError,
			Message: err.Error(),
			Data:    nil,
		},
	)
}
