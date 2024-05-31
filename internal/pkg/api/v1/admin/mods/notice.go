package mods

import (
	"fmt"

	"data-collection-hub-server/internal/pkg/config"
	"data-collection-hub-server/internal/pkg/domain/vo"
	"data-collection-hub-server/internal/pkg/domain/vo/admin"
	adminservice "data-collection-hub-server/internal/pkg/service/admin/mods"
	sysservice "data-collection-hub-server/internal/pkg/service/sys/mods"
	"data-collection-hub-server/pkg/errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NoticeApi struct {
	NoticeService adminservice.NoticeService
	LogsService   sysservice.LogsService
	Validator     *validator.Validate
}

func (n *NoticeApi) InsertNotice(c *fiber.Ctx) error {
	ctx := c.UserContext()
	req := new(admin.InsertNoticeRequest)

	if err := c.BodyParser(req); err != nil {
		return errors.InvalidRequest(err)
	}
	if err := n.Validator.Struct(req); err != nil {
		return errors.InvalidParams(err) // Compare this line with the original one
	}

	noticeIDHex, err := n.NoticeService.InsertNotice(ctx, req.Title, req.Content, req.NoticeType)
	var (
		userID, _  = primitive.ObjectIDFromHex(ctx.Value(config.UserIDKey).(string))
		ipAddr     = c.IP()
		userAgent  = c.Get(fiber.HeaderUserAgent)
		operation  = config.OperationTypeCreate
		entityType = config.EntityTypeNotice
	)
	if err != nil {
		var (
			description = fmt.Sprintf("Insert notice failed: %s", err.Error())
			status      = config.OperationStatusFailure
		)
		_ = n.LogsService.CacheOperationLog(
			ctx, &userID, nil, &ipAddr, &userAgent, &operation, &entityType, &description, &status,
		)
		return err
	}

	var (
		noticeID, _ = primitive.ObjectIDFromHex(noticeIDHex)
		description = fmt.Sprintf("Insert notice: %s", noticeIDHex)
		status      = config.OperationStatusSuccess
	)
	_ = n.LogsService.CacheOperationLog(
		ctx, &userID, &noticeID, &ipAddr, &userAgent, &operation, &entityType, &description, &status,
	)
	return c.JSON(
		vo.Response{
			Code:    errors.CodeSuccess,
			Message: errors.MessageSuccess,
			Data:    nil,
		},
	)
}

func (n *NoticeApi) UpdateNotice(c *fiber.Ctx) error {
	ctx := c.UserContext()
	req := new(admin.UpdateNoticeRequest)

	if err := c.BodyParser(req); err != nil {
		return errors.InvalidRequest(err)
	}
	if err := n.Validator.Struct(req); err != nil {
		return errors.InvalidParams(err) // Compare this line with the original one
	}

	noticeID, err := primitive.ObjectIDFromHex(*req.NoticeID)
	if err != nil {
		return errors.InvalidRequest(err)
	}
	err = n.NoticeService.UpdateNotice(ctx, &noticeID, req.Title, req.Content, req.NoticeType)
	var (
		userID, _  = primitive.ObjectIDFromHex(ctx.Value(config.UserIDKey).(string))
		ipAddr     = c.IP()
		userAgent  = c.Get(fiber.HeaderUserAgent)
		operation  = config.OperationTypeUpdate
		entityType = config.EntityTypeNotice
	)
	if err != nil {
		var (
			description = fmt.Sprintf("Update notice failed: %s", err.Error())
			status      = config.OperationStatusFailure
		)
		_ = n.LogsService.CacheOperationLog(
			ctx, &userID, &noticeID, &ipAddr, &userAgent, &operation, &entityType, &description, &status,
		)
		return err
	}

	var (
		description = fmt.Sprintf("Update notice: %s", *req.NoticeID)
		status      = config.OperationStatusSuccess
	)
	_ = n.LogsService.CacheOperationLog(
		ctx, &userID, &noticeID, &ipAddr, &userAgent, &operation, &entityType, &description, &status,
	)

	return c.JSON(
		vo.Response{
			Code:    errors.CodeSuccess,
			Message: errors.MessageSuccess,
			Data:    nil,
		},
	)
}

func (n *NoticeApi) DeleteNotice(c *fiber.Ctx) error {
	ctx := c.UserContext()
	req := new(admin.DeleteNoticeRequest)

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

	var (
		userID, _  = primitive.ObjectIDFromHex(ctx.Value(config.UserIDKey).(string))
		ipAddr     = c.IP()
		userAgent  = c.Get(fiber.HeaderUserAgent)
		operation  = config.OperationTypeDelete
		entityType = config.EntityTypeNotice
	)
	err = n.NoticeService.DeleteNotice(ctx, &noticeID)
	if err != nil {
		var (
			description = fmt.Sprintf("Delete notice failed: %s", err.Error())
			status      = config.OperationStatusFailure
		)
		_ = n.LogsService.CacheOperationLog(
			ctx, &userID, &noticeID, &ipAddr, &userAgent, &operation, &entityType, &description, &status,
		)
		return err
	}

	var (
		description = fmt.Sprintf("Delete notice: %s", *req.NoticeID)
		status      = config.OperationStatusSuccess
	)
	_ = n.LogsService.CacheOperationLog(
		ctx, &userID, &noticeID, &ipAddr, &userAgent, &operation, &entityType, &description, &status,
	)

	return c.JSON(
		vo.Response{
			Code:    errors.CodeSuccess,
			Message: errors.MessageSuccess,
			Data:    nil,
		},
	)
}
