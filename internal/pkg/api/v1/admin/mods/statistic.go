package mods

import (
	"time"

	"data-collection-hub-server/internal/pkg/schema"
	"data-collection-hub-server/internal/pkg/schema/admin"
	adminservice "data-collection-hub-server/internal/pkg/service/admin/mods"
	"data-collection-hub-server/pkg/errors"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StatisticApi struct {
	adminservice.StatisticService
}

func NewStatisticApi(statisticService adminservice.StatisticService) *StatisticApi {
	return &StatisticApi{statisticService}
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

func (s *StatisticApi) GetUserStatistic(c *fiber.Ctx) error {
	req := new(admin.GetUserStatisticRequest)
	if err := c.QueryParser(req); err != nil {
		return errors.InvalidRequest(err)
	}

	userID, err := primitive.ObjectIDFromHex(*req.UserID)
	if err != nil {
		return errors.InvalidRequest(err)
	}

	resp, err := s.StatisticService.GetUserStatistic(c.Context(), &userID)
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

func (s *StatisticApi) GetUserStatisticList(c *fiber.Ctx) error {
	req := new(admin.GetUserStatisticListRequest)
	if err := c.QueryParser(req); err != nil {
		return errors.InvalidRequest(err)
	}

	var (
		page                        *int
		loginBefore, loginAfter     *time.Time
		createdBefore, createdAfter *time.Time
		err                         error
	)

	if req.Page != nil {
		page = req.Page
	}
	if req.LastLoginBefore != nil {
		*loginBefore, err = time.Parse(time.RFC3339, *req.LastLoginBefore)
		if err != nil {
			return errors.InvalidRequest(err)
		}
	}
	if req.LastLoginAfter != nil {
		*loginAfter, err = time.Parse(time.RFC3339, *req.LastLoginAfter)
		if err != nil {
			return errors.InvalidRequest(err)
		}
	}
	if req.CreatedBefore != nil {
		*createdBefore, err = time.Parse(time.RFC3339, *req.CreatedBefore)
		if err != nil {
			return errors.InvalidRequest(err)
		}
	}
	if req.CreatedAfter != nil {
		*createdAfter, err = time.Parse(time.RFC3339, *req.CreatedAfter)
		if err != nil {
			return errors.InvalidRequest(err)
		}
	}

	resp, err := s.StatisticService.GetUserStatisticList(
		c.Context(), page, loginBefore, loginAfter, createdBefore, createdAfter,
	)
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
