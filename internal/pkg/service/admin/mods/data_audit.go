package mods

import (
	"context"
	e "errors"
	"fmt"
	"time"

	"data-collection-hub-server/internal/pkg/config"
	dao "data-collection-hub-server/internal/pkg/dao/mods"
	"data-collection-hub-server/internal/pkg/domain/vo/admin"
	"data-collection-hub-server/internal/pkg/service"
	"data-collection-hub-server/pkg/errors"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DataAuditService interface {
	GetInstructionData(
		ctx context.Context, instructionDataID primitive.ObjectID,
	) (*admin.GetInstructionDataResponse, error)
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
	ExportInstructionData(
		ctx context.Context, desc *bool, userID *primitive.ObjectID,
		createTimeStart, createTimeEnd, updateTimeStart, updateTimeEnd *time.Time,
		theme, status *string,
	) (*admin.InstructionDataList, error)
	ExportInstructionDataAsAlpaca(
		ctx context.Context, desc *bool, userID *primitive.ObjectID,
		createTimeStart, createTimeEnd, updateTimeStart, updateTimeEnd *time.Time,
		theme, status *string,
	) (*admin.InstructionDataAlpacaList, error)
	DeleteInstructionData(ctx context.Context, instructionDataID *primitive.ObjectID) error
}

type DataAuditServiceImpl struct {
	core               *service.Core
	instructionDataDao dao.InstructionDataDao
}

func NewDataAuditService(core *service.Core, instructionDataDao dao.InstructionDataDao) DataAuditService {
	return &DataAuditServiceImpl{
		core:               core,
		instructionDataDao: instructionDataDao,
	}
}

func (d DataAuditServiceImpl) GetInstructionData(
	ctx context.Context, instructionDataID primitive.ObjectID,
) (*admin.GetInstructionDataResponse, error) {
	instructionData, err := d.instructionDataDao.GetInstructionDataByID(ctx, instructionDataID)
	if err != nil {
		if e.Is(err, qmgo.ErrNoSuchDocuments) {
			return nil, errors.NotFound(fmt.Errorf("instruction data (id: %s) not found", instructionDataID.Hex()))
		} else {
			return nil, errors.OperationFailed(
				fmt.Errorf(
					"failed to get instruction data (id: %s)", instructionDataID.Hex(),
				),
			)
		}
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
		return nil, errors.OperationFailed(fmt.Errorf("failed to get instruction data list")) // TODO: Should contain more situation e.g. sometime it's caused by not found or the input is illegal, not only operation failed
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
		if e.Is(err, qmgo.ErrNoSuchDocuments) {
			return errors.NotFound(fmt.Errorf("instruction data (id: %s) not found", instructionDataID.Hex()))
		} else {
			return errors.OperationFailed(
				fmt.Errorf(
					"failed to approve instruction data (id: %s)", instructionDataID.Hex(),
				),
			)
		}
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
		if e.Is(err, qmgo.ErrNoSuchDocuments) {
			return errors.NotFound(fmt.Errorf("instruction data (id: %s) not found", instructionDataID.Hex()))
		} else {
			return errors.OperationFailed(
				fmt.Errorf(
					"failed to reject instruction data (id: %s)", instructionDataID.Hex(),
				),
			)
		}
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
		if e.Is(err, mongo.ErrNoDocuments) {
			return errors.NotFound(fmt.Errorf("instruction data (id: %s) not found", instructionDataID.Hex()))
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

func (d DataAuditServiceImpl) ExportInstructionData(
	ctx context.Context, desc *bool, userID *primitive.ObjectID,
	createTimeStart, createTimeEnd, updateTimeStart, updateTimeEnd *time.Time,
	theme, status *string,
) (*admin.InstructionDataList, error) {
	var instructionDataList []*admin.InstructionData
	_instructionDataList, _, err := d.instructionDataDao.GetInstructionDataList(
		ctx, 0, 0, *desc, userID, theme, status, createTimeStart, createTimeEnd, updateTimeStart, updateTimeEnd, nil,
	)
	if err != nil {
		return nil, errors.OperationFailed(fmt.Errorf("failed to get instruction data list"))
	}

	for _, instructionData := range _instructionDataList {
		instructionDataList = append(
			instructionDataList, &admin.InstructionData{
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
	return &admin.InstructionDataList{
		InstructionDataList: instructionDataList,
	}, nil
}

func (d DataAuditServiceImpl) ExportInstructionDataAsAlpaca(
	ctx context.Context, desc *bool, userID *primitive.ObjectID,
	createTimeStart, createTimeEnd, updateTimeStart, updateTimeEnd *time.Time,
	theme, status *string,
) (*admin.InstructionDataAlpacaList, error) {
	var instructionDataList []*admin.InstructionDataAlpaca
	_instructionDataList, _, err := d.instructionDataDao.GetInstructionDataList(
		ctx, 0, 0, *desc, userID, theme, status, createTimeStart, createTimeEnd, updateTimeStart, updateTimeEnd, nil,
	)
	if err != nil {
		return nil, errors.OperationFailed(fmt.Errorf("failed to get instruction data list"))
	}

	for _, instructionData := range _instructionDataList {
		instructionDataList = append(
			instructionDataList, &admin.InstructionDataAlpaca{
				Institution: instructionData.Row.Instruction,
				Input:       instructionData.Row.Input,
				Output:      instructionData.Row.Output,
			},
		)
	}
	return &admin.InstructionDataAlpacaList{
		InstructionDataList: instructionDataList,
	}, nil
}

func (d DataAuditServiceImpl) DeleteInstructionData(ctx context.Context, instructionDataID *primitive.ObjectID) error {
	err := d.instructionDataDao.SoftDeleteInstructionData(ctx, *instructionDataID)
	if err != nil {
		if e.Is(err, mongo.ErrNoDocuments) {
			return errors.NotFound(fmt.Errorf("instruction data (id: %s) not found", instructionDataID.Hex()))
		}
		return errors.OperationFailed(fmt.Errorf("failed to delete instruction data (id: %s)", instructionDataID.Hex()))
	}
	return nil
}
