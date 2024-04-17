package dao

import (
	"context"
	"data-collection-hub-server/models"
	"github.com/qiniu/qmgo"
)

type OperationLogDao interface {
	GetOperationLogById(operationLogId string, mongoClient *qmgo.QmgoClient, ctx context.Context) (*models.OperationLogModel, error)
	GetOperationLogList(mongoClient *qmgo.QmgoClient, offset, limit int64, ctx context.Context) ([]models.OperationLogModel, error)
	GetOperationLogListByUserUUID(mongoClient *qmgo.QmgoClient, userUUID string, offset, limit int64, ctx context.Context) ([]models.OperationLogModel, error)
	GetOperationLogListByIPAddress(mongoClient *qmgo.QmgoClient, ipAddress string, offset, limit int64, ctx context.Context) ([]models.OperationLogModel, error)
	GetOperationLogListByOperation(mongoClient *qmgo.QmgoClient, operation string, offset, limit int64, ctx context.Context) ([]models.OperationLogModel, error)
	GetOperationLogListByEntityUUID(mongoClient *qmgo.QmgoClient, entityUUID string, offset, limit int64, ctx context.Context) ([]models.OperationLogModel, error)
	GetOperationLogListByCreatedTime(mongoClient *qmgo.QmgoClient, startTime, endTime string, offset, limit int64, ctx context.Context) ([]models.OperationLogModel, error)
	InsertOperationLog(operationLog *models.OperationLogModel, mongoClient *qmgo.QmgoClient, ctx context.Context) error
	DeleteOperationLog(operationLogId string, mongoClient *qmgo.QmgoClient, ctx context.Context) error
}

type OperationLogDaoImpl struct{}

func (o OperationLogDaoImpl) GetOperationLogById(operationLogId string, mongoClient *qmgo.QmgoClient, ctx context.Context) (*models.OperationLogModel, error) {
	//TODO implement me
	panic("implement me")
}

func (o OperationLogDaoImpl) GetOperationLogList(mongoClient *qmgo.QmgoClient, offset, limit int64, ctx context.Context) ([]models.OperationLogModel, error) {
	//TODO implement me
	panic("implement me")
}

func (o OperationLogDaoImpl) GetOperationLogListByUserUUID(mongoClient *qmgo.QmgoClient, userUUID string, offset, limit int64, ctx context.Context) ([]models.OperationLogModel, error) {
	//TODO implement me
	panic("implement me")
}

func (o OperationLogDaoImpl) GetOperationLogListByIPAddress(mongoClient *qmgo.QmgoClient, ipAddress string, offset, limit int64, ctx context.Context) ([]models.OperationLogModel, error) {
	//TODO implement me
	panic("implement me")
}

func (o OperationLogDaoImpl) GetOperationLogListByOperation(mongoClient *qmgo.QmgoClient, operation string, offset, limit int64, ctx context.Context) ([]models.OperationLogModel, error) {
	//TODO implement me
	panic("implement me")
}

func (o OperationLogDaoImpl) GetOperationLogListByEntityUUID(mongoClient *qmgo.QmgoClient, entityUUID string, offset, limit int64, ctx context.Context) ([]models.OperationLogModel, error) {
	//TODO implement me
	panic("implement me")
}

func (o OperationLogDaoImpl) GetOperationLogListByCreatedTime(mongoClient *qmgo.QmgoClient, startTime, endTime string, offset, limit int64, ctx context.Context) ([]models.OperationLogModel, error) {
	//TODO implement me
	panic("implement me")
}

func (o OperationLogDaoImpl) InsertOperationLog(operationLog *models.OperationLogModel, mongoClient *qmgo.QmgoClient, ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (o OperationLogDaoImpl) DeleteOperationLog(operationLogId string, mongoClient *qmgo.QmgoClient, ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func NewOperationLogDao() OperationLogDao {
	var _ OperationLogDao = new(OperationLogDaoImpl)
	return &OperationLogDaoImpl{}
}
