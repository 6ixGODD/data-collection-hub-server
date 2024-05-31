package mods

import (
	"time"

	"data-collection-hub-server/internal/pkg/domain/vo"
	"data-collection-hub-server/internal/pkg/domain/vo/user"
	userservice "data-collection-hub-server/internal/pkg/service/user/mods"
	"data-collection-hub-server/pkg/errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type StatisticApi struct {
	StatisticService userservice.StatisticService
	Validator        *validator.Validate
}

func (s *StatisticApi) GetDataStatistic(c *fiber.Ctx) error {
	req := new(user.GetDataStatisticRequest)
	if err := c.QueryParser(req); err != nil {
		return errors.InvalidRequest(err)
	}
	if err := s.Validator.Struct(req); err != nil {
		return errors.InvalidParams(err) // Compare this line with the original one
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

	resp, err := s.StatisticService.GetDataStatistic(c.UserContext(), startDate, endDate)
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
