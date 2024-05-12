package mods

import (
	"context"
	"time"

	"data-collection-hub-server/internal/pkg/config"
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
		ctx context.Context, page, pageSize *int64, desc *bool, userID *primitive.ObjectID,
		createTimeStart, createTimeEnd, updateTimeStart, updateTimeEnd *time.Time,
		theme, status, query *string,
	) (*admin.GetInstructionDataListResponse, error)
	ApproveInstructionData(ctx context.Context, instructionDataID *primitive.ObjectID) error
	RejectInstructionData(ctx context.Context, instructionDataID *primitive.ObjectID, message *string) error
	UpdateInstructionData(
		ctx context.Context, instructionDataID, userID *primitive.ObjectID,
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
	instructionData, err := d.instructionDataDao.GetInstructionDataById(ctx, instructionDataID)
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
	ctx context.Context, page, pageSize *int64, desc *bool, userID *primitive.ObjectID,
	createTimeStart, createTimeEnd, updateTimeStart, updateTimeEnd *time.Time,
	theme, status, query *string,
) (*admin.GetInstructionDataListResponse, error) {
	offset := (*page - 1) * *pageSize
	instructionDataList, count, err := d.instructionDataDao.GetInstructionDataList(
		ctx, offset, *pageSize, *desc, userID, theme, status,
		createTimeStart, createTimeEnd, updateTimeStart, updateTimeEnd, query,
	)
	if err != nil {
		return nil, errors.ReadError(err)
	}

	resp := make([]*admin.GetInstructionDataResponse, 0, len(instructionDataList))
	for _, instructionData := range instructionDataList {
		resp = append(
			resp, &admin.GetInstructionDataResponse{
				InstructionDataID: instructionData.InstructionDataID.Hex(),
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
			},
		)
	}

	return &admin.GetInstructionDataListResponse{
		Total:               *count,
		InstructionDataList: resp,
	}, nil
}

func (d DataAuditServiceImpl) ApproveInstructionData(ctx context.Context, instructionDataID *primitive.ObjectID) error {
	status := config.InstructionDataStatusApproved
	message := ""

	err := d.instructionDataDao.UpdateInstructionData(
		ctx,
		*instructionDataID,
		nil, nil, nil, nil, nil, nil, nil,
		&status, &message,
	)
	if err != nil {
		return errors.MongoError(errors.WriteError(err))
	}
	return nil
}

func (d DataAuditServiceImpl) RejectInstructionData(
	ctx context.Context, instructionDataID *primitive.ObjectID, message *string,
) error {
	status := config.InstructionDataStatusRejected
	err := d.instructionDataDao.UpdateInstructionData(
		ctx,
		*instructionDataID,
		nil, nil, nil, nil, nil, nil, nil,
		&status, message,
	)
	if err != nil {
		return errors.MongoError(errors.WriteError(err))
	}
	return nil
}

func (d DataAuditServiceImpl) UpdateInstructionData(
	ctx context.Context, instructionDataID, userID *primitive.ObjectID,
	instruction, input, output, theme, source, note *string,
) error {
	err := d.instructionDataDao.UpdateInstructionData(
		ctx,
		*instructionDataID,
		userID, instruction, input, output, theme, source, note, nil, nil,
	)
	if err != nil {
		return errors.MongoError(errors.WriteError(err))
	}
	return nil
}

func (d DataAuditServiceImpl) DeleteInstructionData(ctx context.Context, instructionDataID *primitive.ObjectID) error {
	err := d.instructionDataDao.SoftDeleteInstructionData(ctx, *instructionDataID)
	if err != nil {
		return errors.MongoError(errors.WriteError(err))
	}
	return nil
}
