package mods

import (
	"fmt"
	"time"

	"data-collection-hub-server/internal/pkg/config"
	"data-collection-hub-server/internal/pkg/domain/vo"
	"data-collection-hub-server/internal/pkg/domain/vo/admin"
	adminservice "data-collection-hub-server/internal/pkg/service/admin/mods"
	sysservice "data-collection-hub-server/internal/pkg/service/sys/mods"
	"data-collection-hub-server/pkg/errors"
	"data-collection-hub-server/pkg/utils/common"
	"github.com/go-playground/validator/v10"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DataAuditApi struct {
	DataAuditService adminservice.DataAuditService
	LogsService      sysservice.LogsService
	Validator        *validator.Validate
}

// GetInstructionData returns the instruction data by ID.
//
//	@description	Get the instruction data by ID.
//	@id				admin-get-instruction-data
//	@summary		get instruction data by ID
//	@tags			Admin API
//	@accept			json
//	@produce		json
//	@param			admin.GetInstructionDataRequest	query	admin.GetInstructionDataRequest	true	"Get instruction data request"
//	@security		Bearer
//	@success		200						{object}	vo.Response{data=admin.GetInstructionDataResponse}	"Success"
//	@failure		400						{object}	vo.Response{data=nil}								"Invalid request"
//	@failure		401						{object}	vo.Response{data=nil}								"Unauthorized"
//	@failure		403						{object}	vo.Response{data=nil}								"Forbidden"
//	@failure		404						{object}	vo.Response{data=nil}								"Instruction data not found"
//	@failure		500						{object}	vo.Response{data=nil}								"Internal server error"
//	@router			/admin/instruction-data	[get]
func (d *DataAuditApi) GetInstructionData(c *fiber.Ctx) error {
	req := new(admin.GetInstructionDataRequest)

	if err := c.QueryParser(req); err != nil {
		return errors.InvalidRequest(fmt.Errorf("failed to parse request"))
	}
	if errs := d.Validator.Struct(req); errs != nil {
		return errors.InvalidRequest(common.FormatValidateError(errs))
	}
	instructionDataID, err := primitive.ObjectIDFromHex(*req.InstructionDataID)
	if err != nil {
		return errors.InvalidRequest(fmt.Errorf("invalid instruction data ID"))
	}
	resp, err := d.DataAuditService.GetInstructionData(c.UserContext(), instructionDataID)
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
//	@id				admin-get-instruction-data-list
//	@summary		get instruction data list
//	@tags			Admin API
//	@accept			json
//	@produce		json
//	@param			admin.GetInstructionDataListRequest	query	admin.GetInstructionDataListRequest	true	"Get instruction data list request"
//	@security		Bearer
//	@success		200	{object}	vo.Response{data=admin.GetInstructionDataListResponse}	"Success"
//	@failure		400	{object}	vo.Response{data=nil}									"Invalid request"
//	@failure		401	{object}	vo.Response{data=nil}									"Unauthorized"
//	@failure		403	{object}	vo.Response{data=nil}									"Forbidden"
//	@failure		500	{object}	vo.Response{data=nil}									"Internal server error"
//	@router			/admin/instruction-data/list [get]
func (d *DataAuditApi) GetInstructionDataList(c *fiber.Ctx) error {
	req := new(admin.GetInstructionDataListRequest)

	if err := c.QueryParser(req); err != nil {
		return errors.InvalidRequest(fmt.Errorf("failed to parse request"))
	}

	if errs := d.Validator.Struct(req); errs != nil {
		return errors.InvalidRequest(common.FormatValidateError(errs))
	}

	var (
		userID                               primitive.ObjectID
		userIDPtr                            *primitive.ObjectID
		createStartTime, createEndTime       time.Time
		updateStartTime, updateEndTime       time.Time
		createStartTimePtr, createEndTimePtr *time.Time
		updateStartTimePtr, updateEndTimePtr *time.Time
		err                                  error
	)

	if req.UserID != nil {
		userID, err = primitive.ObjectIDFromHex(*req.UserID)
		if err != nil {
			return errors.InvalidRequest(fmt.Errorf("invalid user ID"))
		}
		userIDPtr = &userID
	}
	if req.CreateStartTime != nil && req.CreateEndTime != nil {
		createStartTime, err = time.Parse(time.RFC3339, *req.CreateStartTime)
		if err != nil {
			return errors.InvalidRequest(
				fmt.Errorf(
					"invalid create start time %s (should be in `RFC3339` format)", *req.CreateStartTime,
				),
			)
		}
		createEndTime, err = time.Parse(time.RFC3339, *req.CreateEndTime)
		if err != nil {
			return errors.InvalidRequest(
				fmt.Errorf(
					"invalid create end time %s (should be in `RFC3339` format)", *req.CreateEndTime,
				),
			)
		}
		createStartTimePtr = &createStartTime
		createEndTimePtr = &createEndTime
	}
	if req.UpdateStartTime != nil && req.UpdateEndTime != nil {
		updateStartTime, err = time.Parse(time.RFC3339, *req.UpdateStartTime)
		if err != nil {
			return errors.InvalidRequest(
				fmt.Errorf(
					"invalid update start time %s (should be in `RFC3339` format)", *req.UpdateStartTime,
				),
			)
		}
		updateEndTime, err = time.Parse(time.RFC3339, *req.UpdateEndTime)
		if err != nil {
			return errors.InvalidRequest(
				fmt.Errorf(
					"invalid update end time %s (should be in `RFC3339` format)", *req.UpdateEndTime,
				),
			)
		}
		updateStartTimePtr = &updateStartTime
		updateEndTimePtr = &updateEndTime
	}
	resp, err := d.DataAuditService.GetInstructionDataList(
		c.UserContext(),
		req.Page, req.PageSize, req.Desc, userIDPtr, createStartTimePtr, createEndTimePtr, updateStartTimePtr,
		updateEndTimePtr, req.Theme, req.Status, req.Query,
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

// ApproveInstructionData approves the instruction data.
//
//	@description	Approve the instruction data.
//	@id				admin-approve-instruction-data
//	@summary		approve instruction data
//	@tags			Admin API
//	@accept			json
//	@param			admin.ApproveInstructionDataRequest	body	admin.ApproveInstructionDataRequest	true	"Approve instruction data request"
//	@security		Bearer
//	@success		200	{object}	vo.Response{data=nil}	"Success"
//	@failure		400	{object}	vo.Response{data=nil}	"Invalid request"
//	@failure		401	{object}	vo.Response{data=nil}	"Unauthorized"
//	@failure		403	{object}	vo.Response{data=nil}	"Forbidden"
//	@failure		500	{object}	vo.Response{data=nil}	"Internal server error"
//	@router			/admin/instruction-data/approve [post]
func (d *DataAuditApi) ApproveInstructionData(c *fiber.Ctx) error {
	req := new(admin.ApproveInstructionDataRequest)

	if err := c.BodyParser(req); err != nil {
		return errors.InvalidRequest(fmt.Errorf("failed to parse request"))
	}

	if errs := d.Validator.Struct(req); errs != nil {
		return errors.InvalidRequest(common.FormatValidateError(errs))
	}

	instructionDataID, err := primitive.ObjectIDFromHex(*req.InstructionDataID)
	if err != nil {
		return errors.InvalidRequest(fmt.Errorf("invalid instruction data ID"))
	}
	err = d.DataAuditService.ApproveInstructionData(c.UserContext(), &instructionDataID)
	if err != nil {
		return err
	}
	return c.JSON(
		vo.Response{
			Code:    errors.CodeSuccess,
			Message: errors.MessageSuccess,
			Data:    nil,
		},
	)
}

// RejectInstructionData rejects the instruction data.
//
//	@description	Reject the instruction data.
//	@id				admin-reject-instruction-data
//	@summary		reject instruction data
//	@tags			Admin API
//	@accept			json
//	@param			admin.RejectInstructionDataRequest	body	admin.RejectInstructionDataRequest	true	"Reject instruction data request"
//	@security		Bearer
//	@success		200	{object}	vo.Response{data=nil}	"Success"
//	@failure		400	{object}	vo.Response{data=nil}	"Invalid request"
//	@failure		401	{object}	vo.Response{data=nil}	"Unauthorized"
//	@failure		403	{object}	vo.Response{data=nil}	"Forbidden"
//	@failure		500	{object}	vo.Response{data=nil}	"Internal server error"
//	@router			/admin/instruction-data/reject [post]
func (d *DataAuditApi) RejectInstructionData(c *fiber.Ctx) error {
	req := new(admin.RejectInstructionDataRequest)

	if err := c.BodyParser(req); err != nil {
		return errors.InvalidRequest(fmt.Errorf("failed to parse request"))
	}

	if errs := d.Validator.Struct(req); errs != nil {
		return errors.InvalidRequest(common.FormatValidateError(errs))
	}

	instructionDataID, err := primitive.ObjectIDFromHex(*req.InstructionDataID)
	if err != nil {
		return errors.InvalidRequest(fmt.Errorf("failed to parse request"))
	}
	err = d.DataAuditService.RejectInstructionData(c.UserContext(), &instructionDataID, req.Message)
	if err != nil {
		return err
	}
	return c.JSON(
		vo.Response{
			Code:    errors.CodeSuccess,
			Message: errors.MessageSuccess,
			Data:    nil,
		},
	)
}

// UpdateInstructionData updates the instruction data.
//
//	@description	Update the instruction data.
//	@id				admin-update-instruction-data
//	@summary		update instruction data
//	@tags			Admin API
//	@accept			json
//	@param			admin.UpdateInstructionDataRequest	body	admin.UpdateInstructionDataRequest	true	"Update instruction data request"
//	@security		Bearer
//	@success		200	{object}	vo.Response{data=nil}	"Success"
//	@failure		400	{object}	vo.Response{data=nil}	"Invalid request"
//	@failure		401	{object}	vo.Response{data=nil}	"Unauthorized"
//	@failure		403	{object}	vo.Response{data=nil}	"Forbidden"
//	@failure		500	{object}	vo.Response{data=nil}	"Internal server error"
//	@router			/admin/instruction-data/update [post]
func (d *DataAuditApi) UpdateInstructionData(c *fiber.Ctx) error {
	ctx := c.UserContext()
	req := new(admin.UpdateInstructionDataRequest)

	if err := c.BodyParser(req); err != nil {
		return errors.InvalidRequest(fmt.Errorf("failed to parse request"))
	}

	if errs := d.Validator.Struct(req); errs != nil {
		return errors.InvalidRequest(common.FormatValidateError(errs))
	}

	// transform instruction data ID from string to primitive.ObjectID
	instructionDataID, err := primitive.ObjectIDFromHex(*req.InstructionDataID)
	if err != nil {
		return errors.InvalidRequest(fmt.Errorf("invalid instruction data ID"))
	}

	// transform user ID from string to primitive.ObjectID
	var (
		userID    primitive.ObjectID
		userIDPtr *primitive.ObjectID
	)
	if req.UserID != nil {
		userID, err = primitive.ObjectIDFromHex(*req.UserID)
		if err != nil {
			return errors.InvalidRequest(fmt.Errorf("invalid user ID"))
		}
		userIDPtr = &userID
	}

	// transform operator ID from string to primitive.ObjectID for logging
	operatorIDHex, ok := ctx.Value(config.UserIDKey).(string)
	if !ok {
		return errors.NotAuthorized(fmt.Errorf("user is not authorized"))
	}
	operatorID, err := primitive.ObjectIDFromHex(operatorIDHex)
	if err != nil {
		return errors.NotAuthorized(fmt.Errorf("user is not authorized"))
	}

	// update instruction data
	err = d.DataAuditService.UpdateInstructionData(
		ctx, &instructionDataID, userIDPtr, req.Instruction, req.Input, req.Output, req.Theme, req.Source, req.Note,
	)

	var (
		ipAddr     = c.IP()
		userAgent  = c.Get(fiber.HeaderUserAgent)
		operation  = config.OperationTypeUpdate
		entityType = config.EntityTypeInstruction
	)
	if err != nil {
		var (
			description = fmt.Sprintf("Update instruction data failed: %s", err.Error())
			status      = config.OperationStatusFailure
		)
		_ = d.LogsService.CacheOperationLog(
			ctx, &operatorID, &instructionDataID, &ipAddr, &userAgent, &operation, &entityType, &description, &status,
		)
		return err
	}

	var (
		description = fmt.Sprintf("Update instruction data: %s", *req.InstructionDataID)
		status      = config.OperationStatusSuccess
	)
	_ = d.LogsService.CacheOperationLog(
		ctx, &operatorID, &instructionDataID, &ipAddr, &userAgent, &operation, &entityType, &description, &status,
	)
	return c.JSON(
		vo.Response{
			Code:    errors.CodeSuccess,
			Message: errors.MessageSuccess,
			Data:    nil,
		},
	)
}

// ExportInstructionData exports the instruction data.
//
//	@description	Export the instruction data.
//	@id				admin-export-instruction-data
//	@summary		export instruction data
//	@tags			Admin API
//	@accept			json
//	@produce		json
//	@param			admin.ExportInstructionDataRequest	query	admin.ExportInstructionDataRequest	true	"Export instruction data request"
//	@security		Bearer
//	@success		200	{object}	vo.Response{data=string}	"Success"
//	@failure		400	{object}	vo.Response{data=nil}		"Invalid request"
//	@failure		401	{object}	vo.Response{data=nil}		"Unauthorized"
//	@failure		403	{object}	vo.Response{data=nil}		"Forbidden"
//	@failure		500	{object}	vo.Response{data=nil}		"Internal server error"
//	@router			/admin/instruction-data/export [get]
func (d *DataAuditApi) ExportInstructionData(c *fiber.Ctx) error {
	req := new(admin.ExportInstructionDataRequest)

	if err := c.QueryParser(req); err != nil {
		return errors.InvalidRequest(fmt.Errorf("failed to parse request"))
	}

	if errs := d.Validator.Struct(req); errs != nil {
		return errors.InvalidRequest(common.FormatValidateError(errs))
	}

	var (
		userID                               primitive.ObjectID
		userIDPtr                            *primitive.ObjectID
		createStartTime, createEndTime       time.Time
		updateStartTime, updateEndTime       time.Time
		createStartTimePtr, createEndTimePtr *time.Time
		updateStartTimePtr, updateEndTimePtr *time.Time
		err                                  error
	)

	if req.UserID != nil {
		userID, err = primitive.ObjectIDFromHex(*req.UserID)
		if err != nil {
			return errors.InvalidRequest(fmt.Errorf("failed to parse request"))
		}
		userIDPtr = &userID
	}
	if req.CreateEndTime != nil && req.CreateStartTime != nil {
		createStartTime, err = time.Parse(time.RFC3339, *req.CreateStartTime)
		if err != nil {
			return errors.InvalidRequest(fmt.Errorf("failed to parse request"))
		}
		createEndTime, err = time.Parse(time.RFC3339, *req.CreateEndTime)
		if err != nil {
			return errors.InvalidRequest(fmt.Errorf("failed to parse request"))
		}
		createStartTimePtr = &createStartTime
		createEndTimePtr = &createEndTime
	}
	if req.UpdateEndTime != nil && req.UpdateStartTime != nil {
		updateStartTime, err = time.Parse(time.RFC3339, *req.UpdateStartTime)
		if err != nil {
			return errors.InvalidRequest(fmt.Errorf("failed to parse request"))
		}
		updateEndTime, err = time.Parse(time.RFC3339, *req.UpdateEndTime)
		if err != nil {
			return errors.InvalidRequest(fmt.Errorf("failed to parse request"))
		}
		updateStartTimePtr = &updateStartTime
		updateEndTimePtr = &updateEndTime
	}
	data, err := d.DataAuditService.ExportInstructionData(
		c.UserContext(), req.Desc,
		userIDPtr, createStartTimePtr, createEndTimePtr, updateStartTimePtr, updateEndTimePtr, req.Theme, req.Status,
	)
	if err != nil {
		return err
	}
	dataJSON, err := json.Marshal(data)
	if err != nil {
		return errors.ServiceError(fmt.Errorf("failed to marshal data"))
	}
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	c.Set(
		fiber.HeaderContentDisposition,
		fmt.Sprintf("attachment; filename=%s", fmt.Sprintf("instruction_%s.json", time.Now().Format(time.RFC3339))),
	)
	return c.Send(dataJSON)
}

// ExportInstructionDataAsAlpaca exports the instruction data as Alpaca format.
//
//	@description	Export the instruction data as Alpaca format.
//	@id				admin-export-instruction-data-as-alpaca
//	@summary		export instruction data as Alpaca
//	@tags			Admin API
//	@accept			json
//	@produce		json
//	@param			admin.ExportInstructionDataRequest	query	admin.ExportInstructionDataRequest	true	"Export instruction data request"
//	@security		Bearer
//	@success		200	{object}	vo.Response{data=string}	"Success"
//	@failure		400	{object}	vo.Response{data=nil}		"Invalid request"
//	@failure		401	{object}	vo.Response{data=nil}		"Unauthorized"
//	@failure		403	{object}	vo.Response{data=nil}		"Forbidden"
//	@failure		500	{object}	vo.Response{data=nil}		"Internal server error"
//	@router			/admin/instruction-data/export/alpaca [get]
func (d *DataAuditApi) ExportInstructionDataAsAlpaca(c *fiber.Ctx) error {
	req := new(admin.ExportInstructionDataRequest)

	if err := c.QueryParser(req); err != nil {
		return errors.InvalidRequest(fmt.Errorf("failed to parse request"))
	}

	if errs := d.Validator.Struct(req); errs != nil {
		return errors.InvalidRequest(common.FormatValidateError(errs))
	}

	var (
		userID                               primitive.ObjectID
		userIDPtr                            *primitive.ObjectID
		createStartTime, createEndTime       time.Time
		updateStartTime, updateEndTime       time.Time
		createStartTimePtr, createEndTimePtr *time.Time
		updateStartTimePtr, updateEndTimePtr *time.Time
		err                                  error
	)

	if req.UserID != nil {
		userID, err = primitive.ObjectIDFromHex(*req.UserID)
		if err != nil {
			return errors.InvalidRequest(fmt.Errorf("invalid user ID %s", *req.UserID))
		}
		userIDPtr = &userID
	}
	if req.CreateStartTime != nil && req.CreateEndTime != nil {
		createStartTime, err = time.Parse(time.RFC3339, *req.CreateStartTime)
		if err != nil {
			return errors.InvalidRequest(
				fmt.Errorf(
					"invalid create start time %s (should be in `RFC3339` format)", *req.CreateStartTime,
				),
			)
		}
		createEndTime, err = time.Parse(time.RFC3339, *req.CreateEndTime)
		if err != nil {
			return errors.InvalidRequest(
				fmt.Errorf(
					"invalid create end time %s (should be in `RFC3339` format)", *req.CreateEndTime,
				),
			)
		}
		createStartTimePtr = &createStartTime
		createEndTimePtr = &createEndTime
	}
	if req.UpdateStartTime != nil && req.UpdateEndTime != nil {
		updateStartTime, err = time.Parse(time.RFC3339, *req.UpdateStartTime)
		if err != nil {
			return errors.InvalidRequest(
				fmt.Errorf(
					"invalid update start time %s (should be in `RFC3339` format)", *req.UpdateStartTime,
				),
			)
		}
		updateEndTime, err = time.Parse(time.RFC3339, *req.UpdateEndTime)
		if err != nil {
			return errors.InvalidRequest(
				fmt.Errorf(
					"invalid update end time %s (should be in `RFC3339` format)", *req.UpdateEndTime,
				),
			)
		}
		updateStartTimePtr = &updateStartTime
		updateEndTimePtr = &updateEndTime
	}
	data, err := d.DataAuditService.ExportInstructionDataAsAlpaca(
		c.UserContext(), req.Desc,
		userIDPtr, createStartTimePtr, createEndTimePtr, updateStartTimePtr, updateEndTimePtr, req.Theme, req.Status,
	)
	if err != nil {
		return err
	}
	dataJSON, err := json.Marshal(data)
	if err != nil {
		return errors.ServiceError(fmt.Errorf("failed to marshal data"))
	}
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	c.Set(
		fiber.HeaderContentDisposition, fmt.Sprintf(
			"attachment; filename=%s",
			fmt.Sprintf("instruction_alpaca_%s.json", time.Now().Format(time.RFC3339)),
		),
	)
	return c.Send(dataJSON)
}

// DeleteInstructionData deletes the instruction data.
//
//	@description	Delete the instruction data.
//	@id				admin-delete-instruction-data
//	@summary		delete instruction data
//	@tags			Admin API
//	@accept			json
//	@param			admin.DeleteInstructionDataRequest	query	admin.DeleteInstructionDataRequest	true	"Delete instruction data request"
//	@security		Bearer
//	@success		200	{object}	vo.Response{data=nil}	"Success"
//	@failure		400	{object}	vo.Response{data=nil}	"Invalid request"
//	@failure		401	{object}	vo.Response{data=nil}	"Unauthorized"
//	@failure		403	{object}	vo.Response{data=nil}	"Forbidden"
//	@failure		500	{object}	vo.Response{data=nil}	"Internal server error"
//	@router			/admin/instruction-data/delete [delete]
func (d *DataAuditApi) DeleteInstructionData(c *fiber.Ctx) error {
	req := new(admin.DeleteInstructionDataRequest)

	if err := c.QueryParser(req); err != nil {
		return errors.InvalidRequest(fmt.Errorf("failed to parse request"))
	}

	if errs := d.Validator.Struct(req); errs != nil {
		return errors.InvalidRequest(common.FormatValidateError(errs))
	}

	instructionDataID, err := primitive.ObjectIDFromHex(*req.InstructionDataID)
	if err != nil {
		return errors.InvalidRequest(fmt.Errorf("failed to parse request"))
	}
	err = d.DataAuditService.DeleteInstructionData(c.UserContext(), &instructionDataID)
	var (
		userID, _  = primitive.ObjectIDFromHex(c.UserContext().Value(config.UserIDKey).(string))
		ipAddr     = c.IP()
		userAgent  = c.Get(fiber.HeaderUserAgent)
		operation  = config.OperationTypeDelete
		entityType = config.EntityTypeInstruction
	)

	if err != nil {
		var (
			description = fmt.Sprintf("Delete instruction data failed: %s", err.Error())
			status      = config.OperationStatusFailure
		)
		_ = d.LogsService.CacheOperationLog(
			c.UserContext(), &userID, &instructionDataID, &ipAddr, &userAgent, &operation, &entityType, &description,
			&status,
		)
		return err
	}

	var (
		description = fmt.Sprintf("Delete instruction data: %s", *req.InstructionDataID)
		status      = config.OperationStatusSuccess
	)
	_ = d.LogsService.CacheOperationLog(
		c.UserContext(), &userID, &instructionDataID, &ipAddr, &userAgent, &operation, &entityType, &description,
		&status,
	)
	return c.JSON(
		vo.Response{
			Code:    errors.CodeSuccess,
			Message: errors.MessageSuccess,
			Data:    nil,
		},
	)
}
