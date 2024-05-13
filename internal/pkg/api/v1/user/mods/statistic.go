package mods

import (
	"time"

	"data-collection-hub-server/internal/pkg/schema"
	"data-collection-hub-server/internal/pkg/schema/user"
	userservice "data-collection-hub-server/internal/pkg/service/user/mods"
	"data-collection-hub-server/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

type StatisticApi struct {
	userservice.StatisticService
}

func (s *StatisticApi) GetDataStatistic(c *fiber.Ctx) error {
	req := new(user.GetDataStatisticRequest)
	if err := c.QueryParser(req); err != nil {
		return errors.InvalidRequest(err)
	}

	var (
		startDate, endDate *time.Time
		err                error
	)
	if req.StartDate != nil {
		*startDate, err = time.Parse(time.RFC3339, *req.StartDate)
		if err != nil {
			return errors.InvalidRequest(err)
		}
	}
	if req.EndDate != nil {
		*endDate, err = time.Parse(time.RFC3339, *req.EndDate)
		if err != nil {
			return errors.InvalidRequest(err)
		}
	}

	resp, err := s.StatisticService.GetDataStatistic(c.Context(), startDate, endDate)
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
