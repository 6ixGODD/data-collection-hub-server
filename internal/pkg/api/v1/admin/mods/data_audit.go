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

type DataAuditApi struct {
	adminservice.DataAuditService
}

func NewDataAuditApi(dataAuditService adminservice.DataAuditService) *DataAuditApi {
	return &DataAuditApi{dataAuditService}
}

func (d *DataAuditApi) GetInstructionData(c *fiber.Ctx) error {
	req := new(admin.GetInstructionDataRequest)

	if err := c.QueryParser(req); err != nil {
		return errors.InvalidRequest(err)
	}
	instructionDataID, err := primitive.ObjectIDFromHex(*req.InstructionDataID)
	if err != nil {
		return errors.InvalidRequest(err)
	}
	resp, err := d.DataAuditService.GetInstructionData(c.Context(), &instructionDataID)
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

func (d *DataAuditApi) GetInstructionDataList(c *fiber.Ctx) error {
	req := new(admin.GetInstructionDataListRequest)

	if err := c.QueryParser(req); err != nil {
		return errors.InvalidRequest(err)
	}

	var (
		userID                    *primitive.ObjectID
		updateBefore, updateAfter *time.Time
		err                       error
	)

	if req.UserID != nil {
		*userID, err = primitive.ObjectIDFromHex(*req.UserID)
		if err != nil {
			return errors.InvalidRequest(err)
		}
	}
	if req.UpdateBefore != nil {
		*updateBefore, err = time.Parse(time.RFC3339, *req.UpdateBefore)
		if err != nil {
			return errors.InvalidRequest(err)
		}
	}
	if req.UpdateAfter != nil {
		*updateAfter, err = time.Parse(time.RFC3339, *req.UpdateAfter)
		if err != nil {
			return errors.InvalidRequest(err)
		}
	}
	resp, err := d.DataAuditService.GetInstructionDataList(
		c.Context(), req.Page, req.Desc, userID, updateBefore, updateAfter, req.Theme, req.Status,
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

func (d *DataAuditApi) ApproveInstructionData(c *fiber.Ctx) error {
	req := new(admin.ApproveInstructionDataRequest)

	if err := c.BodyParser(req); err != nil {
		return errors.InvalidRequest(err)
	}
	instructionDataID, err := primitive.ObjectIDFromHex(*req.InstructionDataID)
	if err != nil {
		return errors.InvalidRequest(err)
	}
	err = d.DataAuditService.ApproveInstructionData(c.Context(), &instructionDataID)
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

func (d *DataAuditApi) RejectInstructionData(c *fiber.Ctx) error {
	req := new(admin.RejectInstructionDataRequest)

	if err := c.BodyParser(req); err != nil {
		return errors.InvalidRequest(err)
	}
	instructionDataID, err := primitive.ObjectIDFromHex(*req.InstructionDataID)
	if err != nil {
		return errors.InvalidRequest(err)
	}
	err = d.DataAuditService.RejectInstructionData(c.Context(), &instructionDataID, req.Message)
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

func (d *DataAuditApi) UpdateInstructionData(c *fiber.Ctx) error {
	req := new(admin.UpdateInstructionDataRequest)

	if err := c.BodyParser(req); err != nil {
		return errors.InvalidRequest(err)
	}
	instructionDataID, err := primitive.ObjectIDFromHex(*req.InstructionDataID)
	if err != nil {
		return errors.InvalidRequest(err)
	}
	err = d.DataAuditService.UpdateInstructionData(
		c.Context(), &instructionDataID, req.Instruction, req.Input, req.Output, req.Theme, req.Source, req.Note,
	)
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

func (d *DataAuditApi) DeleteInstructionData(c *fiber.Ctx) error {
	req := new(admin.DeleteInstructionDataRequest)

	if err := c.QueryParser(req); err != nil {
		return errors.InvalidRequest(err)
	}
	instructionDataID, err := primitive.ObjectIDFromHex(*req.InstructionDataID)
	if err != nil {
		return errors.InvalidRequest(err)
	}
	err = d.DataAuditService.DeleteInstructionData(c.Context(), &instructionDataID)
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
