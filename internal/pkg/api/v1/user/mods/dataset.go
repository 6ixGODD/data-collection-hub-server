package mods

import (
	"fmt"
	"time"

	"data-collection-hub-server/internal/pkg/config"
	"data-collection-hub-server/internal/pkg/domain/vo"
	"data-collection-hub-server/internal/pkg/domain/vo/user"
	sysservice "data-collection-hub-server/internal/pkg/service/sys/mods"
	userservice "data-collection-hub-server/internal/pkg/service/user/mods"
	"data-collection-hub-server/pkg/errors"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DatasetApi struct {
	DatasetService userservice.DatasetService
	LogsService    sysservice.LogsService
}

func (d *DatasetApi) InsertInstructionData(c *fiber.Ctx) error {
	ctx := c.UserContext()
	req := new(user.InsertInstructionDataRequest)

	if err := c.BodyParser(req); err != nil {
		return errors.InvalidRequest(err)
	}

	userID, err := primitive.ObjectIDFromHex(ctx.Value(config.UserIDKey).(string))
	if err != nil {
		return errors.InvalidRequest(err) // TODO: Change to not login error
	}

	instructionDataIDHex, err := d.DatasetService.InsertInstructionData(
		ctx, req.Instruction, req.Input, req.Output, req.Theme, req.Source, req.Note,
	)
	var (
		entityID, _ = primitive.ObjectIDFromHex(instructionDataIDHex)
		ipAddr      = c.IP()
		userAgent   = c.Get(fiber.HeaderUserAgent)
		operation   = config.OperationTypeCreate
		entityType  = config.EntityTypeInstruction
	)

	if err != nil {
		var (
			description = fmt.Sprintf("Insert instruction data failed: %s", err.Error())
			status      = config.OperationStatusFailure
		)
		_ = d.LogsService.CacheOperationLog(
			ctx, &userID, &entityID, &ipAddr, &userAgent, &operation, &entityType, &description, &status,
		)
		return err
	}

	var (
		description = fmt.Sprintf("Insert instruction data: %s", instructionDataIDHex)
		status      = config.OperationStatusSuccess
	)
	_ = d.LogsService.CacheOperationLog(
		ctx, &userID, &entityID, &ipAddr, &userAgent, &operation, &entityType, &description, &status,
	)

	return c.JSON(
		vo.Response{
			Code:    errors.CodeSuccess,
			Message: errors.MessageSuccess,
			Data:    nil,
		},
	)
}

func (d *DatasetApi) GetInstructionData(c *fiber.Ctx) error {
	req := new(user.GetInstructionDataRequest)

	if err := c.QueryParser(req); err != nil {
		return errors.InvalidRequest(err)
	}

	instructionDataID, err := primitive.ObjectIDFromHex(*req.InstructionDataID)
	if err != nil {
		return errors.InvalidRequest(err)
	}

	resp, err := d.DatasetService.GetInstructionData(c.UserContext(), instructionDataID)
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

func (d *DatasetApi) GetInstructionDataList(c *fiber.Ctx) error {
	req := new(user.GetInstructionDataListRequest)

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

	resp, err := d.DatasetService.GetInstructionDataList(
		c.UserContext(), req.Page, req.PageSize, updateBefore, updateAfter, req.Theme, req.Status,
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

func (d *DatasetApi) UpdateInstructionData(c *fiber.Ctx) error {
	ctx := c.UserContext()
	req := new(user.UpdateInstructionDataRequest)

	if err := c.BodyParser(req); err != nil {
		return errors.InvalidRequest(err)
	}

	instructionDataID, err := primitive.ObjectIDFromHex(*req.InstructionDataID)
	if err != nil {
		return errors.InvalidRequest(err)
	}

	err = d.DatasetService.UpdateInstructionData(
		ctx, &instructionDataID, req.Instruction, req.Input, req.Output, req.Theme, req.Source, req.Note,
	)
	var (
		userID, _   = primitive.ObjectIDFromHex(ctx.Value(config.UserIDKey).(string))
		entityID, _ = primitive.ObjectIDFromHex(*req.InstructionDataID)
		ipAddr      = c.IP()
		userAgent   = c.Get(fiber.HeaderUserAgent)
		operation   = config.OperationTypeUpdate
		entityType  = config.EntityTypeInstruction
	)

	if err != nil {
		var (
			description = fmt.Sprintf("Update instruction data failed: %s", err.Error())
			status      = config.OperationStatusFailure
		)
		_ = d.LogsService.CacheOperationLog(
			ctx, &userID, &entityID, &ipAddr, &userAgent, &operation, &entityType, &description, &status,
		)
		return err
	}

	var (
		description = fmt.Sprintf("Update instruction data: %s", *req.InstructionDataID)
		status      = config.OperationStatusSuccess
	)
	_ = d.LogsService.CacheOperationLog(
		ctx, &userID, &entityID, &ipAddr, &userAgent, &operation, &entityType, &description, &status,
	)
	return c.JSON(
		vo.Response{
			Code:    errors.CodeSuccess,
			Message: errors.MessageSuccess,
			Data:    nil,
		},
	)
}

func (d *DatasetApi) DeleteInstructionData(c *fiber.Ctx) error {
	ctx := c.UserContext()
	req := new(user.DeleteInstructionDataRequest)

	if err := c.QueryParser(req); err != nil {
		return errors.InvalidRequest(err)
	}

	instructionDataID, err := primitive.ObjectIDFromHex(*req.InstructionDataID)
	if err != nil {
		return errors.InvalidRequest(err)
	}

	err = d.DatasetService.DeleteInstructionData(ctx, &instructionDataID)
	var (
		userID, _   = primitive.ObjectIDFromHex(ctx.Value(config.UserIDKey).(string))
		entityID, _ = primitive.ObjectIDFromHex(*req.InstructionDataID)
		ipAddr      = c.IP()
		userAgent   = c.Get(fiber.HeaderUserAgent)
		operation   = config.OperationTypeDelete
		entityType  = config.EntityTypeInstruction
	)

	if err != nil {
		var (
			description = fmt.Sprintf("Delete instruction data failed: %s", err.Error())
			status      = config.OperationStatusFailure
		)
		_ = d.LogsService.CacheOperationLog(
			ctx, &userID, &entityID, &ipAddr, &userAgent, &operation, &entityType, &description, &status,
		)
		return err
	}

	var (
		description = fmt.Sprintf("Delete instruction data: %s", *req.InstructionDataID)
		status      = config.OperationStatusSuccess
	)
	_ = d.LogsService.CacheOperationLog(
		ctx, &userID, &entityID, &ipAddr, &userAgent, &operation, &entityType, &description, &status,
	)

	return c.JSON(
		vo.Response{
			Code:    errors.CodeSuccess,
			Message: errors.MessageSuccess,
			Data:    nil,
		},
	)
}
