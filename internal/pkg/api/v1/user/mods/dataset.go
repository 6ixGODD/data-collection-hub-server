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
	"data-collection-hub-server/pkg/utils/common"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DatasetApi struct {
	DatasetService userservice.DatasetService
	LogsService    sysservice.LogsService
	Validator      *validator.Validate
}

// InsertInstructionData inserts the instruction data.
//
//	@description	Insert the instruction data.
//	@id				user-insert-instruction-data
//	@summary		insert instruction data
//	@tags			User API
//	@accept			json
//	@produce		json
//	@param			user.InsertInstructionDataRequest	body	user.InsertInstructionDataRequest	true	"Insert instruction data request"
//	@security		Bearer
//	@success		200						{object}	vo.Response{data=nil}	"Success"
//	@failure		400						{object}	vo.Response{data=nil}	"Invalid request"
//	@failure		401						{object}	vo.Response{data=nil}	"Unauthorized"
//	@failure		403						{object}	vo.Response{data=nil}	"Forbidden"
//	@failure		500						{object}	vo.Response{data=nil}	"Internal server error"
//	@router			/user/instruction-data	[post]
func (d *DatasetApi) InsertInstructionData(c *fiber.Ctx) error {
	ctx := c.UserContext()
	req := new(user.InsertInstructionDataRequest)

	if err := c.BodyParser(req); err != nil {
		return errors.InvalidRequest(fmt.Errorf("failed to parse request"))
	}
	if errs := d.Validator.Struct(req); errs != nil {
		return errors.InvalidRequest(common.FormatValidateError(errs))
	}

	userIDHex, ok := ctx.Value(config.UserIDKey).(string)
	if !ok {
		return errors.NotAuthorized(fmt.Errorf("user is not authorized"))
	}
	userID, err := primitive.ObjectIDFromHex(userIDHex)
	if err != nil {
		return errors.NotAuthorized(fmt.Errorf("user is not authorized"))
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

// GetInstructionData returns the instruction data.
//
//	@description	Get the instruction data.
//	@id				user-get-instruction-data
//	@summary		get instruction data
//	@tags			User API
//	@accept			json
//	@produce		json
//	@param			user.GetInstructionDataRequest	query	user.GetInstructionDataRequest	true	"Get instruction data request"
//	@security		Bearer
//	@success		200						{object}	vo.Response{data=user.GetInstructionDataResponse}	"Success"
//	@failure		400						{object}	vo.Response{data=nil}								"Invalid request"
//	@failure		401						{object}	vo.Response{data=nil}								"Unauthorized"
//	@failure		404						{object}	vo.Response{data=nil}								"Instruction data not found"
//	@failure		500						{object}	vo.Response{data=nil}								"Internal server error"
//	@router			/user/instruction-data	[get]
func (d *DatasetApi) GetInstructionData(c *fiber.Ctx) error {
	req := new(user.GetInstructionDataRequest)

	if err := c.QueryParser(req); err != nil {
		return errors.InvalidRequest(fmt.Errorf("failed to parse request"))
	}
	if errs := d.Validator.Struct(req); errs != nil {
		return errors.InvalidRequest(common.FormatValidateError(errs))
	}

	instructionDataID, err := primitive.ObjectIDFromHex(*req.InstructionDataID)
	if err != nil {
		return errors.InvalidRequest(fmt.Errorf("invalid instruciton data id"))
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

// GetInstructionDataList returns the instruction data list.
//
//	@description	Get the instruction data list.
//	@id				user-get-instruction-data-list
//	@summary		get instruction data list
//	@tags			User API
//	@accept			json
//	@produce		json
//	@param			user.GetInstructionDataListRequest	query	user.GetInstructionDataListRequest	true	"Get instruction data list request"
//	@security		Bearer
//	@success		200							{object}	vo.Response{data=user.GetInstructionDataListResponse}	"Success"
//	@failure		400							{object}	vo.Response{data=nil}									"Invalid request"
//	@failure		401							{object}	vo.Response{data=nil}									"Unauthorized"
//	@failure		500							{object}	vo.Response{data=nil}									"Internal server error"
//	@router			/user/instruction-data/list	[get]
func (d *DatasetApi) GetInstructionDataList(c *fiber.Ctx) error {
	req := new(user.GetInstructionDataListRequest)

	if err := c.QueryParser(req); err != nil {
		return errors.InvalidRequest(fmt.Errorf("failed to parse request"))
	}
	if errs := d.Validator.Struct(req); errs != nil {
		return errors.InvalidRequest(common.FormatValidateError(errs))
	}

	var (
		updateBefore, updateAfter       time.Time
		updateBeforePtr, updateAfterPtr *time.Time
		err                             error
	)

	if req.UpdateStartTime != nil && req.UpdateEndTime != nil {
		updateBefore, err = time.Parse(time.RFC3339, *req.UpdateStartTime)
		if err != nil {
			return errors.InvalidRequest(fmt.Errorf("failed to parse request"))
		}
		updateAfter, err = time.Parse(time.RFC3339, *req.UpdateEndTime)
		if err != nil {
			return errors.InvalidRequest(fmt.Errorf("failed to parse request"))
		}
		updateBeforePtr = &updateBefore
		updateAfterPtr = &updateAfter
	}

	resp, err := d.DatasetService.GetInstructionDataList(
		c.UserContext(), req.Page, req.PageSize, updateBeforePtr, updateAfterPtr, req.Theme, req.Status,
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

// UpdateInstructionData updates the instruction data.
//
//	@description	Update the instruction data.
//	@id				user-update-instruction-data
//	@summary		update instruction data
//	@tags			User API
//	@accept			json
//	@produce		json
//	@param			user.UpdateInstructionDataRequest	body	user.UpdateInstructionDataRequest	true	"Update instruction data request"
//	@security		Bearer
//	@success		200						{object}	vo.Response{data=nil}	"Success"
//	@failure		400						{object}	vo.Response{data=nil}	"Invalid request"
//	@failure		401						{object}	vo.Response{data=nil}	"Unauthorized"
//	@failure		403						{object}	vo.Response{data=nil}	"Forbidden"
//	@failure		404						{object}	vo.Response{data=nil}	"Instruction data not found"
//	@failure		500						{object}	vo.Response{data=nil}	"Internal server error"
//	@router			/user/instruction-data	[put]
func (d *DatasetApi) UpdateInstructionData(c *fiber.Ctx) error {
	ctx := c.UserContext()
	req := new(user.UpdateInstructionDataRequest)

	if err := c.BodyParser(req); err != nil {
		return errors.InvalidRequest(fmt.Errorf("failed to parse request"))
	}
	if errs := d.Validator.Struct(req); errs != nil {
		return errors.InvalidRequest(common.FormatValidateError(errs))
	}

	instructionDataID, err := primitive.ObjectIDFromHex(*req.InstructionDataID)
	if err != nil {
		return errors.InvalidRequest(fmt.Errorf("invalid instruction data id"))
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

// DeleteInstructionData deletes the instruction data.
//
//	@description	Delete the instruction data.
//	@id				user-delete-instruction-data
//	@summary		delete instruction data
//	@tags			User API
//	@accept			json
//	@produce		json
//	@param			user.DeleteInstructionDataRequest	query	user.DeleteInstructionDataRequest	true	"Delete instruction data request"
//	@security		Bearer
//	@success		200						{object}	vo.Response{data=nil}	"Success"
//	@failure		400						{object}	vo.Response{data=nil}	"Invalid request"
//	@failure		401						{object}	vo.Response{data=nil}	"Unauthorized"
//	@failure		403						{object}	vo.Response{data=nil}	"Forbidden"
//	@failure		404						{object}	vo.Response{data=nil}	"Instruction data not found"
//	@failure		500						{object}	vo.Response{data=nil}	"Internal server error"
//	@router			/user/instruction-data	[delete]
func (d *DatasetApi) DeleteInstructionData(c *fiber.Ctx) error {
	ctx := c.UserContext()
	req := new(user.DeleteInstructionDataRequest)

	if err := c.QueryParser(req); err != nil {
		return errors.InvalidRequest(fmt.Errorf("failed to parse request"))
	}
	if errs := d.Validator.Struct(req); errs != nil {
		return errors.InvalidRequest(common.FormatValidateError(errs))
	}

	instructionDataID, err := primitive.ObjectIDFromHex(*req.InstructionDataID)
	if err != nil {
		return errors.InvalidRequest(fmt.Errorf("invalid instruction data id"))
	}

	userIDHex, ok := ctx.Value(config.UserIDKey).(string)
	if !ok {
		return errors.NotAuthorized(fmt.Errorf("user is not authorized"))
	}
	userID, err := primitive.ObjectIDFromHex(userIDHex)
	if err != nil {
		return errors.NotAuthorized(fmt.Errorf("user is not authorized"))
	}
	entityID, err := primitive.ObjectIDFromHex(*req.InstructionDataID)
	if err != nil {
		return errors.InvalidRequest(fmt.Errorf("invalid instruction data id"))
	}
	var (
		ipAddr     = c.IP()
		userAgent  = c.Get(fiber.HeaderUserAgent)
		operation  = config.OperationTypeDelete
		entityType = config.EntityTypeInstruction
	)

	err = d.DatasetService.DeleteInstructionData(ctx, &instructionDataID)
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
