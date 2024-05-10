package mods

import (
	"data-collection-hub-server/internal/pkg/schema"
	"data-collection-hub-server/internal/pkg/schema/admin"
	adminservice "data-collection-hub-server/internal/pkg/service/admin/mods"
	"data-collection-hub-server/pkg/errors"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NoticeApi struct {
	adminservice.NoticeService
}

func NewNoticeApi(noticeService adminservice.NoticeService) NoticeApi {
	return NoticeApi{noticeService}
}

func (api *NoticeApi) InsertNotice(c *fiber.Ctx) error {
	req := new(admin.InsertNoticeRequest)

	if err := c.BodyParser(req); err != nil {
		return errors.InvalidRequest(err)
	}

	err := api.NoticeService.InsertNotice(c.Context(), req.Title, req.Content)
	if err != nil {
		return err
	}

	return c.JSON(
		schema.Response{
			Code:    errors.CodeSuccess,
			Message: errors.MessageSuccess,
			Data:    nil,
		},
	)
}

func (api *NoticeApi) UpdateNotice(c *fiber.Ctx) error {
	req := new(admin.UpdateNoticeRequest)

	if err := c.BodyParser(req); err != nil {
		return errors.InvalidRequest(err)
	}

	noticeID, err := primitive.ObjectIDFromHex(*req.NoticeID)
	if err != nil {
		return errors.InvalidRequest(err)
	}
	err = api.NoticeService.UpdateNotice(c.Context(), &noticeID, req.Title, req.Content)
	if err != nil {
		return err
	}

	return c.JSON(
		schema.Response{
			Code:    errors.CodeSuccess,
			Message: errors.MessageSuccess,
			Data:    nil,
		},
	)
}

func (api *NoticeApi) DeleteNotice(c *fiber.Ctx) error {
	req := new(admin.DeleteNoticeRequest)

	if err := c.QueryParser(req); err != nil {
		return errors.InvalidRequest(err)
	}

	noticeID, err := primitive.ObjectIDFromHex(*req.NoticeID)
	if err != nil {
		return errors.InvalidRequest(err)
	}
	err = api.NoticeService.DeleteNotice(c.Context(), &noticeID)
	if err != nil {
		return err
	}

	return c.JSON(
		schema.Response{
			Code:    errors.CodeSuccess,
			Message: errors.MessageSuccess,
			Data:    nil,
		},
	)
}
