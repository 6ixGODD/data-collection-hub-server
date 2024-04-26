package modules

import (
	"context"

	"data-collection-hub-server/dal"
	"data-collection-hub-server/models"
	"go.mongodb.org/mongo-driver/bson"
)

var OperationLogCollectionName = "operation_log"

type OperationLogDao interface {
	GetOperationLogById(operationLogId string, ctx context.Context) (*models.OperationLogModel, error)
	GetOperationLogList(offset, limit int64, ctx context.Context) ([]models.OperationLogModel, error)
	GetOperationLogListByUserUUID(userUUID string, offset, limit int64, ctx context.Context) ([]models.OperationLogModel, error)
	GetOperationLogListByIPAddress(ipAddress string, offset, limit int64, ctx context.Context) ([]models.OperationLogModel, error)
	GetOperationLogListByOperation(operation string, offset, limit int64, ctx context.Context) ([]models.OperationLogModel, error)
	GetOperationLogListByEntityUUID(entityUUID string, offset, limit int64, ctx context.Context) ([]models.OperationLogModel, error)
	GetOperationLogListByCreatedTime(startTime, endTime string, offset, limit int64, ctx context.Context) ([]models.OperationLogModel, error)
	InsertOperationLog(operationLog *models.OperationLogModel, ctx context.Context) error
	DeleteOperationLog(operationLogId string, ctx context.Context) error
}

type OperationLogDaoImpl struct{ *dal.Dao }

func NewOperationLogDao(dao *dal.Dao) OperationLogDao {
	var _ OperationLogDao = new(OperationLogDaoImpl)
	return &OperationLogDaoImpl{dao}
}

func (o OperationLogDaoImpl) GetOperationLogById(operationLogId string, ctx context.Context) (*models.OperationLogModel, error) {
	var operationLog models.OperationLogModel
	collection := o.Dao.MongoDB.Collection(OperationLogCollectionName)
	err := collection.Find(ctx, bson.M{"operation_log_id": operationLogId}).One(&operationLog)
	if err != nil {
		return nil, err
	} else {
		return &operationLog, nil
	}
}

func (o OperationLogDaoImpl) GetOperationLogList(offset, limit int64, ctx context.Context) ([]models.OperationLogModel, error) {
	var operationLogList []models.OperationLogModel
	collection := o.Dao.MongoDB.Collection(OperationLogCollectionName)
	err := collection.Find(ctx, bson.M{}).Skip(offset).Limit(limit).All(&operationLogList)
	if err != nil {
		return nil, err
	} else {
		return operationLogList, nil
	}
}

func (o OperationLogDaoImpl) GetOperationLogListByUserUUID(userUUID string, offset, limit int64, ctx context.Context) ([]models.OperationLogModel, error) {
	var operationLogList []models.OperationLogModel
	collection := o.Dao.MongoDB.Collection(OperationLogCollectionName)
	err := collection.Find(ctx, bson.M{"user_uuid": userUUID}).Skip(offset).Limit(limit).All(&operationLogList)
	if err != nil {
		return nil, err
	} else {
		return operationLogList, nil
	}
}

func (o OperationLogDaoImpl) GetOperationLogListByIPAddress(ipAddress string, offset, limit int64, ctx context.Context) ([]models.OperationLogModel, error) {
	var operationLogList []models.OperationLogModel
	collection := o.Dao.MongoDB.Collection(OperationLogCollectionName)
	err := collection.Find(ctx, bson.M{"ip_address": ipAddress}).Skip(offset).Limit(limit).All(&operationLogList)
	if err != nil {
		return nil, err
	} else {
		return operationLogList, nil
	}
}

func (o OperationLogDaoImpl) GetOperationLogListByOperation(operation string, offset, limit int64, ctx context.Context) ([]models.OperationLogModel, error) {
	var operationLogList []models.OperationLogModel
	collection := o.Dao.MongoDB.Collection(OperationLogCollectionName)
	err := collection.Find(ctx, bson.M{"operation": operation}).Skip(offset).Limit(limit).All(&operationLogList)
	if err != nil {
		return nil, err
	} else {
		return operationLogList, nil
	}
}

func (o OperationLogDaoImpl) GetOperationLogListByEntityUUID(entityUUID string, offset, limit int64, ctx context.Context) ([]models.OperationLogModel, error) {
	var operationLogList []models.OperationLogModel
	collection := o.Dao.MongoDB.Collection(OperationLogCollectionName)
	err := collection.Find(ctx, bson.M{"entity_uuid": entityUUID}).Skip(offset).Limit(limit).All(&operationLogList)
	if err != nil {
		return nil, err
	} else {
		return operationLogList, nil
	}
}

func (o OperationLogDaoImpl) GetOperationLogListByCreatedTime(startTime, endTime string, offset, limit int64, ctx context.Context) ([]models.OperationLogModel, error) {
	var operationLogList []models.OperationLogModel
	collection := o.Dao.MongoDB.Collection(OperationLogCollectionName)
	err := collection.Find(ctx, bson.M{"created_time": bson.M{"$gte": startTime, "$lte": endTime}}).Skip(offset).Limit(limit).All(&operationLogList)
	if err != nil {
		return nil, err
	} else {
		return operationLogList, nil
	}
}

func (o OperationLogDaoImpl) InsertOperationLog(operationLog *models.OperationLogModel, ctx context.Context) error {
	collection := o.Dao.MongoDB.Collection(OperationLogCollectionName)
	_, err := collection.InsertOne(ctx, operationLog)
	return err
}

func (o OperationLogDaoImpl) DeleteOperationLog(operationLogId string, ctx context.Context) error {
	collection := o.Dao.MongoDB.Collection(OperationLogCollectionName)
	err := collection.RemoveId(ctx, operationLogId)
	return err
}
