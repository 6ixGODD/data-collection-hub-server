package mods

import (
	"context"
	"time"

	"data-collection-hub-server/internal/pkg/dal"
	"data-collection-hub-server/internal/pkg/models"
	"github.com/goccy/go-json"
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

type OperationLogDaoImpl struct{ *dal.Dao }

func NewOperationLogDao(dao *dal.Dao) OperationLogDao {
	var _ OperationLogDao = (*OperationLogDaoImpl)(nil) // Ensure that the interface is implemented
	return &OperationLogDaoImpl{dao}
}

func (o *OperationLogDaoImpl) GetOperationLogById(
	ctx context.Context, operationLogID primitive.ObjectID,
) (*models.OperationLogModel, error) {
	collection := o.Dao.Mongo.MongoDatabase.Collection(operationLogCollectionName)
	var operationLog models.OperationLogModel
	err := collection.Find(ctx, bson.M{"_id": operationLogID}).One(&operationLog)
	if err != nil {
		o.Dao.Zap.Logger.Error(
			"OperationLogDaoImpl.GetOperationLogById",
			zap.String("operationLogID", operationLogID.Hex()),
			zap.Error(err),
		)
		return nil, err
	} else {
		o.Dao.Zap.Logger.Info(
			"OperationLogDaoImpl.GetOperationLogById",
			zap.String("operationLogID", operationLogID.Hex()),
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
	collection := o.Dao.Mongo.MongoDatabase.Collection(operationLogCollectionName)
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
		o.Dao.Zap.Logger.Error(
			"OperationLogDaoImpl.GetOperationLogList",
			zap.Int64("offset", offset), zap.Int64("limit", limit),
			zap.Bool("desc", desc),
			zap.ByteString(operationLogCollectionName, docJSON),
			zap.Error(err),
		)
		return nil, nil, err
	}
	count, err := collection.Find(ctx, doc).Count()
	if err != nil {
		o.Dao.Zap.Logger.Error(
			"OperationLogDaoImpl.GetOperationLogList",
			zap.Int64("offset", offset), zap.Int64("limit", limit),
			zap.Bool("desc", desc),
			zap.ByteString(operationLogCollectionName, docJSON),
			zap.Error(err),
		)
		return nil, nil, err
	}
	o.Dao.Zap.Logger.Info(
		"OperationLogDaoImpl.GetOperationLogList",
		zap.Int64("offset", offset), zap.Int64("limit", limit),
		zap.Bool("desc", desc),
		zap.ByteString(operationLogCollectionName, docJSON),
		zap.Int64("count", count),
	)
	return operationLogList, &count, nil
}

func (o *OperationLogDaoImpl) GetOperationLogListByUserID(
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

func (o *OperationLogDaoImpl) GetOperationLogListByIPAddress(
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

func (o *OperationLogDaoImpl) GetOperationLogListByOperation(
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

func (o *OperationLogDaoImpl) GetOperationLogListByEntityID(
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

func (o *OperationLogDaoImpl) GetOperationLogListByEntityType(
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

func (o *OperationLogDaoImpl) GetOperationLogListByStatus(
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

func (o *OperationLogDaoImpl) GetOperationLogListByCreatedTime(
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

func (o *OperationLogDaoImpl) InsertOperationLog(
	ctx context.Context,
	userID, entityID primitive.ObjectID,
	username, email, ipAddress, userAgent, operation, entityType, description, status string,
) (primitive.ObjectID, error) {
	collection := o.Dao.Mongo.MongoDatabase.Collection(operationLogCollectionName)
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
		o.Dao.Zap.Logger.Error(
			"OperationLogDaoImpl.InsertOperationLog",
			zap.ByteString("doc", docJSON),
			zap.Error(err),
		)
	} else {
		o.Dao.Zap.Logger.Info(
			"OperationLogDaoImpl.InsertOperationLog",
			zap.ByteString("doc", docJSON),
			zap.String("operationLogID", result.InsertedID.(primitive.ObjectID).Hex()),
		)
	}
	return result.InsertedID.(primitive.ObjectID), err
}

func (o *OperationLogDaoImpl) DeleteOperationLog(ctx context.Context, operationLogID primitive.ObjectID) error {
	collection := o.Dao.Mongo.MongoDatabase.Collection(operationLogCollectionName)
	err := collection.RemoveId(ctx, operationLogID)
	if err != nil {
		o.Dao.Zap.Logger.Error(
			"OperationLogDaoImpl.DeleteOperationLog",
			zap.String("operationLogID", operationLogID.Hex()),
			zap.Error(err),
		)
	} else {
		o.Dao.Zap.Logger.Info(
			"OperationLogDaoImpl.DeleteOperationLog",
			zap.String("operationLogID", operationLogID.Hex()),
		)
	}
	return err
}

func (o *OperationLogDaoImpl) DeleteOperationLogList(
	ctx context.Context, startTime, endTime *time.Time, userID, entityID *primitive.ObjectID,
	ipAddress, operation, entityType, status *string,
) (*int64, error) {
	collection := o.Dao.Mongo.MongoDatabase.Collection(operationLogCollectionName)
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
		o.Dao.Zap.Logger.Error(
			"OperationLogDaoImpl.DeleteOperationLogList",
			zap.ByteString(operationLogCollectionName, docJSON),
			zap.Error(err),
		)
	} else {
		o.Dao.Zap.Logger.Info(
			"OperationLogDaoImpl.DeleteOperationLogList",
			zap.ByteString(operationLogCollectionName, docJSON),
			zap.Int64("count", result.DeletedCount),
		)
	}
	return &result.DeletedCount, err
}
