package modules

import (
	"context"

	"data-collection-hub-server/internal/pkg/dal"
	"data-collection-hub-server/internal/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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

type OperationLogDaoImpl struct{ *dal.Core }

func NewOperationLogDao(core *dal.Core) OperationLogDao {
	var _ OperationLogDao = (*OperationLogDaoImpl)(nil) // Ensure that the interface is implemented
	return &OperationLogDaoImpl{core}
}

func (o OperationLogDaoImpl) GetOperationLogById(operationLogId string, ctx context.Context) (*models.OperationLogModel, error) {
	var operationLog models.OperationLogModel
	collection := o.Core.Mongo.MongoDatabase.Collection(OperationLogCollectionName)
	err := collection.Find(ctx, bson.M{"operation_log_id": operationLogId}).One(&operationLog)
	if err != nil {
		o.Core.Zap.Logger.Error(
			"OperationLogDaoImpl.GetOperationLogById",
			zap.Field{Key: "operationLogId", Type: zapcore.StringType, String: operationLogId},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		o.Core.Zap.Logger.Info(
			"OperationLogDaoImpl.GetOperationLogById",
			zap.Field{Key: "operationLogId", Type: zapcore.StringType, String: operationLogId},
		)
		return &operationLog, nil
	}
}

func (o OperationLogDaoImpl) GetOperationLogList(offset, limit int64, ctx context.Context) ([]models.OperationLogModel, error) {
	var operationLogList []models.OperationLogModel
	collection := o.Core.Mongo.MongoDatabase.Collection(OperationLogCollectionName)
	err := collection.Find(ctx, bson.M{}).Skip(offset).Limit(limit).All(&operationLogList)
	if err != nil {
		o.Core.Zap.Logger.Error(
			"OperationLogDaoImpl.GetOperationLogList",
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		o.Core.Zap.Logger.Info(
			"OperationLogDaoImpl.GetOperationLogList",
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return operationLogList, nil
	}
}

func (o OperationLogDaoImpl) GetOperationLogListByUserUUID(userUUID string, offset, limit int64, ctx context.Context) ([]models.OperationLogModel, error) {
	var operationLogList []models.OperationLogModel
	collection := o.Core.Mongo.MongoDatabase.Collection(OperationLogCollectionName)
	err := collection.Find(ctx, bson.M{"user_uuid": userUUID}).Skip(offset).Limit(limit).All(&operationLogList)
	if err != nil {
		o.Core.Zap.Logger.Error(
			"OperationLogDaoImpl.GetOperationLogListByUserUUID",
			zap.Field{Key: "userUUID", Type: zapcore.StringType, String: userUUID},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		o.Core.Zap.Logger.Info(
			"OperationLogDaoImpl.GetOperationLogListByUserUUID",
			zap.Field{Key: "userUUID", Type: zapcore.StringType, String: userUUID},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return operationLogList, nil
	}
}

func (o OperationLogDaoImpl) GetOperationLogListByIPAddress(ipAddress string, offset, limit int64, ctx context.Context) ([]models.OperationLogModel, error) {
	var operationLogList []models.OperationLogModel
	collection := o.Core.Mongo.MongoDatabase.Collection(OperationLogCollectionName)
	err := collection.Find(ctx, bson.M{"ip_address": ipAddress}).Skip(offset).Limit(limit).All(&operationLogList)
	if err != nil {
		o.Core.Zap.Logger.Error(
			"OperationLogDaoImpl.GetOperationLogListByIPAddress",
			zap.Field{Key: "ipAddress", Type: zapcore.StringType, String: ipAddress},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		o.Core.Zap.Logger.Info(
			"OperationLogDaoImpl.GetOperationLogListByIPAddress",
			zap.Field{Key: "ipAddress", Type: zapcore.StringType, String: ipAddress},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return operationLogList, nil
	}
}

func (o OperationLogDaoImpl) GetOperationLogListByOperation(operation string, offset, limit int64, ctx context.Context) ([]models.OperationLogModel, error) {
	var operationLogList []models.OperationLogModel
	collection := o.Core.Mongo.MongoDatabase.Collection(OperationLogCollectionName)
	err := collection.Find(ctx, bson.M{"operation": operation}).Skip(offset).Limit(limit).All(&operationLogList)
	if err != nil {
		o.Core.Zap.Logger.Error(
			"OperationLogDaoImpl.GetOperationLogListByOperation",
			zap.Field{Key: "operation", Type: zapcore.StringType, String: operation},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		o.Core.Zap.Logger.Info(
			"OperationLogDaoImpl.GetOperationLogListByOperation",
			zap.Field{Key: "operation", Type: zapcore.StringType, String: operation},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return operationLogList, nil
	}
}

func (o OperationLogDaoImpl) GetOperationLogListByEntityUUID(entityUUID string, offset, limit int64, ctx context.Context) ([]models.OperationLogModel, error) {
	var operationLogList []models.OperationLogModel
	collection := o.Core.Mongo.MongoDatabase.Collection(OperationLogCollectionName)
	err := collection.Find(ctx, bson.M{"entity_uuid": entityUUID}).Skip(offset).Limit(limit).All(&operationLogList)
	if err != nil {
		o.Core.Zap.Logger.Error(
			"OperationLogDaoImpl.GetOperationLogListByEntityUUID",
			zap.Field{Key: "entityUUID", Type: zapcore.StringType, String: entityUUID},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		o.Core.Zap.Logger.Info(
			"OperationLogDaoImpl.GetOperationLogListByEntityUUID",
			zap.Field{Key: "entityUUID", Type: zapcore.StringType, String: entityUUID},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return operationLogList, nil
	}
}

func (o OperationLogDaoImpl) GetOperationLogListByCreatedTime(startTime, endTime string, offset, limit int64, ctx context.Context) ([]models.OperationLogModel, error) {
	var operationLogList []models.OperationLogModel
	collection := o.Core.Mongo.MongoDatabase.Collection(OperationLogCollectionName)
	err := collection.Find(ctx, bson.M{"created_time": bson.M{"$gte": startTime, "$lte": endTime}}).Skip(offset).Limit(limit).All(&operationLogList)
	if err != nil {
		o.Core.Zap.Logger.Error(
			"OperationLogDaoImpl.GetOperationLogListByCreatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		o.Core.Zap.Logger.Info(
			"OperationLogDaoImpl.GetOperationLogListByCreatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return operationLogList, nil
	}
}

func (o OperationLogDaoImpl) InsertOperationLog(operationLog *models.OperationLogModel, ctx context.Context) error {
	collection := o.Core.Mongo.MongoDatabase.Collection(OperationLogCollectionName)
	result, err := collection.InsertOne(ctx, operationLog)
	if err != nil {
		o.Core.Zap.Logger.Error(
			"OperationLogDaoImpl.InsertOperationLog",
			zap.Field{Key: "operationLog", Type: zapcore.ObjectMarshalerType, Interface: operationLog},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
	} else {
		o.Core.Zap.Logger.Info(
			"OperationLogDaoImpl.InsertOperationLog",
			zap.Field{Key: "operationLog", Type: zapcore.ObjectMarshalerType, Interface: operationLog},
			zap.Field{Key: "result", Type: zapcore.ObjectMarshalerType, Interface: result},
		)
	}
	return err
}

func (o OperationLogDaoImpl) DeleteOperationLog(operationLogId string, ctx context.Context) error {
	collection := o.Core.Mongo.MongoDatabase.Collection(OperationLogCollectionName)
	err := collection.RemoveId(ctx, operationLogId)
	if err != nil {
		o.Core.Zap.Logger.Error(
			"OperationLogDaoImpl.DeleteOperationLog",
			zap.Field{Key: "operationLogId", Type: zapcore.StringType, String: operationLogId},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
	} else {
		o.Core.Zap.Logger.Info(
			"OperationLogDaoImpl.DeleteOperationLog",
			zap.Field{Key: "operationLogId", Type: zapcore.StringType, String: operationLogId},
		)
	}
	return err
}
