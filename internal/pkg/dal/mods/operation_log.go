package mods

import (
	"context"
	"time"

	"data-collection-hub-server/internal/pkg/dal"
	"data-collection-hub-server/internal/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const operationLogCollectionName = "operation_log"

type OperationLogDao interface {
	GetOperationLogById(operationLogID primitive.ObjectID, ctx context.Context) (*models.OperationLogModel, error)
	GetOperationLogList(offset, limit int64, ctx context.Context) ([]models.OperationLogModel, error)
	GetOperationLogListByUserID(
		userID primitive.ObjectID, offset, limit int64, ctx context.Context,
	) ([]models.OperationLogModel, error)
	GetOperationLogListByIPAddress(
		ipAddress string, offset, limit int64, ctx context.Context,
	) ([]models.OperationLogModel, error)
	GetOperationLogListByOperation(
		operation string, offset, limit int64, ctx context.Context,
	) ([]models.OperationLogModel, error)
	GetOperationLogListByEntityID(
		entityID primitive.ObjectID, offset, limit int64, ctx context.Context,
	) ([]models.OperationLogModel, error)
	GetOperationLogListByEntityType(
		entityType string, offset, limit int64, ctx context.Context,
	) ([]models.OperationLogModel, error)
	GetOperationLogListByStatus(status string, offset, limit int64, ctx context.Context) (
		[]models.OperationLogModel, error,
	)
	GetOperationLogListByCreatedTime(
		startTime, endTime time.Time, offset, limit int64, ctx context.Context,
	) ([]models.OperationLogModel, error)
	InsertOperationLog(
		userID, entityID primitive.ObjectID,
		username, email, ipAddress, userAgent, operation, entityType, description, status string, ctx context.Context,
	) (primitive.ObjectID, error)
	DeleteOperationLog(operationLogID primitive.ObjectID, ctx context.Context) error
}

type OperationLogDaoImpl struct{ *dal.Dao }

func NewOperationLogDao(dao *dal.Dao) OperationLogDao {
	var _ OperationLogDao = (*OperationLogDaoImpl)(nil) // Ensure that the interface is implemented
	return &OperationLogDaoImpl{dao}
}

func (o OperationLogDaoImpl) GetOperationLogById(
	operationLogID primitive.ObjectID, ctx context.Context,
) (*models.OperationLogModel, error) {
	var operationLog models.OperationLogModel
	collection := o.Dao.Mongo.MongoDatabase.Collection(operationLogCollectionName)
	err := collection.Find(ctx, bson.M{"operation_log_id": operationLogID}).One(&operationLog)
	if err != nil {
		o.Dao.Zap.Logger.Error(
			"OperationLogDaoImpl.GetOperationLogById",
			zap.Field{Key: "operationLogID", Type: zapcore.StringType, String: operationLogID.Hex()},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		o.Dao.Zap.Logger.Info(
			"OperationLogDaoImpl.GetOperationLogById",
			zap.Field{Key: "operationLogID", Type: zapcore.StringType, String: operationLogID.Hex()},
		)
		return &operationLog, nil
	}
}

func (o OperationLogDaoImpl) GetOperationLogList(offset, limit int64, ctx context.Context) (
	[]models.OperationLogModel, error,
) {
	var operationLogList []models.OperationLogModel
	collection := o.Dao.Mongo.MongoDatabase.Collection(operationLogCollectionName)
	err := collection.Find(ctx, bson.M{}).Skip(offset).Limit(limit).All(&operationLogList)
	if err != nil {
		o.Dao.Zap.Logger.Error(
			"OperationLogDaoImpl.GetOperationLogList",
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		o.Dao.Zap.Logger.Info(
			"OperationLogDaoImpl.GetOperationLogList",
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return operationLogList, nil
	}
}

func (o OperationLogDaoImpl) GetOperationLogListByUserID(
	userID primitive.ObjectID, offset, limit int64, ctx context.Context,
) ([]models.OperationLogModel, error) {
	var operationLogList []models.OperationLogModel
	collection := o.Dao.Mongo.MongoDatabase.Collection(operationLogCollectionName)
	err := collection.Find(ctx, bson.M{"user_id": userID}).Skip(offset).Limit(limit).All(&operationLogList)
	if err != nil {
		o.Dao.Zap.Logger.Error(
			"OperationLogDaoImpl.GetOperationLogListByUserID",
			zap.Field{Key: "userID", Type: zapcore.StringType, String: userID.Hex()},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		o.Dao.Zap.Logger.Info(
			"OperationLogDaoImpl.GetOperationLogListByUserID",
			zap.Field{Key: "userID", Type: zapcore.StringType, String: userID.Hex()},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return operationLogList, nil
	}
}

func (o OperationLogDaoImpl) GetOperationLogListByIPAddress(
	ipAddress string, offset, limit int64, ctx context.Context,
) ([]models.OperationLogModel, error) {
	var operationLogList []models.OperationLogModel
	collection := o.Dao.Mongo.MongoDatabase.Collection(operationLogCollectionName)
	err := collection.Find(ctx, bson.M{"ip_address": ipAddress}).Skip(offset).Limit(limit).All(&operationLogList)
	if err != nil {
		o.Dao.Zap.Logger.Error(
			"OperationLogDaoImpl.GetOperationLogListByIPAddress",
			zap.Field{Key: "ipAddress", Type: zapcore.StringType, String: ipAddress},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		o.Dao.Zap.Logger.Info(
			"OperationLogDaoImpl.GetOperationLogListByIPAddress",
			zap.Field{Key: "ipAddress", Type: zapcore.StringType, String: ipAddress},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return operationLogList, nil
	}
}

func (o OperationLogDaoImpl) GetOperationLogListByOperation(
	operation string, offset, limit int64, ctx context.Context,
) ([]models.OperationLogModel, error) {
	var operationLogList []models.OperationLogModel
	collection := o.Dao.Mongo.MongoDatabase.Collection(operationLogCollectionName)
	err := collection.Find(ctx, bson.M{"operation": operation}).Skip(offset).Limit(limit).All(&operationLogList)
	if err != nil {
		o.Dao.Zap.Logger.Error(
			"OperationLogDaoImpl.GetOperationLogListByOperation",
			zap.Field{Key: "operation", Type: zapcore.StringType, String: operation},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		o.Dao.Zap.Logger.Info(
			"OperationLogDaoImpl.GetOperationLogListByOperation",
			zap.Field{Key: "operation", Type: zapcore.StringType, String: operation},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return operationLogList, nil
	}
}

func (o OperationLogDaoImpl) GetOperationLogListByEntityID(
	entityID primitive.ObjectID, offset, limit int64, ctx context.Context,
) ([]models.OperationLogModel, error) {
	var operationLogList []models.OperationLogModel
	collection := o.Dao.Mongo.MongoDatabase.Collection(operationLogCollectionName)
	err := collection.Find(ctx, bson.M{"entity_id": entityID}).Skip(offset).Limit(limit).All(&operationLogList)
	if err != nil {
		o.Dao.Zap.Logger.Error(
			"OperationLogDaoImpl.GetOperationLogListByEntityID",
			zap.Field{Key: "entityID", Type: zapcore.StringType, String: entityID.Hex()},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		o.Dao.Zap.Logger.Info(
			"OperationLogDaoImpl.GetOperationLogListByEntityID",
			zap.Field{Key: "entityID", Type: zapcore.StringType, String: entityID.Hex()},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return operationLogList, nil
	}
}

func (o OperationLogDaoImpl) GetOperationLogListByEntityType(
	entityType string, offset, limit int64, ctx context.Context,
) ([]models.OperationLogModel, error) {
	var operationLogList []models.OperationLogModel
	collection := o.Dao.Mongo.MongoDatabase.Collection(operationLogCollectionName)
	err := collection.Find(ctx, bson.M{"entity_type": entityType}).Skip(offset).Limit(limit).All(&operationLogList)
	if err != nil {
		o.Dao.Zap.Logger.Error(
			"OperationLogDaoImpl.GetOperationLogListByEntityType",
			zap.Field{Key: "entityType", Type: zapcore.StringType, String: entityType},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		o.Dao.Zap.Logger.Info(
			"OperationLogDaoImpl.GetOperationLogListByEntityType",
			zap.Field{Key: "entityType", Type: zapcore.StringType, String: entityType},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return operationLogList, nil
	}
}

func (o OperationLogDaoImpl) GetOperationLogListByStatus(
	status string, offset, limit int64, ctx context.Context,
) ([]models.OperationLogModel, error) {
	var operationLogList []models.OperationLogModel
	collection := o.Dao.Mongo.MongoDatabase.Collection(operationLogCollectionName)
	err := collection.Find(ctx, bson.M{"status": status}).Skip(offset).Limit(limit).All(&operationLogList)
	if err != nil {
		o.Dao.Zap.Logger.Error(
			"OperationLogDaoImpl.GetOperationLogListByStatus",
			zap.Field{Key: "status", Type: zapcore.StringType, String: status},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		o.Dao.Zap.Logger.Info(
			"OperationLogDaoImpl.GetOperationLogListByStatus",
			zap.Field{Key: "status", Type: zapcore.StringType, String: status},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return operationLogList, nil
	}
}

func (o OperationLogDaoImpl) GetOperationLogListByCreatedTime(
	startTime, endTime time.Time, offset, limit int64, ctx context.Context,
) ([]models.OperationLogModel, error) {
	var operationLogList []models.OperationLogModel
	collection := o.Dao.Mongo.MongoDatabase.Collection(operationLogCollectionName)
	err := collection.Find(
		ctx, bson.M{"created_time": bson.M{"$gte": startTime, "$lte": endTime}},
	).Skip(offset).Limit(limit).All(&operationLogList)
	if err != nil {
		o.Dao.Zap.Logger.Error(
			"OperationLogDaoImpl.GetOperationLogListByCreatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime.Format(time.RFC3339)},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime.Format(time.RFC3339)},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		o.Dao.Zap.Logger.Info(
			"OperationLogDaoImpl.GetOperationLogListByCreatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime.Format(time.RFC3339)},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime.Format(time.RFC3339)},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return operationLogList, nil
	}
}

func (o OperationLogDaoImpl) InsertOperationLog(
	userID, entityID primitive.ObjectID,
	username, email, ipAddress, userAgent, operation, entityType, description, status string,
	ctx context.Context,
) (primitive.ObjectID, error) {
	collection := o.Dao.Mongo.MongoDatabase.Collection(operationLogCollectionName)
	operationLog := models.OperationLogModel{
		OperationLogID: primitive.NewObjectID(),
		UserID:         userID,
		EntityID:       entityID,
		Username:       username,
		Email:          email,
		IPAddress:      ipAddress,
		UserAgent:      userAgent,
		Operation:      operation,
		EntityType:     entityType,
		Description:    description,
		Status:         status,
		CreatedAt:      time.Now(),
	}
	result, err := collection.InsertOne(ctx, operationLog)
	if err != nil {
		o.Dao.Zap.Logger.Error(
			"OperationLogDaoImpl.InsertOperationLog",
			zap.Field{Key: "operationLog", Type: zapcore.ObjectMarshalerType, Interface: operationLog},
			zap.Field{Key: "operationLogID", Type: zapcore.StringType, String: operationLog.OperationLogID.Hex()},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
	} else {
		o.Dao.Zap.Logger.Info(
			"OperationLogDaoImpl.InsertOperationLog",
			zap.Field{Key: "operationLog", Type: zapcore.ObjectMarshalerType, Interface: operationLog},
			zap.Field{Key: "operationLogID", Type: zapcore.StringType, String: operationLog.OperationLogID.Hex()},
		)
	}
	return result.InsertedID.(primitive.ObjectID), err
}

func (o OperationLogDaoImpl) DeleteOperationLog(operationLogID primitive.ObjectID, ctx context.Context) error {
	collection := o.Dao.Mongo.MongoDatabase.Collection(operationLogCollectionName)
	err := collection.RemoveId(ctx, operationLogID)
	if err != nil {
		o.Dao.Zap.Logger.Error(
			"OperationLogDaoImpl.DeleteOperationLog",
			zap.Field{Key: "operationLogID", Type: zapcore.StringType, String: operationLogID.Hex()},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
	} else {
		o.Dao.Zap.Logger.Info(
			"OperationLogDaoImpl.DeleteOperationLog",
			zap.Field{Key: "operationLogID", Type: zapcore.StringType, String: operationLogID.Hex()},
		)
	}
	return err
}
