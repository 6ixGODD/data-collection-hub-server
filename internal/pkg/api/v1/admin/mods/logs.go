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

type LogsApi struct {
	adminservice.LogsService
}

func (l *LogsApi) GetLoginLog(c *fiber.Ctx) error {
	req := new(admin.GetLoginLogRequest)

	if err := c.QueryParser(req); err != nil {
		return errors.InvalidRequest(err)
	}

	loginLogID, err := primitive.ObjectIDFromHex(*req.LoginLogID)
	if err != nil {
		return errors.InvalidRequest(err)
	}

	resp, err := l.LogsService.GetLoginLog(c.UserContext(), &loginLogID)
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

func (l *LogsApi) GetLoginLogList(c *fiber.Ctx) error {
	req := new(admin.GetLoginLogListRequest)

	if err := c.QueryParser(req); err != nil {
		return errors.InvalidRequest(err)
	}

	var (
		createdBefore, createdAfter *time.Time
		err                         error
	)

	if req.CreateStartTime != nil {
		*createdBefore, err = time.Parse(time.RFC3339, *req.CreateStartTime)
		if err != nil {
			return errors.InvalidRequest(err)
		}
	}
	if req.CreateEndTime != nil {
		*createdAfter, err = time.Parse(time.RFC3339, *req.CreateEndTime)
		if err != nil {
			return errors.InvalidRequest(err)
		}
	}

	resp, err := l.LogsService.GetLoginLogList(
		c.UserContext(), req.Page, req.PageSize, req.Desc, req.Query, createdBefore, createdAfter,
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

func (l *LogsApi) GetOperationLog(c *fiber.Ctx) error {
	req := new(admin.GetOperationLogRequest)

	if err := c.QueryParser(req); err != nil {
		return errors.InvalidRequest(err)
	}

	operationLogID, err := primitive.ObjectIDFromHex(*req.OperationLogID)
	if err != nil {
		return errors.InvalidRequest(err)
	}

	resp, err := l.LogsService.GetOperationLog(c.UserContext(), &operationLogID)
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

func (l *LogsApi) GetOperationLogList(c *fiber.Ctx) error {
	req := new(admin.GetOperationLogListRequest)

	if err := c.QueryParser(req); err != nil {
		return errors.InvalidRequest(err)
	}

	var (
		createdBefore, createdAfter *time.Time
		err                         error
	)

	if req.CreateStartTime != nil {
		*createdBefore, err = time.Parse(time.RFC3339, *req.CreateStartTime)
		if err != nil {
			return errors.InvalidRequest(err)
		}
	}
	if req.CreateEndTime != nil {
		*createdAfter, err = time.Parse(time.RFC3339, *req.CreateEndTime)
		if err != nil {
			return errors.InvalidRequest(err)
		}
	}

	resp, err := l.LogsService.GetOperationLogList(
		c.UserContext(), req.Page, req.PageSize, req.Desc, req.Query, req.Operation, createdBefore, createdAfter,
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
