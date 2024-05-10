package mods

import (
	"context"
	"time"

	dao "data-collection-hub-server/internal/pkg/dal/mods"
	"data-collection-hub-server/internal/pkg/schema/user"
	"data-collection-hub-server/internal/pkg/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DatasetService interface {
	InsertInstructionData(ctx context.Context, Instruction, Input, Output, Theme, Source, Note *string) error
	GetInstructionData(ctx context.Context, instructionDataID *primitive.ObjectID) (
		*user.GetInstructionDataResponse, error,
	)
	GetInstructionDataList(
		ctx context.Context, page *int, updateBefore, updateAfter *time.Time, theme, status *string,
	) (*user.GetInstructionDataListResponse, error)
	UpdateInstructionData(
		ctx context.Context, instructionDataID *primitive.ObjectID,
		Instruction, Input, Output, Theme, Source, Note *string,
	) error
	DeleteInstructionData(ctx context.Context, instructionDataID *primitive.ObjectID) error
}

type DatasetServiceImpl struct {
	service            *service.Service
	instructionDataDao dao.InstructionDataDao
	operationLogDao    dao.OperationLogDao
}

func NewDatasetService(
	s *service.Service, instructionDataDao dao.InstructionDataDao, operationLogDao dao.OperationLogDao,
) DatasetService {
	return &DatasetServiceImpl{
		service:            s,
		instructionDataDao: instructionDataDao,
		operationLogDao:    operationLogDao,
	}
}

func (d DatasetServiceImpl) InsertInstructionData(
	ctx context.Context, Instruction, Input, Output, Theme, Source, Note *string,
) error {
	// TODO implement me
	panic("implement me")
}

func (d DatasetServiceImpl) GetInstructionData(ctx context.Context, instructionDataID *primitive.ObjectID) (
	*user.GetInstructionDataResponse, error,
) {
	// TODO implement me
	panic("implement me")
}

func (d DatasetServiceImpl) GetInstructionDataList(
	ctx context.Context, page *int, updateBefore, updateAfter *time.Time, theme, status *string,
) (*user.GetInstructionDataListResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (d DatasetServiceImpl) UpdateInstructionData(
	ctx context.Context, instructionDataID *primitive.ObjectID, Instruction, Input, Output, Theme, Source, Note *string,
) error {
	// TODO implement me
	panic("implement me")
}

func (d DatasetServiceImpl) DeleteInstructionData(ctx context.Context, instructionDataID *primitive.ObjectID) error {
	// TODO implement me
	panic("implement me")
}
