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

func NewLogsApi(logsService adminservice.LogsService) LogsApi {
	return LogsApi{logsService}
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

	resp, err := l.LogsService.GetLoginLog(c.Context(), &loginLogID)
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

	resp, err := l.LogsService.GetLoginLogList(c.Context(), req.Page, req.Query, createdBefore, createdAfter)
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

	resp, err := l.LogsService.GetOperationLog(c.Context(), &operationLogID)
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

	resp, err := l.LogsService.GetOperationLogList(
		c.Context(), req.Page, req.Query, req.Operation, createdBefore, createdAfter,
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

func (l *LogsApi) GetErrorLog(c *fiber.Ctx) error {
	req := new(admin.GetErrorLogRequest)

	if err := c.QueryParser(req); err != nil {
		return errors.InvalidRequest(err)
	}

	errorLogID, err := primitive.ObjectIDFromHex(*req.ErrorLogID)
	if err != nil {
		return errors.InvalidRequest(err)
	}

	resp, err := l.LogsService.GetErrorLog(c.Context(), &errorLogID)
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

func (l *LogsApi) GetErrorLogList(c *fiber.Ctx) error {
	req := new(admin.GetErrorLogListRequest)

	if err := c.QueryParser(req); err != nil {
		return errors.InvalidRequest(err)
	}

	var (
		createdBefore, createdAfter *time.Time
		err                         error
	)

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

	resp, err := l.LogsService.GetErrorLogList(
		c.Context(), req.Page, req.RequestURL, req.ErrorCode, createdBefore, createdAfter,
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
