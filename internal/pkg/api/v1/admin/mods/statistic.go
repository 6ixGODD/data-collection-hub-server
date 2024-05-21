package mods

import (
	"time"

	"data-collection-hub-server/internal/pkg/domain/vo"
	"data-collection-hub-server/internal/pkg/domain/vo/admin"
	adminservice "data-collection-hub-server/internal/pkg/service/admin/mods"
	"data-collection-hub-server/pkg/errors"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StatisticApi struct {
	StatisticService adminservice.StatisticService
}

func (s *StatisticApi) GetDataStatistic(c *fiber.Ctx) error {
	req := new(admin.GetDataStatisticRequest)
	if err := c.QueryParser(req); err != nil {
		return errors.InvalidRequest(err)
	}

	// Parse start and end date
	var startDate, endDate *time.Time
	var err error

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

func (s *StatisticApi) GetUserStatistic(c *fiber.Ctx) error {
	req := new(admin.GetUserStatisticRequest)
	if err := c.QueryParser(req); err != nil {
		return errors.InvalidRequest(err)
	}

	userID, err := primitive.ObjectIDFromHex(*req.UserID)
	if err != nil {
		return errors.InvalidRequest(err)
	}

	resp, err := s.StatisticService.GetUserStatistic(c.UserContext(), &userID)
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

func (s *StatisticApi) GetUserStatisticList(c *fiber.Ctx) error {
	req := new(admin.GetUserStatisticListRequest)
	if err := c.QueryParser(req); err != nil {
		return errors.InvalidRequest(err)
	}

	var (
		loginStartTime, loginEndTime   *time.Time
		createStartTime, createEndTime *time.Time
		err                            error
	)

	if req.LastLoginStartTime != nil {
		*loginStartTime, err = time.Parse(time.RFC3339, *req.LastLoginStartTime)
		if err != nil {
			return errors.InvalidRequest(err)
		}
	}
	if req.LastLoginEndTime != nil {
		*loginEndTime, err = time.Parse(time.RFC3339, *req.LastLoginEndTime)
		if err != nil {
			return errors.InvalidRequest(err)
		}
	}
	if req.CreateStartTime != nil {
		*createStartTime, err = time.Parse(time.RFC3339, *req.CreateStartTime)
		if err != nil {
			return errors.InvalidRequest(err)
		}
	}
	if req.CreateEndTime != nil {
		*createEndTime, err = time.Parse(time.RFC3339, *req.CreateEndTime)
		if err != nil {
			return errors.InvalidRequest(err)
		}
	}

	resp, err := s.StatisticService.GetUserStatisticList(
		c.UserContext(), req.Page, req.PageSize, loginStartTime, loginEndTime, createStartTime, createEndTime,
	)
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
