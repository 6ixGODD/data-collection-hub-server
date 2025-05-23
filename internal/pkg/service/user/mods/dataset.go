package mods

import (
	"context"
	e "errors"
	"fmt"
	"time"

	"data-collection-hub-server/internal/pkg/config"
	dao "data-collection-hub-server/internal/pkg/dao/mods"
	"data-collection-hub-server/internal/pkg/domain/vo/user"
	"data-collection-hub-server/internal/pkg/service"
	"data-collection-hub-server/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
	ctx context.Context, instruction, input, output, theme, source, note *string,
) (string, error) {
	userIDHex, ok := ctx.Value(config.UserIDKey).(string)
	if !ok {
		return "", errors.NotAuthorized(fmt.Errorf("user is not authorized"))
	}
	userID, err := primitive.ObjectIDFromHex(userIDHex)
	if err != nil {
		return "", errors.NotAuthorized(fmt.Errorf("user is not authorized"))
	}

	var t, n string
	if theme == nil {
		t = "Default"
	} else {
		t = *theme
	}
	if note == nil {
		n = ""
	} else {
		n = *note
	}
	instructionDataID, err := d.instructionDataDao.InsertInstructionData(
		ctx, userID, *instruction, *input, *output, t, *source, n, config.InstructionDataStatusPending, "",
	)
	if err != nil {
		return "", errors.OperationFailed(fmt.Errorf("failed to insert instruction data"))
	}
	return instructionDataID.Hex(), nil
}

func (d datasetServiceImpl) GetInstructionData(
	ctx context.Context, instructionDataID primitive.ObjectID,
) (*user.GetInstructionDataResponse, error) {
	instructionData, err := d.instructionDataDao.GetInstructionDataByID(ctx, instructionDataID)
	if err != nil {
		if e.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.NotFound(fmt.Errorf("instruction (id: %s) data not found", instructionDataID.Hex()))
		} else {
			return nil, errors.OperationFailed(
				fmt.Errorf("failed to get instruction data (id: %s)", instructionDataID.Hex()),
			)
		}
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
	userIDHex, ok := ctx.Value(config.UserIDKey).(string)
	if !ok {
		return nil, errors.NotAuthorized(fmt.Errorf("user is not authorized"))
	}
	userID, err := primitive.ObjectIDFromHex(userIDHex)
	if err != nil {
		return nil, errors.NotAuthorized(fmt.Errorf("user is not authorized"))
	}
	instructionDataList, count, err := d.instructionDataDao.GetInstructionDataList(
		ctx, offset, *pageSize, false, &userID, theme, status,
		nil, nil, updateBefore, updateAfter, nil,
	)
	if err != nil {
		return nil, errors.OperationFailed(fmt.Errorf("failed to get instruction data list"))
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
	ctx context.Context, instructionDataID *primitive.ObjectID, instruction, input, output, theme, source, note *string,
) error {
	// Check if the instruction data exists and is in pending status (only pending status can be updated by the user)
	instructionData, err := d.instructionDataDao.GetInstructionDataByID(ctx, *instructionDataID)
	if err != nil {
		if e.Is(err, mongo.ErrNoDocuments) {
			return errors.NotFound(fmt.Errorf("instruction data (id: %s) not found", instructionDataID.Hex()))
		} else {
			return errors.OperationFailed(
				fmt.Errorf("failed to get instruction data (id: %s)", instructionDataID.Hex()),
			)
		}
	}
	if instructionData.Status.Code != config.InstructionDataStatusPending {
		return errors.PermissionDeny(
			fmt.Errorf("instruction data (id: %s) is not in pending status", instructionDataID.Hex()),
		)
	}

	err = d.instructionDataDao.UpdateInstructionData(
		ctx, *instructionDataID, nil, instruction, input, output, theme, source, note, nil, nil,
	)
	if err != nil {
		if e.Is(err, mongo.ErrNoDocuments) {
			return errors.NotFound(fmt.Errorf("instruction (id: %s) data not found", instructionDataID.Hex()))
		} else {
			return errors.OperationFailed(
				fmt.Errorf(
					"failed to update instruction data (id: %s)", instructionDataID.Hex(),
				),
			)
		}
	}
	return nil
}

func (d datasetServiceImpl) DeleteInstructionData(ctx context.Context, instructionDataID *primitive.ObjectID) error {
	// Check if the instruction data exists and is in pending status (only pending status can be deleted by the user)
	instructionData, err := d.instructionDataDao.GetInstructionDataByID(ctx, *instructionDataID)
	if err != nil {
		if e.Is(err, mongo.ErrNoDocuments) {
			return errors.NotFound(fmt.Errorf("instruction data (id: %s) not found", instructionDataID.Hex()))
		} else {
			return errors.OperationFailed(
				fmt.Errorf("failed to get instruction data (id: %s)", instructionDataID.Hex()),
			)
		}
	}
	if instructionData.Status.Code != config.InstructionDataStatusPending {
		return errors.PermissionDeny(
			fmt.Errorf("instruction data (id: %s) is not in pending status", instructionDataID.Hex()),
		)
	}

	err = d.instructionDataDao.SoftDeleteInstructionData(ctx, *instructionDataID)
	if err != nil {
		return errors.OperationFailed(fmt.Errorf("failed to delete instruction data (id: %s)", instructionDataID.Hex()))
	}
	return nil
}
