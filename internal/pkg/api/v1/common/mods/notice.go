package mods

import (
	"time"

	"data-collection-hub-server/internal/pkg/domain/vo"
	"data-collection-hub-server/internal/pkg/domain/vo/common"
	commonservice "data-collection-hub-server/internal/pkg/service/common/mods"
	"data-collection-hub-server/pkg/errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NoticeApi struct {
	NoticeService commonservice.NoticeService
	Validator     *validator.Validate
}

func (n *NoticeApi) GetNotice(c *fiber.Ctx) error {
	req := new(common.GetNoticeRequest)

	if err := c.QueryParser(req); err != nil {
		return errors.InvalidRequest(err)
	}
	if err := n.Validator.Struct(req); err != nil {
		return errors.InvalidParams(err) // Compare this line with the original one
	}

	noticeID, err := primitive.ObjectIDFromHex(*req.NoticeID)
	if err != nil {
		return errors.InvalidRequest(err)
	}
	resp, err := n.NoticeService.GetNotice(c.UserContext(), &noticeID)
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

func (n *NoticeApi) GetNoticeList(c *fiber.Ctx) error {
	req := new(common.GetNoticeListRequest)

	if err := c.QueryParser(req); err != nil {
		return errors.InvalidRequest(err)
	}
	if err := n.Validator.Struct(req); err != nil {
		return errors.InvalidParams(err) // Compare this line with the original one
	}

	var (
		updateStartTime, updateEndTime time.Time
		err                            error
	)

	if req.UpdateStartTime != nil {
		updateStartTime, err = time.Parse(time.RFC3339, *req.UpdateStartTime)
		if err != nil {
			return errors.InvalidRequest(err)
		}
	}
	if req.UpdateEndTime != nil {
		updateEndTime, err = time.Parse(time.RFC3339, *req.UpdateEndTime)
		if err != nil {
			return errors.InvalidRequest(err)
		}
	}

	resp, err := n.NoticeService.GetNoticeList(
		c.UserContext(), req.Page, req.PageSize, req.NoticeType, &updateStartTime, &updateEndTime,
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
