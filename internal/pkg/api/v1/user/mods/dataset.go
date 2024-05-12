package mods

import (
	"time"

	"data-collection-hub-server/internal/pkg/schema"
	"data-collection-hub-server/internal/pkg/schema/user"
	userservice "data-collection-hub-server/internal/pkg/service/user/mods"
	"data-collection-hub-server/pkg/errors"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DatasetApi struct {
	userservice.DatasetService
}

func NewDatasetApi(datasetService userservice.DatasetService) *DatasetApi {
	return &DatasetApi{datasetService}
}

func (d *DatasetApi) InsertInstructionData(c *fiber.Ctx) error {
	req := new(user.InsertInstructionDataRequest)

	if err := c.BodyParser(req); err != nil {
		return errors.InvalidRequest(err)
	}

	err := d.DatasetService.InsertInstructionData(
		c.Context(), req.Instruction, req.Input, req.Output, req.Theme, req.Source, req.Note,
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

func (d *DatasetApi) GetInstructionData(c *fiber.Ctx) error {
	req := new(user.GetInstructionDataRequest)

	if err := c.QueryParser(req); err != nil {
		return errors.InvalidRequest(err)
	}

	instructionDataID, err := primitive.ObjectIDFromHex(*req.InstructionDataID)
	if err != nil {
		return errors.InvalidRequest(err)
	}

	resp, err := d.DatasetService.GetInstructionData(c.Context(), &instructionDataID)
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
		c.Context(), req.Page, req.PageSize, updateBefore, updateAfter, req.Theme, req.Status,
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

func (d *DatasetApi) UpdateInstructionData(c *fiber.Ctx) error {
	req := new(user.UpdateInstructionDataRequest)

	if err := c.BodyParser(req); err != nil {
		return errors.InvalidRequest(err)
	}

	instructionDataID, err := primitive.ObjectIDFromHex(*req.InstructionDataID)
	if err != nil {
		return errors.InvalidRequest(err)
	}

	err = d.DatasetService.UpdateInstructionData(
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

func (d *DatasetApi) DeleteInstructionData(c *fiber.Ctx) error {
	req := new(user.DeleteInstructionDataRequest)

	if err := c.QueryParser(req); err != nil {
		return errors.InvalidRequest(err)
	}

	instructionDataID, err := primitive.ObjectIDFromHex(*req.InstructionDataID)
	if err != nil {
		return errors.InvalidRequest(err)
	}

	err = d.DatasetService.DeleteInstructionData(c.Context(), &instructionDataID)
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
