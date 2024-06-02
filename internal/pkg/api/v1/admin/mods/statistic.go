package mods

import (
	"fmt"
	"time"

	"data-collection-hub-server/internal/pkg/domain/vo"
	"data-collection-hub-server/internal/pkg/domain/vo/admin"
	adminservice "data-collection-hub-server/internal/pkg/service/admin/mods"
	"data-collection-hub-server/pkg/errors"
	"data-collection-hub-server/pkg/utils/common"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StatisticApi struct {
	StatisticService adminservice.StatisticService
	Validator        *validator.Validate
}

// GetDataStatistic returns the data statistic.
//
//	@description	Get the data statistic.
//	@id				admin-get-data-statistic
//	@summary		get data statistic
//	@tags			Admin API
//	@accept			json
//	@produce		json
//	@param			admin.GetDataStatisticRequest	query	admin.GetDataStatisticRequest	true	"Get data statistic request"
//	@security		Bearer
//	@success		200						{object}	vo.Response{data=admin.GetDataStatisticResponse}	"Success"
//	@failure		400						{object}	vo.Response{data=nil}								"Invalid request"
//	@failure		401						{object}	vo.Response{data=nil}								"Unauthorized"
//	@failure		403						{object}	vo.Response{data=nil}								"Forbidden"
//	@failure		404						{object}	vo.Response{data=nil}								"Data statistic not found"
//	@failure		500						{object}	vo.Response{data=nil}								"Internal server error"
//	@router			/admin/data-statistic	[get]
func (s *StatisticApi) GetDataStatistic(c *fiber.Ctx) error {
	req := new(admin.GetDataStatisticRequest)
	if err := c.QueryParser(req); err != nil {
		return errors.InvalidRequest(fmt.Errorf("failed to parse request"))
	}
	if errs := s.Validator.Struct(req); errs != nil {
		return errors.InvalidRequest(common.FormatValidateError(errs))
	}

	// Parse start and end date
	var startDate, endDate *time.Time
	if req.StartDate != nil {
		*startDate, _ = time.Parse(time.RFC3339, *req.StartDate)
	}
	if req.EndDate != nil {
		*endDate, _ = time.Parse(time.RFC3339, *req.EndDate)
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

// GetUserStatistic returns the user statistic.
//
//	@description	Get the user statistic.
//	@id				admin-get-user-statistic
//	@summary		get user statistic
//	@tags			Admin API
//	@accept			json
//	@produce		json
//	@param			admin.GetUserStatisticRequest	query	admin.GetUserStatisticRequest	true	"Get user statistic request"
//	@security		Bearer
//	@success		200						{object}	vo.Response{data=admin.GetUserStatisticResponse}	"Success"
//	@failure		400						{object}	vo.Response{data=nil}								"Invalid request"
//	@failure		401						{object}	vo.Response{data=nil}								"Unauthorized"
//	@failure		403						{object}	vo.Response{data=nil}								"Forbidden"
//	@failure		404						{object}	vo.Response{data=nil}								"User statistic not found"
//	@failure		500						{object}	vo.Response{data=nil}								"Internal server error"
//	@router			/admin/user-statistic	[get]
func (s *StatisticApi) GetUserStatistic(c *fiber.Ctx) error {
	req := new(admin.GetUserStatisticRequest)
	if err := c.QueryParser(req); err != nil {
		return errors.InvalidRequest(fmt.Errorf("failed to parse request"))
	}
	if errs := s.Validator.Struct(req); errs != nil {
		return errors.InvalidRequest(common.FormatValidateError(errs))
	}

	userID, err := primitive.ObjectIDFromHex(*req.UserID)
	if err != nil {
		return errors.InvalidRequest(fmt.Errorf("invalid user id"))
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

// GetUserStatisticList returns the user statistic list.
//
//	@description	Get the user statistic list.
//	@id				admin-get-user-statistic-list
//	@summary		get user statistic list
//	@tags			Admin API
//	@accept			json
//	@produce		json
//	@param			admin.GetUserStatisticListRequest	query	admin.GetUserStatisticListRequest	true	"Get user statistic list request"
//	@security		Bearer
//	@success		200							{object}	vo.Response{data=admin.GetUserStatisticListResponse}	"Success"
//	@failure		400							{object}	vo.Response{data=nil}									"Invalid request"
//	@failure		401							{object}	vo.Response{data=nil}									"Unauthorized"
//	@failure		403							{object}	vo.Response{data=nil}									"Forbidden"
//	@failure		500							{object}	vo.Response{data=nil}									"Internal server error"
//	@router			/admin/user-statistic/list	[get]
func (s *StatisticApi) GetUserStatisticList(c *fiber.Ctx) error {
	req := new(admin.GetUserStatisticListRequest)
	if err := c.QueryParser(req); err != nil {
		return errors.InvalidRequest(fmt.Errorf("failed to parse request"))
	}
	if errs := s.Validator.Struct(req); errs != nil {
		return errors.InvalidRequest(common.FormatValidateError(errs))
	}

	var (
		loginStartTime, loginEndTime         time.Time
		loginStartTimePtr, loginEndTimePtr   *time.Time
		createStartTime, createEndTime       time.Time
		createStartTimePtr, createEndTimePtr *time.Time
		err                                  error
	)

	if req.LastLoginStartTime != nil && req.LastLoginEndTime != nil {
		loginStartTime, err = time.Parse(time.RFC3339, *req.LastLoginStartTime)
		if err != nil {
			return errors.InvalidRequest(
				fmt.Errorf(
					"invalid last login start time %s (should be in RFC3339 format)", *req.LastLoginStartTime,
				),
			)
		}
		loginEndTime, err = time.Parse(time.RFC3339, *req.LastLoginEndTime)
		if err != nil {
			return errors.InvalidRequest(
				fmt.Errorf(
					"invalid last login end time %s (should be in RFC3339 format)", *req.LastLoginEndTime,
				),
			)
		}
		loginStartTimePtr = &loginStartTime
		loginEndTimePtr = &loginEndTime
	}
	if req.CreateStartTime != nil && req.CreateEndTime != nil {
		createStartTime, err = time.Parse(time.RFC3339, *req.CreateStartTime)
		if err != nil {
			return errors.InvalidRequest(
				fmt.Errorf(
					"invalid create start time %s (should be in RFC3339 format)", *req.CreateStartTime,
				),
			)
		}
		createEndTime, err = time.Parse(time.RFC3339, *req.CreateEndTime)
		if err != nil {
			return errors.InvalidRequest(
				fmt.Errorf(
					"invalid create end time %s (should be in RFC3339 format)", *req.CreateEndTime,
				),
			)
		}
		createStartTimePtr = &createStartTime
		createEndTimePtr = &createEndTime
	}

	resp, err := s.StatisticService.GetUserStatisticList(
		c.UserContext(), req.Page, req.PageSize, loginStartTimePtr, loginEndTimePtr, createStartTimePtr,
		createEndTimePtr,
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
