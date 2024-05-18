package mods

import (
	"context"
	"fmt"
	"time"

	"data-collection-hub-server/internal/pkg/config"
	"data-collection-hub-server/internal/pkg/dao"
	"data-collection-hub-server/internal/pkg/models"
	"github.com/goccy/go-json"
	"github.com/qiniu/qmgo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type OperationLogDao interface {
	GetOperationLogById(ctx context.Context, operationLogID primitive.ObjectID) (*models.OperationLogModel, error)
	GetOperationLogList(
		ctx context.Context,
		offset, limit int64, desc bool, startTime, endTime *time.Time, userID, entityID *primitive.ObjectID,
		ipAddress, operation, entityType, status, query *string,
	) ([]models.OperationLogModel, *int64, error)
	InsertOperationLog(
		ctx context.Context,
		userID, entityID primitive.ObjectID,
		username, email, ipAddress, userAgent, operation, entityType, description, status string,
	) (primitive.ObjectID, error)
	DeleteOperationLog(ctx context.Context, operationLogID primitive.ObjectID) error
	DeleteOperationLogList(
		ctx context.Context, startTime, endTime *time.Time, userID, entityID *primitive.ObjectID,
		ipAddress, operation, entityType, status *string,
	) (*int64, error)
}

type OperationLogDaoImpl struct{ *dao.Dao }

func NewOperationLogDao(ctx context.Context, dao *dao.Dao) (OperationLogDao, error) {
	var _ OperationLogDao = (*OperationLogDaoImpl)(nil) // Ensure that the interface is implemented
	collection := dao.Mongo.MongoClient.Database(dao.Mongo.DatabaseName).Collection(config.OperationLogCollectionName)
	err := collection.CreateIndexes(
		ctx, []options.IndexModel{
			{Key: []string{"created_at"}}, {Key: []string{"operation"}}, {Key: []string{"entity_type"}},
			{Key: []string{"status"}},
		},
	)
	if err != nil {
		dao.Logger.Error(
			fmt.Sprintf("Failed to create index for %s", config.OperationLogCollectionName),
			zap.Error(err),
		)
		return nil, err
	}
	return &OperationLogDaoImpl{dao}, nil
}

func (o *OperationLogDaoImpl) GetOperationLogById(
	ctx context.Context, operationLogID primitive.ObjectID,
) (*models.OperationLogModel, error) {
	collection := o.Dao.Mongo.MongoClient.Database(o.Dao.Mongo.DatabaseName).Collection(config.OperationLogCollectionName)
	var operationLog models.OperationLogModel
	err := collection.Find(ctx, bson.M{"_id": operationLogID}).One(&operationLog)
	if err != nil {
		o.Dao.Logger.Error(
			"OperationLogDaoImpl.GetOperationLogById: error", zap.Error(err),
			zap.String("operationLogID", operationLogID.Hex()),
		)
		return nil, err
	} else {
		o.Dao.Logger.Info(
			"OperationLogDaoImpl.GetOperationLogById: success", zap.String("operationLogID", operationLogID.Hex()),
		)
		return &operationLog, nil
	}
}

func (o *OperationLogDaoImpl) GetOperationLogList(
	ctx context.Context,
	offset, limit int64, desc bool, startTime, endTime *time.Time, userID, entityID *primitive.ObjectID,
	ipAddress, operation, entityType, status, query *string,
) ([]models.OperationLogModel, *int64, error) {
	var operationLogList []models.OperationLogModel
	var err error
	collection := o.Dao.Mongo.MongoClient.Database(o.Dao.Mongo.DatabaseName).Collection(config.OperationLogCollectionName)
	doc := bson.M{}
	if startTime != nil && endTime != nil {
		doc["created_at"] = bson.M{"$gte": startTime, "$lte": endTime}
	}
	if userID != nil {
		doc["user_id"] = *userID
	}
	if entityID != nil {
		doc["entity_id"] = *entityID
	}
	if ipAddress != nil {
		doc["ip_address"] = *ipAddress
	}
	if operation != nil {
		doc["operation"] = *operation
	}
	if entityType != nil {
		doc["entity_type"] = *entityType
	}
	if status != nil {
		doc["status"] = *status
	}
	if query != nil {
		doc["$or"] = []bson.M{
			{"username": bson.M{"$regex": *query, "$options": "i"}},
			{"email": bson.M{"$regex": *query, "$options": "i"}},
			{"ip_address": bson.M{"$regex": *query, "$options": "i"}},
			{"operation": bson.M{"$regex": *query, "$options": "i"}},
			{"entity_type": bson.M{"$regex": *query, "$options": "i"}},
			{"description": bson.M{"$regex": *query, "$options": "i"}},
			{"status": bson.M{"$regex": *query, "$options": "i"}},
		}
	}
	docJSON, _ := json.Marshal(doc)
	if desc {
		err = collection.Find(ctx, doc).Sort("-created_at").Skip(offset).Limit(limit).All(&operationLogList)
	} else {
		err = collection.Find(ctx, doc).Skip(offset).Limit(limit).All(&operationLogList)
	}

	if err != nil {
		o.Dao.Logger.Error(
			"OperationLogDaoImpl.GetOperationLogList",
			zap.Error(err), zap.ByteString(config.OperationLogCollectionName, docJSON),
		)
		return nil, nil, err
	}
	count, err := collection.Find(ctx, doc).Count()
	if err != nil {
		o.Dao.Logger.Error(
			"OperationLogDaoImpl.GetOperationLogList",
			zap.Error(err), zap.ByteString(config.OperationLogCollectionName, docJSON),
		)
		return nil, nil, err
	}
	o.Dao.Logger.Info(
		"OperationLogDaoImpl.GetOperationLogList",
		zap.Int64("count", count), zap.ByteString(config.OperationLogCollectionName, docJSON),
	)
	return operationLogList, &count, nil
}

func (o *OperationLogDaoImpl) GetOperationLogListByUserID(
	userID primitive.ObjectID, offset, limit int64, ctx context.Context,
) ([]models.OperationLogModel, error) {
	var operationLogList []models.OperationLogModel
	collection := o.Dao.Mongo.MongoClient.Database(o.Dao.Mongo.DatabaseName).Collection(config.OperationLogCollectionName)
	err := collection.Find(ctx, bson.M{"user_id": userID}).Skip(offset).Limit(limit).All(&operationLogList)
	if err != nil {
		o.Dao.Logger.Error(
			"OperationLogDaoImpl.GetOperationLogListByUserID: failed to find operation logs",
			zap.Field{Key: "userID", Type: zapcore.StringType, String: userID.Hex()},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		o.Dao.Logger.Info(
			"OperationLogDaoImpl.GetOperationLogListByUserID: success",
			zap.Field{Key: "userID", Type: zapcore.StringType, String: userID.Hex()},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return operationLogList, nil
	}
}

func (o *OperationLogDaoImpl) GetOperationLogListByIPAddress(
	ipAddress string, offset, limit int64, ctx context.Context,
) ([]models.OperationLogModel, error) {
	var operationLogList []models.OperationLogModel
	collection := o.Dao.Mongo.MongoClient.Database(o.Dao.Mongo.DatabaseName).Collection(config.OperationLogCollectionName)
	err := collection.Find(ctx, bson.M{"ip_address": ipAddress}).Skip(offset).Limit(limit).All(&operationLogList)
	if err != nil {
		o.Dao.Logger.Error(
			"OperationLogDaoImpl.GetOperationLogListByIPAddress: failed to find operation logs",
			zap.Field{Key: "ipAddress", Type: zapcore.StringType, String: ipAddress},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		o.Dao.Logger.Info(
			"OperationLogDaoImpl.GetOperationLogListByIPAddress: success",
			zap.Field{Key: "ipAddress", Type: zapcore.StringType, String: ipAddress},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return operationLogList, nil
	}
}

func (o *OperationLogDaoImpl) GetOperationLogListByOperation(
	operation string, offset, limit int64, ctx context.Context,
) ([]models.OperationLogModel, error) {
	var operationLogList []models.OperationLogModel
	collection := o.Dao.Mongo.MongoClient.Database(o.Dao.Mongo.DatabaseName).Collection(config.OperationLogCollectionName)
	err := collection.Find(ctx, bson.M{"operation": operation}).Skip(offset).Limit(limit).All(&operationLogList)
	if err != nil {
		o.Dao.Logger.Error(
			"OperationLogDaoImpl.GetOperationLogListByOperation: failed to find operation logs",
			zap.Field{Key: "operation", Type: zapcore.StringType, String: operation},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		o.Dao.Logger.Info(
			"OperationLogDaoImpl.GetOperationLogListByOperation: success",
			zap.Field{Key: "operation", Type: zapcore.StringType, String: operation},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return operationLogList, nil
	}
}

func (o *OperationLogDaoImpl) GetOperationLogListByEntityID(
	entityID primitive.ObjectID, offset, limit int64, ctx context.Context,
) ([]models.OperationLogModel, error) {
	var operationLogList []models.OperationLogModel
	collection := o.Dao.Mongo.MongoClient.Database(o.Dao.Mongo.DatabaseName).Collection(config.OperationLogCollectionName)
	err := collection.Find(ctx, bson.M{"entity_id": entityID}).Skip(offset).Limit(limit).All(&operationLogList)
	if err != nil {
		o.Dao.Logger.Error(
			"OperationLogDaoImpl.GetOperationLogListByEntityID: failed to find operation logs",
			zap.Field{Key: "entityID", Type: zapcore.StringType, String: entityID.Hex()},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		o.Dao.Logger.Info(
			"OperationLogDaoImpl.GetOperationLogListByEntityID: success",
			zap.Field{Key: "entityID", Type: zapcore.StringType, String: entityID.Hex()},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return operationLogList, nil
	}
}

func (o *OperationLogDaoImpl) GetOperationLogListByEntityType(
	entityType string, offset, limit int64, ctx context.Context,
) ([]models.OperationLogModel, error) {
	var operationLogList []models.OperationLogModel
	collection := o.Dao.Mongo.MongoClient.Database(o.Dao.Mongo.DatabaseName).Collection(config.OperationLogCollectionName)
	err := collection.Find(ctx, bson.M{"entity_type": entityType}).Skip(offset).Limit(limit).All(&operationLogList)
	if err != nil {
		o.Dao.Logger.Error(
			"OperationLogDaoImpl.GetOperationLogListByEntityType: failed to find operation logs",
			zap.Field{Key: "entityType", Type: zapcore.StringType, String: entityType},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		o.Dao.Logger.Info(
			"OperationLogDaoImpl.GetOperationLogListByEntityType: success",
			zap.Field{Key: "entityType", Type: zapcore.StringType, String: entityType},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return operationLogList, nil
	}
}

func (o *OperationLogDaoImpl) GetOperationLogListByStatus(
	status string, offset, limit int64, ctx context.Context,
) ([]models.OperationLogModel, error) {
	var operationLogList []models.OperationLogModel
	collection := o.Dao.Mongo.MongoClient.Database(o.Dao.Mongo.DatabaseName).Collection(config.OperationLogCollectionName)
	err := collection.Find(ctx, bson.M{"status": status}).Skip(offset).Limit(limit).All(&operationLogList)
	if err != nil {
		o.Dao.Logger.Error(
			"OperationLogDaoImpl.GetOperationLogListByStatus: failed to find operation logs",
			zap.Field{Key: "status", Type: zapcore.StringType, String: status},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		o.Dao.Logger.Info(
			"OperationLogDaoImpl.GetOperationLogListByStatus: success",
			zap.Field{Key: "status", Type: zapcore.StringType, String: status},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return operationLogList, nil
	}
}

func (o *OperationLogDaoImpl) GetOperationLogListByCreatedTime(
	startTime, endTime time.Time, offset, limit int64, ctx context.Context,
) ([]models.OperationLogModel, error) {
	var operationLogList []models.OperationLogModel
	collection := o.Dao.Mongo.MongoClient.Database(o.Dao.Mongo.DatabaseName).Collection(config.OperationLogCollectionName)
	err := collection.Find(
		ctx, bson.M{"created_time": bson.M{"$gte": startTime, "$lte": endTime}},
	).Skip(offset).Limit(limit).All(&operationLogList)
	if err != nil {
		o.Dao.Logger.Error(
			"OperationLogDaoImpl.GetOperationLogListByCreatedTime: failed to find operation logs",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime.Format(time.RFC3339)},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime.Format(time.RFC3339)},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		o.Dao.Logger.Info(
			"OperationLogDaoImpl.GetOperationLogListByCreatedTime: success",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime.Format(time.RFC3339)},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime.Format(time.RFC3339)},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return operationLogList, nil
	}
}

func (o *OperationLogDaoImpl) InsertOperationLog(
	ctx context.Context,
	userID, entityID primitive.ObjectID,
	username, email, ipAddress, userAgent, operation, entityType, description, status string,
) (primitive.ObjectID, error) {
	collection := o.Dao.Mongo.MongoClient.Database(o.Dao.Mongo.DatabaseName).Collection(config.OperationLogCollectionName)
	doc := bson.M{
		"user_id":     userID,
		"entity_id":   entityID,
		"username":    username,
		"email":       email,
		"ip_address":  ipAddress,
		"user_agent":  userAgent,
		"operation":   operation,
		"entity_type": entityType,
		"description": description,
		"status":      status,
		"created_at":  time.Now(),
	}
	docJSON, _ := json.Marshal(doc)
	result, err := collection.InsertOne(ctx, doc)
	if err != nil {
		o.Dao.Logger.Error(
			"OperationLogDaoImpl.InsertOperationLog: failed to insert operation log",
			zap.ByteString("doc", docJSON), zap.Error(err),
		)
	} else {
		o.Dao.Logger.Info(
			"OperationLogDaoImpl.InsertOperationLog: success",
			zap.ByteString("doc", docJSON), zap.String("operationLogID", result.InsertedID.(primitive.ObjectID).Hex()),
		)
	}
	return result.InsertedID.(primitive.ObjectID), err
}

func (o *OperationLogDaoImpl) DeleteOperationLog(ctx context.Context, operationLogID primitive.ObjectID) error {
	collection := o.Dao.Mongo.MongoClient.Database(o.Dao.Mongo.DatabaseName).Collection(config.OperationLogCollectionName)
	err := collection.RemoveId(ctx, operationLogID)
	if err != nil {
		o.Dao.Logger.Error(
			"OperationLogDaoImpl.DeleteOperationLog: failed to delete operation log",
			zap.Error(err), zap.String("operationLogID", operationLogID.Hex()),
		)
	} else {
		o.Dao.Logger.Info(
			"OperationLogDaoImpl.DeleteOperationLog: success", zap.String("operationLogID", operationLogID.Hex()),
		)
	}
	return err
}

func (o *OperationLogDaoImpl) DeleteOperationLogList(
	ctx context.Context, startTime, endTime *time.Time, userID, entityID *primitive.ObjectID,
	ipAddress, operation, entityType, status *string,
) (*int64, error) {
	collection := o.Dao.Mongo.MongoClient.Database(o.Dao.Mongo.DatabaseName).Collection(config.OperationLogCollectionName)
	doc := bson.M{}
	if startTime != nil && endTime != nil {
		doc["created_at"] = bson.M{"$gte": startTime, "$lte": endTime}
	}
	if userID != nil {
		doc["user_id"] = *userID
	}
	if entityID != nil {
		doc["entity_id"] = *entityID
	}
	if ipAddress != nil {
		doc["ip_address"] = *ipAddress
	}
	if operation != nil {
		doc["operation"] = *operation
	}
	if entityType != nil {
		doc["entity_type"] = *entityType
	}
	if status != nil {
		doc["status"] = *status
	}

	docJSON, _ := json.Marshal(doc)
	result, err := collection.RemoveAll(ctx, doc)
	if err != nil {
		o.Dao.Logger.Error(
			"OperationLogDaoImpl.DeleteOperationLogList: failed to delete operation logs",
			zap.Error(err), zap.ByteString(config.OperationLogCollectionName, docJSON),
		)
	} else {
		o.Dao.Logger.Info(
			"OperationLogDaoImpl.DeleteOperationLogList: success",
			zap.Int64("count", result.DeletedCount),
			zap.ByteString(config.OperationLogCollectionName, docJSON),
		)
	}
	return &result.DeletedCount, err
}
