package mods

import (
	"fmt"
	"time"

	"data-collection-hub-server/internal/pkg/domain/vo"
	"data-collection-hub-server/internal/pkg/domain/vo/user"
	userservice "data-collection-hub-server/internal/pkg/service/user/mods"
	"data-collection-hub-server/pkg/errors"
	"data-collection-hub-server/pkg/utils/common"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type StatisticApi struct {
	StatisticService userservice.StatisticService
	Validator        *validator.Validate
}

// GetDataStatistic returns the data statistic.
//
//	@description	Get the data statistic.
//	@id				user-get-data-statistic
//	@summary		get data statistic
//	@tags			User API
//	@accept			json
//	@produce		json
//	@param			user.GetDataStatisticRequest	query	user.GetDataStatisticRequest	false	"Get data statistic request"
//	@security		Bearer
//	@success		200						{object}	vo.Response{data=user.GetDataStatisticResponse}	"Success"
//	@failure		400						{object}	vo.Response{data=nil}							"Invalid request"
//	@failure		401						{object}	vo.Response{data=nil}							"Unauthorized"
//	@failure		500						{object}	vo.Response{data=nil}							"Internal server error"
//	@router			/user/data-statistic	[get]
func (s *StatisticApi) GetDataStatistic(c *fiber.Ctx) error {
	req := new(user.GetDataStatisticRequest)
	if err := c.QueryParser(req); err != nil {
		return errors.InvalidRequest(fmt.Errorf("failed to parse request"))
	}
	if errs := s.Validator.Struct(req); errs != nil {
		return errors.InvalidRequest(common.FormatValidateError(errs))
	}

	var (
		startDate, endDate       time.Time
		startDatePtr, endDatePtr *time.Time
		err                      error
	)
	if req.StartDate != nil {
		startDate, err = time.Parse(time.RFC3339, *req.StartDate)
		if err != nil {
			return errors.InvalidRequest(
				fmt.Errorf(
					"invalid start date %s (should be in RFC3339 format)", *req.StartDate,
				),
			)
		}
		startDatePtr = &startDate
	}
	if req.EndDate != nil {
		endDate, err = time.Parse(time.RFC3339, *req.EndDate)
		if err != nil {
			return errors.InvalidRequest(fmt.Errorf("invalid end date %s (should be in RFC3339 format)", *req.EndDate))
		}
		endDatePtr = &endDate
	}

	resp, err := s.StatisticService.GetDataStatistic(c.UserContext(), startDatePtr, endDatePtr)
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
