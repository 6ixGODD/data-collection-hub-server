package mods

import (
	"context"
	"fmt"
	"time"

	"data-collection-hub-server/internal/pkg/config"
	dao "data-collection-hub-server/internal/pkg/dao/mods"
	"data-collection-hub-server/internal/pkg/domain/vo/user"
	"data-collection-hub-server/internal/pkg/service"
	"data-collection-hub-server/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DatasetService interface {
	InsertInstructionData(
		ctx context.Context, Instruction, Input, Output, Theme, Source, Note *string,
	) (string, error)
	GetInstructionData(ctx context.Context, instructionDataID primitive.ObjectID) (
		*user.GetInstructionDataResponse, error,
	)
	GetInstructionDataList(
		ctx context.Context, page, pageSize *int64, updateBefore, updateAfter *time.Time, theme, status *string,
	) (*user.GetInstructionDataListResponse, error)
	UpdateInstructionData(
		ctx context.Context, instructionDataID *primitive.ObjectID,
		Instruction, Input, Output, Theme, Source, Note *string,
	) error
	DeleteInstructionData(ctx context.Context, instructionDataID *primitive.ObjectID) error
}

type datasetServiceImpl struct {
	core               *service.Core
	instructionDataDao dao.InstructionDataDao
	userDao            dao.UserDao
	operationLogDao    dao.OperationLogDao
}

func NewDatasetService(
	core *service.Core, instructionDataDao dao.InstructionDataDao, operationLogDao dao.OperationLogDao,
) DatasetService {
	return &datasetServiceImpl{
		core:               core,
		instructionDataDao: instructionDataDao,
		operationLogDao:    operationLogDao,
	}
}

func (d datasetServiceImpl) InsertInstructionData(
	ctx context.Context, Instruction, Input, Output, Theme, Source, Note *string,
) (string, error) {
	var (
		userIDHex string
		ok        bool
	)
	if userIDHex, ok = ctx.Value(config.UserIDKey).(string); !ok {
		return "", errors.UserNotFound(fmt.Errorf("user id not found in context")) // TODO: change error type
	}
	userID, err := primitive.ObjectIDFromHex(userIDHex)
	if err != nil {
		return "", errors.UserNotFound(err) // TODO: change error type
	}
	instructionDataID, err := d.instructionDataDao.InsertInstructionData(
		ctx, userID, *Instruction, *Input, *Output, *Theme, *Source, *Note,
		config.InstructionDataStatusPending, "",
	)
	if err != nil {
		return "", errors.DBError(errors.WriteError(err))
	}
	return instructionDataID.Hex(), nil
}

func (d datasetServiceImpl) GetInstructionData(ctx context.Context, instructionDataID primitive.ObjectID) (
	*user.GetInstructionDataResponse, error,
) {
	instructionData, err := d.instructionDataDao.GetInstructionDataByID(ctx, instructionDataID)
	if err != nil {
		return nil, errors.DBError(errors.ReadError(err))
	}
	return &user.GetInstructionDataResponse{
		InstructionDataID: instructionData.InstructionDataID.Hex(),
		Row: struct {
			Instruction string `json:"instruction"`
			Input       string `json:"input"`
			Output      string `json:"output"`
		}{
			Instruction: instructionData.Row.Instruction,
			Input:       instructionData.Row.Input,
			Output:      instructionData.Row.Output,
		},
		Theme:  instructionData.Theme,
		Source: instructionData.Source,
		Note:   instructionData.Note,
		Status: struct {
			Code    string `json:"code"`
			Message string `json:"message"`
		}{
			Code:    instructionData.Status.Code,
			Message: instructionData.Status.Message,
		},
		CreatedAt: instructionData.CreatedAt.Format(time.RFC3339),
		UpdatedAt: instructionData.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (d datasetServiceImpl) GetInstructionDataList(
	ctx context.Context, page, pageSize *int64, updateBefore, updateAfter *time.Time, theme, status *string,
) (*user.GetInstructionDataListResponse, error) {
	offset := (*page - 1) * *pageSize
	var (
		userIDHex string
		ok        bool
	)
	if userIDHex, ok = ctx.Value(config.UserIDKey).(string); !ok {
		return nil, errors.UserNotFound(fmt.Errorf("user id not found in context")) // TODO: change error type
	}
	userID, err := primitive.ObjectIDFromHex(userIDHex)
	if err != nil {
		return nil, errors.UserNotFound(err) // TODO: change error type
	}
	instructionDataList, count, err := d.instructionDataDao.GetInstructionDataList(
		ctx, offset, *pageSize, false, &userID, theme, status,
		nil, nil, updateBefore, updateAfter, nil,
	)
	if err != nil {
		return nil, errors.DBError(errors.ReadError(err))
	}
	resp := make([]*user.GetInstructionDataResponse, 0, len(instructionDataList))
	for _, instructionData := range instructionDataList {
		resp = append(
			resp, &user.GetInstructionDataResponse{
				InstructionDataID: instructionData.InstructionDataID.Hex(),
				Row: struct {
					Instruction string `json:"instruction"`
					Input       string `json:"input"`
					Output      string `json:"output"`
				}{
					Instruction: instructionData.Row.Instruction,
					Input:       instructionData.Row.Input,
					Output:      instructionData.Row.Output,
				},
				Theme:  instructionData.Theme,
				Source: instructionData.Source,
				Note:   instructionData.Note,
				Status: struct {
					Code    string `json:"code"`
					Message string `json:"message"`
				}{
					Code:    instructionData.Status.Code,
					Message: instructionData.Status.Message,
				},
				CreatedAt: instructionData.CreatedAt.Format(time.RFC3339),
				UpdatedAt: instructionData.UpdatedAt.Format(time.RFC3339),
			},
		)
	}
	return &user.GetInstructionDataListResponse{
		Total:               *count,
		InstructionDataList: resp,
	}, nil
}

func (d datasetServiceImpl) UpdateInstructionData(
	ctx context.Context, instructionDataID *primitive.ObjectID, Instruction, Input, Output, Theme, Source, Note *string,
) error {
	err := d.instructionDataDao.UpdateInstructionData(
		ctx, *instructionDataID, nil, Instruction, Input, Output, Theme, Source, Note, nil, nil,
	)
	if err != nil {
		return errors.DBError(errors.WriteError(err))
	}
	return nil
}

func (d datasetServiceImpl) DeleteInstructionData(ctx context.Context, instructionDataID *primitive.ObjectID) error {
	instructionData, err := d.instructionDataDao.GetInstructionDataByID(ctx, *instructionDataID)
	if err != nil {
		return errors.DBError(errors.ReadError(err))
	}
	if instructionData.Status.Code != config.InstructionDataStatusPending {
		return errors.PermissionDeny(fmt.Errorf("instruction data status is not pending")) // TODO: change error type
	}
	err = d.instructionDataDao.SoftDeleteInstructionData(ctx, *instructionDataID)
	if err != nil {
		return errors.DBError(errors.WriteError(err))
	}
	return nil
}
