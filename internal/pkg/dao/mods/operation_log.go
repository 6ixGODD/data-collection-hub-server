package mods

import (
	"context"
	"errors"
	"fmt"
	"time"

	"data-collection-hub-server/internal/pkg/config"
	"data-collection-hub-server/internal/pkg/dao"
	entity2 "data-collection-hub-server/internal/pkg/domain/entity"
	"github.com/goccy/go-json"
	"github.com/qiniu/qmgo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

type OperationLogDao interface {
	GetOperationLogByID(ctx context.Context, operationLogID primitive.ObjectID) (*entity2.OperationLogModel, error)
	GetOperationLogList(
		ctx context.Context,
		offset, limit int64, desc bool, startTime, endTime *time.Time, userID, entityID *primitive.ObjectID,
		ipAddress, operation, entityType, status, query *string,
	) ([]entity2.OperationLogModel, *int64, error)
	InsertOperationLog(
		ctx context.Context,
		userID, entityID primitive.ObjectID,
		username, email, ipAddress, userAgent, operation, entityType, description, status string,
	) (primitive.ObjectID, error)
	CacheOperationLog(
		ctx context.Context, userID, entityID primitive.ObjectID,
		ipAddress, userAgent, operation, entityType, description, status string,
	) error
	SyncOperationLog(ctx context.Context)
	DeleteOperationLog(ctx context.Context, operationLogID primitive.ObjectID) error
	DeleteOperationLogList(
		ctx context.Context, startTime, endTime *time.Time, userID, entityID *primitive.ObjectID,
		ipAddress, operation, entityType, status *string,
	) (*int64, error)
}

type OperationLogDaoImpl struct {
	core    *dao.Core
	cache   *dao.Cache
	userDao UserDao
}

func NewOperationLogDao(ctx context.Context, core *dao.Core, cache *dao.Cache, userDao UserDao) (
	OperationLogDao, error,
) {
	var _ OperationLogDao = (*OperationLogDaoImpl)(nil) // Ensure that the interface is implemented
	collection := core.Mongo.MongoClient.Database(core.Mongo.DatabaseName).Collection(config.OperationLogCollectionName)
	err := collection.CreateIndexes(
		ctx, []options.IndexModel{
			{Key: []string{"created_at"}}, {Key: []string{"operation"}}, {Key: []string{"entity_type"}},
			{Key: []string{"status"}},
		},
	)
	if err != nil {
		core.Logger.Error(
			fmt.Sprintf("Failed to create index for %s", config.OperationLogCollectionName),
			zap.Error(err),
		)
		return nil, err
	}
	return &OperationLogDaoImpl{
		core:    core,
		cache:   cache,
		userDao: userDao,
	}, nil
}

func (o *OperationLogDaoImpl) GetOperationLogByID(
	ctx context.Context, operationLogID primitive.ObjectID,
) (*entity2.OperationLogModel, error) {
	collection := o.core.Mongo.MongoClient.Database(o.core.Mongo.DatabaseName).Collection(config.OperationLogCollectionName)
	var operationLog entity2.OperationLogModel
	err := collection.Find(ctx, bson.M{"_id": operationLogID}).One(&operationLog)
	if err != nil {
		o.core.Logger.Error(
			"OperationLogDaoImpl.GetOperationLogByID: error", zap.Error(err),
			zap.String("operationLogID", operationLogID.Hex()),
		)
		return nil, err
	} else {
		o.core.Logger.Info(
			"OperationLogDaoImpl.GetOperationLogByID: success", zap.String("operationLogID", operationLogID.Hex()),
		)
		return &operationLog, nil
	}
}

func (o *OperationLogDaoImpl) GetOperationLogList(
	ctx context.Context,
	offset, limit int64, desc bool, startTime, endTime *time.Time, userID, entityID *primitive.ObjectID,
	ipAddress, operation, entityType, status, query *string,
) ([]entity2.OperationLogModel, *int64, error) {
	var operationLogList []entity2.OperationLogModel
	var err error
	collection := o.core.Mongo.MongoClient.Database(o.core.Mongo.DatabaseName).Collection(config.OperationLogCollectionName)
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
		o.core.Logger.Error(
			"OperationLogDaoImpl.GetOperationLogList",
			zap.Error(err), zap.ByteString(config.OperationLogCollectionName, docJSON),
		)
		return nil, nil, err
	}
	count, err := collection.Find(ctx, doc).Count()
	if err != nil {
		o.core.Logger.Error(
			"OperationLogDaoImpl.GetOperationLogList",
			zap.Error(err), zap.ByteString(config.OperationLogCollectionName, docJSON),
		)
		return nil, nil, err
	}
	o.core.Logger.Info(
		"OperationLogDaoImpl.GetOperationLogList",
		zap.Int64("count", count), zap.ByteString(config.OperationLogCollectionName, docJSON),
	)
	return operationLogList, &count, nil
}

func (o *OperationLogDaoImpl) InsertOperationLog(
	ctx context.Context,
	userID, entityID primitive.ObjectID,
	username, email, ipAddress, userAgent, operation, entityType, description, status string,
) (primitive.ObjectID, error) {
	collection := o.core.Mongo.MongoClient.Database(o.core.Mongo.DatabaseName).Collection(config.OperationLogCollectionName)
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
		o.core.Logger.Error(
			"OperationLogDaoImpl.InsertOperationLog: failed to insert operation log",
			zap.ByteString("doc", docJSON), zap.Error(err),
		)
	} else {
		o.core.Logger.Info(
			"OperationLogDaoImpl.InsertOperationLog: success",
			zap.ByteString("doc", docJSON), zap.String("operationLogID", result.InsertedID.(primitive.ObjectID).Hex()),
		)
	}
	return result.InsertedID.(primitive.ObjectID), err
}

func (o *OperationLogDaoImpl) CacheOperationLog(
	ctx context.Context, userID, entityID primitive.ObjectID,
	ipAddress, userAgent, operation, entityType, description, status string,
) error {
	user, err := o.userDao.GetUserByID(ctx, userID)
	if err != nil {
		o.core.Logger.Error(
			"OperationLogDaoImpl.CacheOperationLog: failed to get user",
			zap.Error(err), zap.String("userID", userID.Hex()),
		)
		return err
	}
	operationLog := entity2.OperationLogCache{
		UserIDHex:   user.UserID.Hex(),
		Username:    user.Username,
		Email:       user.Email,
		IPAddress:   ipAddress,
		UserAgent:   userAgent,
		Operation:   operation,
		EntityIDHex: entityID.Hex(),
		EntityType:  entityType,
		Description: description,
		Status:      status,
		CreatedAt:   time.Now(),
	}
	operationLogJSON, err := json.Marshal(operationLog)
	if err != nil {
		o.core.Logger.Error(
			"OperationLogDaoImpl.CacheOperationLog: failed to marshal operation log",
			zap.Error(err), zap.String("userID", userID.Hex()),
		)
		return err
	}
	return o.cache.RightPush(ctx, config.OperationLogCacheKey, string(operationLogJSON))
}

func (o *OperationLogDaoImpl) SyncOperationLog(ctx context.Context) {
	for {
		operationLogJSON, err := o.cache.LeftPop(ctx, config.OperationLogCacheKey)
		if err != nil {
			if errors.Is(err, dao.CacheNil{}) {
				break
			}
			o.core.Logger.Error(
				"OperationLogDaoImpl.SyncOperationLog: failed to left pop operation log",
				zap.Error(err),
			)
		}
		var operationLog entity2.OperationLogCache
		err = json.Unmarshal([]byte(*operationLogJSON), &operationLog)
		if err != nil {
			o.core.Logger.Error(
				"OperationLogDaoImpl.SyncOperationLog: failed to unmarshal operation log",
				zap.Error(err), zap.String("operationLogJSON", *operationLogJSON),
			)
			continue
		}
		userID, err := primitive.ObjectIDFromHex(operationLog.UserIDHex)
		if err != nil {
			o.core.Logger.Error(
				"OperationLogDaoImpl.SyncOperationLog: failed to get user ID",
				zap.Error(err), zap.String("operationLogJSON", *operationLogJSON),
			)
			continue
		}
		entityID, err := primitive.ObjectIDFromHex(operationLog.EntityIDHex)
		if err != nil {
			o.core.Logger.Error(
				"OperationLogDaoImpl.SyncOperationLog: failed to get entity ID",
				zap.Error(err), zap.String("operationLogJSON", *operationLogJSON),
			)
			continue
		}
		_, err = o.InsertOperationLog(
			ctx, userID, entityID,
			operationLog.Username, operationLog.Email, operationLog.IPAddress, operationLog.UserAgent,
			operationLog.Operation, operationLog.EntityType, operationLog.Description, operationLog.Status,
		)
		if err != nil {
			o.core.Logger.Error(
				"OperationLogDaoImpl.SyncOperationLog: failed to insert operation log",
				zap.Error(err), zap.String("operationLogJSON", *operationLogJSON),
			)
		}
	}
}

func (o *OperationLogDaoImpl) DeleteOperationLog(ctx context.Context, operationLogID primitive.ObjectID) error {
	collection := o.core.Mongo.MongoClient.Database(o.core.Mongo.DatabaseName).Collection(config.OperationLogCollectionName)
	err := collection.RemoveId(ctx, operationLogID)
	if err != nil {
		o.core.Logger.Error(
			"OperationLogDaoImpl.DeleteOperationLog: failed to delete operation log",
			zap.Error(err), zap.String("operationLogID", operationLogID.Hex()),
		)
	} else {
		o.core.Logger.Info(
			"OperationLogDaoImpl.DeleteOperationLog: success", zap.String("operationLogID", operationLogID.Hex()),
		)
	}
	return err
}

func (o *OperationLogDaoImpl) DeleteOperationLogList(
	ctx context.Context, startTime, endTime *time.Time, userID, entityID *primitive.ObjectID,
	ipAddress, operation, entityType, status *string,
) (*int64, error) {
	collection := o.core.Mongo.MongoClient.Database(o.core.Mongo.DatabaseName).Collection(config.OperationLogCollectionName)
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
		o.core.Logger.Error(
			"OperationLogDaoImpl.DeleteOperationLogList: failed to delete operation logs",
			zap.Error(err), zap.ByteString(config.OperationLogCollectionName, docJSON),
		)
	} else {
		o.core.Logger.Info(
			"OperationLogDaoImpl.DeleteOperationLogList: success",
			zap.Int64("count", result.DeletedCount),
			zap.ByteString(config.OperationLogCollectionName, docJSON),
		)
	}
	return &result.DeletedCount, err
}
