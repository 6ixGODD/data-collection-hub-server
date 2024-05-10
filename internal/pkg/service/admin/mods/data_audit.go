package mods

import (
	"context"
	"time"

	dao "data-collection-hub-server/internal/pkg/dal/mods"
	"data-collection-hub-server/internal/pkg/schema/admin"
	"data-collection-hub-server/internal/pkg/service"
	"data-collection-hub-server/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DataAuditService interface {
	GetInstructionData(ctx context.Context, instructionDataID *primitive.ObjectID) (
		*admin.GetInstructionDataResponse, error,
	)
	GetInstructionDataList(
		ctx context.Context, page *int, desc *bool, userID *primitive.ObjectID, updateBefore, updateAfter *time.Time,
		theme, status *string,
	) (*admin.GetInstructionDataListResponse, error)
	ApproveInstructionData(ctx context.Context, instructionDataID *primitive.ObjectID) error
	RejectInstructionData(ctx context.Context, instructionDataID *primitive.ObjectID, message *string) error
	UpdateInstructionData(
		ctx context.Context, instructionDataID *primitive.ObjectID,
		instruction, input, output, theme, source, note *string,
	) error
	DeleteInstructionData(ctx context.Context, instructionDataID *primitive.ObjectID) error
}

type DataAuditServiceImpl struct {
	service            *service.Service
	instructionDataDao dao.InstructionDataDao
}

func NewDataAuditService(s *service.Service, instructionDataDao dao.InstructionDataDao) DataAuditService {
	return &DataAuditServiceImpl{
		service:            s,
		instructionDataDao: instructionDataDao,
	}
}

func (d DataAuditServiceImpl) GetInstructionData(
	ctx context.Context, instructionDataID *primitive.ObjectID,
) (*admin.GetInstructionDataResponse, error) {
	instructionData, err := d.instructionDataDao.GetInstructionDataById(*instructionDataID, ctx)
	if err != nil {
		return nil, errors.ReadError(err)
	}
	return &admin.GetInstructionDataResponse{
		InstructionDataID: instructionDataID.Hex(),
		UserID:            instructionData.UserID.Hex(),
		Username:          instructionData.Username,
		Row: struct {
			Instruction string `json:"instruction"`
			Input       string `json:"input"`
			Output      string `json:"output"`
		}(struct {
			Instruction string
			Input       string
			Output      string
		}{
			Instruction: instructionData.Row.Instruction,
			Input:       instructionData.Row.Input,
			Output:      instructionData.Row.Output,
		}),
		Theme:  instructionData.Theme,
		Source: instructionData.Source,
		Note:   instructionData.Note,
		Status: struct {
			Code    string `json:"code"`
			Message string `json:"message"`
		}(struct {
			Code    string
			Message string
		}{
			Code:    instructionData.Status.Code,
			Message: instructionData.Status.Message,
		}),
		CreatedAt: instructionData.CreatedAt.Format(time.RFC3339),
		UpdatedAt: instructionData.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (d DataAuditServiceImpl) GetInstructionDataList(
	ctx context.Context, page *int, desc *bool, userID *primitive.ObjectID, updateBefore, updateAfter *time.Time,
	theme, status *string,
) (*admin.GetInstructionDataListResponse, error) {
	offset := d.service.Config.BaseConfig.PageLimit * (*page - 1)

}

func (d DataAuditServiceImpl) ApproveInstructionData(ctx context.Context, instructionDataID *primitive.ObjectID) error {
	// TODO implement me
	panic("implement me")
}

func (d DataAuditServiceImpl) RejectInstructionData(
	ctx context.Context, instructionDataID *primitive.ObjectID, message *string,
) error {
	// TODO implement me
	panic("implement me")
}

func (d DataAuditServiceImpl) UpdateInstructionData(
	ctx context.Context, instructionDataID *primitive.ObjectID, instruction, input, output, theme, source, note *string,
) error {
	// TODO implement me
	panic("implement me")
}

func (d DataAuditServiceImpl) DeleteInstructionData(ctx context.Context, instructionDataID *primitive.ObjectID) error {
	// TODO implement me
	panic("implement me")
}
