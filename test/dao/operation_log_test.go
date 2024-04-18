package dao

import (
	"context"

	"data-collection-hub-server/models"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
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
	var operationLog models.OperationLogModel
	err := mongoClient.Find(ctx, bson.M{"_id": operationLogId}).One(&operationLog)
	if err != nil {
		return nil, err
	} else {
		return &operationLog, nil
	}
}

func (o OperationLogDaoImpl) GetOperationLogList(mongoClient *qmgo.QmgoClient, offset, limit int64, ctx context.Context) ([]models.OperationLogModel, error) {
	var operationLogList []models.OperationLogModel
	err := mongoClient.Find(ctx, bson.M{}).Skip(offset).Limit(limit).All(&operationLogList)
	if err != nil {
		return nil, err
	} else {
		return operationLogList, nil
	}
}

func (o OperationLogDaoImpl) GetOperationLogListByUserUUID(mongoClient *qmgo.QmgoClient, userUUID string, offset, limit int64, ctx context.Context) ([]models.OperationLogModel, error) {
	var operationLogList []models.OperationLogModel
	err := mongoClient.Find(ctx, bson.M{"user_uuid": userUUID}).Skip(offset).Limit(limit).All(&operationLogList)
	if err != nil {
		return nil, err
	} else {
		return operationLogList, nil
	}
}

func (o OperationLogDaoImpl) GetOperationLogListByIPAddress(mongoClient *qmgo.QmgoClient, ipAddress string, offset, limit int64, ctx context.Context) ([]models.OperationLogModel, error) {
	var operationLogList []models.OperationLogModel
	err := mongoClient.Find(ctx, bson.M{"ip_address": ipAddress}).Skip(offset).Limit(limit).All(&operationLogList)
	if err != nil {
		return nil, err
	} else {
		return operationLogList, nil
	}
}

func (o OperationLogDaoImpl) GetOperationLogListByOperation(mongoClient *qmgo.QmgoClient, operation string, offset, limit int64, ctx context.Context) ([]models.OperationLogModel, error) {
	var operationLogList []models.OperationLogModel
	err := mongoClient.Find(ctx, bson.M{"operation": operation}).Skip(offset).Limit(limit).All(&operationLogList)
	if err != nil {
		return nil, err
	} else {
		return operationLogList, nil
	}
}

func (o OperationLogDaoImpl) GetOperationLogListByEntityUUID(mongoClient *qmgo.QmgoClient, entityUUID string, offset, limit int64, ctx context.Context) ([]models.OperationLogModel, error) {
	var operationLogList []models.OperationLogModel
	err := mongoClient.Find(ctx, bson.M{"entity_uuid": entityUUID}).Skip(offset).Limit(limit).All(&operationLogList)
	if err != nil {
		return nil, err
	} else {
		return operationLogList, nil
	}
}

func (o OperationLogDaoImpl) GetOperationLogListByCreatedTime(mongoClient *qmgo.QmgoClient, startTime, endTime string, offset, limit int64, ctx context.Context) ([]models.OperationLogModel, error) {
	var operationLogList []models.OperationLogModel
	err := mongoClient.Find(
		ctx, bson.M{"created_at": bson.M{"$gte": startTime, "$lte": endTime}},
	).Skip(offset).Limit(limit).All(&operationLogList)
	if err != nil {
		return nil, err
	} else {
		return operationLogList, nil
	}
}

func (o OperationLogDaoImpl) InsertOperationLog(operationLog *models.OperationLogModel, mongoClient *qmgo.QmgoClient, ctx context.Context) error {
	_, err := mongoClient.InsertOne(ctx, operationLog)
	return err
}

func (o OperationLogDaoImpl) DeleteOperationLog(operationLogId string, mongoClient *qmgo.QmgoClient, ctx context.Context) error {
	err := mongoClient.RemoveId(ctx, operationLogId)
	return err
}

func NewOperationLogDao() OperationLogDao {
	var _ OperationLogDao = new(OperationLogDaoImpl)
	return &OperationLogDaoImpl{}
}
