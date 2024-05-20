package mods

import (
	"time"

	"data-collection-hub-server/internal/pkg/schema"
	"data-collection-hub-server/internal/pkg/schema/common"
	commonservice "data-collection-hub-server/internal/pkg/service/common/mods"
	"data-collection-hub-server/pkg/errors"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NoticeApi struct {
	commonservice.NoticeService
}

func (n *NoticeApi) GetNotice(c *fiber.Ctx) error {
	req := new(common.GetNoticeRequest)

	if err := c.QueryParser(req); err != nil {
		return errors.InvalidRequest(err)
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
		schema.Response{
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

	var (
		updateBefore, updateAfter *time.Time
		err                       error
	)

	if req.UpdateStartTime != nil {
		*updateBefore, err = time.Parse(time.RFC3339, *req.UpdateStartTime)
		if err != nil {
			return errors.InvalidRequest(err)
		}
	}
	if req.UpdateEndTime != nil {
		*updateAfter, err = time.Parse(time.RFC3339, *req.UpdateEndTime)
		if err != nil {
			return errors.InvalidRequest(err)
		}
	}

	resp, err := n.NoticeService.GetNoticeList(
		c.UserContext(), req.Page, req.PageSize, req.NoticeType, updateBefore, updateAfter,
	)
	if err != nil {
		return err
	}
	return c.JSON(resp)
}
