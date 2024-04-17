package dao

import (
	"context"

	"data-collection-hub-server/models"
	"github.com/qiniu/qmgo"
)

type InstructionDataDao interface {
	GetInstructionDataById(instructionDataId string, mongoClient *qmgo.QmgoClient, ctx context.Context) (*models.InstructionDataModel, error)
	GetInstructionDataByRowId(rowId string, mongoClient *qmgo.QmgoClient, ctx context.Context) (*models.InstructionDataModel, error)
	GetInstructionDataList(mongoClient *qmgo.QmgoClient, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByUserUUID(mongoClient *qmgo.QmgoClient, userUUID string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByFuzzyQuery(mongoClient *qmgo.QmgoClient, query string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByFuzzyQueryAndTheme(mongoClient *qmgo.QmgoClient, query, theme string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByFuzzyQueryAndThemeAndStatusCode(mongoClient *qmgo.QmgoClient, query, theme, statusCode string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByFuzzyQueryAndThemeAndCreatedTime(mongoClient *qmgo.QmgoClient, query, theme, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByFuzzyQueryAndThemeAndStatusCodeAndCreatedTime(mongoClient *qmgo.QmgoClient, query, theme, statusCode, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByFuzzyQueryAndStatusCode(mongoClient *qmgo.QmgoClient, query, statusCode string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByFuzzyQueryAndStatusCodeAndCreatedTime(mongoClient *qmgo.QmgoClient, query, statusCode, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByFuzzyQueryAndCreatedTime(mongoClient *qmgo.QmgoClient, query, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByTheme(mongoClient *qmgo.QmgoClient, theme string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByThemeAndStatusCode(mongoClient *qmgo.QmgoClient, theme, statusCode string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByThemeAndCreatedTime(mongoClient *qmgo.QmgoClient, theme, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByThemeAndStatusCodeAndCreatedTime(mongoClient *qmgo.QmgoClient, theme, statusCode, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByStatusCode(mongoClient *qmgo.QmgoClient, statusCode string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByStatusCodeAndCreatedTime(mongoClient *qmgo.QmgoClient, statusCode, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByCreatedTime(mongoClient *qmgo.QmgoClient, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	GetInstructionDataListByUpdatedTime(mongoClient *qmgo.QmgoClient, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error)
	InsertInstructionData(instructionData *models.InstructionDataModel, mongoClient *qmgo.QmgoClient, ctx context.Context) error
	UpdateInstructionData(instructionData *models.InstructionDataModel, mongoClient *qmgo.QmgoClient, ctx context.Context) error
	DeleteInstructionData(instructionData *models.InstructionDataModel, mongoClient *qmgo.QmgoClient, ctx context.Context) error
}

type InstructionData struct{}

func (i InstructionData) GetInstructionDataById(instructionDataId string, mongoClient *qmgo.QmgoClient, ctx context.Context) (*models.InstructionDataModel, error) {
	// TODO implement me
	panic("implement me")
}

func (i InstructionData) GetInstructionDataByRowId(rowId string, mongoClient *qmgo.QmgoClient, ctx context.Context) (*models.InstructionDataModel, error) {
	// TODO implement me
	panic("implement me")
}

func (i InstructionData) GetInstructionDataList(mongoClient *qmgo.QmgoClient, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	// TODO implement me
	panic("implement me")
}

func (i InstructionData) GetInstructionDataListByUserUUID(mongoClient *qmgo.QmgoClient, userUUID string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	// TODO implement me
	panic("implement me")
}

func (i InstructionData) GetInstructionDataListByFuzzyQuery(mongoClient *qmgo.QmgoClient, query string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	// TODO implement me
	panic("implement me")
}

func (i InstructionData) GetInstructionDataListByFuzzyQueryAndTheme(mongoClient *qmgo.QmgoClient, query, theme string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	// TODO implement me
	panic("implement me")
}

func (i InstructionData) GetInstructionDataListByFuzzyQueryAndThemeAndStatusCode(mongoClient *qmgo.QmgoClient, query, theme, statusCode string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	// TODO implement me
	panic("implement me")
}

func (i InstructionData) GetInstructionDataListByFuzzyQueryAndThemeAndCreatedTime(mongoClient *qmgo.QmgoClient, query, theme, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	// TODO implement me
	panic("implement me")
}

func (i InstructionData) GetInstructionDataListByFuzzyQueryAndThemeAndStatusCodeAndCreatedTime(mongoClient *qmgo.QmgoClient, query, theme, statusCode, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	// TODO implement me
	panic("implement me")
}

func (i InstructionData) GetInstructionDataListByFuzzyQueryAndStatusCode(mongoClient *qmgo.QmgoClient, query, statusCode string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	// TODO implement me
	panic("implement me")
}

func (i InstructionData) GetInstructionDataListByFuzzyQueryAndStatusCodeAndCreatedTime(mongoClient *qmgo.QmgoClient, query, statusCode, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	// TODO implement me
	panic("implement me")
}

func (i InstructionData) GetInstructionDataListByFuzzyQueryAndCreatedTime(mongoClient *qmgo.QmgoClient, query, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	// TODO implement me
	panic("implement me")
}

func (i InstructionData) GetInstructionDataListByTheme(mongoClient *qmgo.QmgoClient, theme string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	// TODO implement me
	panic("implement me")
}

func (i InstructionData) GetInstructionDataListByThemeAndStatusCode(mongoClient *qmgo.QmgoClient, theme, statusCode string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	// TODO implement me
	panic("implement me")
}

func (i InstructionData) GetInstructionDataListByThemeAndCreatedTime(mongoClient *qmgo.QmgoClient, theme, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	// TODO implement me
	panic("implement me")
}

func (i InstructionData) GetInstructionDataListByThemeAndStatusCodeAndCreatedTime(mongoClient *qmgo.QmgoClient, theme, statusCode, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	// TODO implement me
	panic("implement me")
}

func (i InstructionData) GetInstructionDataListByStatusCode(mongoClient *qmgo.QmgoClient, statusCode string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	// TODO implement me
	panic("implement me")
}

func (i InstructionData) GetInstructionDataListByStatusCodeAndCreatedTime(mongoClient *qmgo.QmgoClient, statusCode, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	// TODO implement me
	panic("implement me")
}

func (i InstructionData) GetInstructionDataListByCreatedTime(mongoClient *qmgo.QmgoClient, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	// TODO implement me
	panic("implement me")
}

func (i InstructionData) GetInstructionDataListByUpdatedTime(mongoClient *qmgo.QmgoClient, startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.InstructionDataModel, error) {
	// TODO implement me
	panic("implement me")
}

func (i InstructionData) InsertInstructionData(instructionData *models.InstructionDataModel, mongoClient *qmgo.QmgoClient, ctx context.Context) error {
	// TODO implement me
	panic("implement me")
}

func (i InstructionData) UpdateInstructionData(instructionData *models.InstructionDataModel, mongoClient *qmgo.QmgoClient, ctx context.Context) error {
	// TODO implement me
	panic("implement me")
}

func (i InstructionData) DeleteInstructionData(instructionData *models.InstructionDataModel, mongoClient *qmgo.QmgoClient, ctx context.Context) error {
	// TODO implement me
	panic("implement me")
}

func NewInstructionDataDao() InstructionDataDao {
	var _ InstructionDataDao = new(InstructionData)
	return &InstructionData{}
}
