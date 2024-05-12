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

type ErrorLogDao interface {
	GetErrorLogById(ctx context.Context, errorLogID primitive.ObjectID) (*models.ErrorLogModel, error)
	GetErrorLogList(
		ctx context.Context,
		offset, limit int64, desc bool, startTime, endTime *time.Time, userID *primitive.ObjectID,
		ipAddress, requestURL, errorCode, query *string,
	) ([]models.ErrorLogModel, *int64, error)
	InsertErrorLog(
		ctx context.Context,
		userID primitive.ObjectID,
		Username, IPAddress, UserAgent, RequestURL, RequestMethod, RequestPayload, ErrorCode, ErrorMsg, Stack string,
	) (primitive.ObjectID, error)
	DeleteErrorLog(ctx context.Context, errorLogID primitive.ObjectID) error
	DeleteErrorLogList(
		ctx context.Context, startTime, endTime *time.Time, userID *primitive.ObjectID,
		ipAddress, requestURL, errorCode, query *string,
	) (*int64, error)
}

type ErrorLogDaoImpl struct{ *dal.Dao }

func NewErrorLogDao(dao *dal.Dao) ErrorLogDao {
	var _ ErrorLogDao = (*ErrorLogDaoImpl)(nil) // Ensure that the interface is implemented
	return &ErrorLogDaoImpl{dao}
}

func (e *ErrorLogDaoImpl) GetErrorLogById(ctx context.Context, errorLogID primitive.ObjectID) (
	*models.ErrorLogModel, error,
) {
	var errorLog models.ErrorLogModel
	collection := e.Dao.Mongo.MongoDatabase.Collection(errorLogCollectionName)
	err := collection.Find(ctx, bson.M{"_id": errorLogID}).One(&errorLog)
	if err != nil {
		e.Dao.Zap.Logger.Error(
			"ErrorLogDaoImpl.GetErrorLogById",
			zap.Field{Key: "errorLogID", Type: zapcore.StringType, String: errorLogID.Hex()},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		e.Dao.Zap.Logger.Info(
			"ErrorLogDaoImpl.GetErrorLogById",
			zap.Field{Key: "errorLogID", Type: zapcore.StringType, String: errorLogID.Hex()},
		)
		return &errorLog, nil
	}
}

func (e *ErrorLogDaoImpl) GetErrorLogList(
	ctx context.Context,
	offset, limit int64, desc bool, startTime, endTime *time.Time, userID *primitive.ObjectID,
	ipAddress, requestURL, errorCode, query *string,
) ([]models.ErrorLogModel, *int64, error) {
	var errorLogList []models.ErrorLogModel
	var err error
	collection := e.Dao.Mongo.MongoDatabase.Collection(errorLogCollectionName)

	doc := bson.M{}
	if userID != nil {
		doc["user_id"] = *userID
	}
	if ipAddress != nil {
		doc["ip_address"] = *ipAddress
	}
	if requestURL != nil {
		doc["request_url"] = *requestURL
	}
	if errorCode != nil {
		doc["error_code"] = *errorCode
	}
	if query != nil {
		doc["$or"] = []bson.M{
			{"user_id": bson.M{"$regex": *query}},
			{"username": bson.M{"$regex": *query}},
			{"ip_address": bson.M{"$regex": *query}},
			{"request_url": bson.M{"$regex": *query}},
			{"error_code": bson.M{"$regex": *query}},
			{"error_msg": bson.M{"$regex": *query}},
			{"stack": bson.M{"$regex": *query}},
		}
	}
	if startTime != nil && endTime != nil {
		doc["created_at"] = bson.M{"$gte": *startTime, "$lte": *endTime}
	}
	docJSON, _ := json.Marshal(doc)

	if desc {
		err = collection.Find(ctx, doc).Sort("-created_at").Skip(offset).Limit(limit).All(&errorLogList)
	} else {
		err = collection.Find(ctx, doc).Skip(offset).Limit(limit).All(&errorLogList)
	}
	if err != nil {
		e.Dao.Zap.Logger.Error(
			"ErrorLogDaoImpl.GetErrorLogList",
			zap.Int64("offset", offset), zap.Int64("limit", limit), zap.Bool("desc", desc),
			zap.ByteString(errorLogCollectionName, docJSON), zap.Error(err),
		)
		return nil, nil, err
	}
	count, err := collection.Find(ctx, doc).Count()
	if err != nil {
		e.Dao.Zap.Logger.Error(
			"ErrorLogDaoImpl.GetErrorLogList",
			zap.Int64("offset", offset), zap.Int64("limit", limit), zap.Bool("desc", desc),
			zap.ByteString(errorLogCollectionName, docJSON), zap.Error(err),
		)
		return nil, nil, err
	}
	e.Dao.Zap.Logger.Info(
		"ErrorLogDaoImpl.GetErrorLogList",
		zap.Int64("offset", offset), zap.Int64("limit", limit), zap.Bool("desc", desc),
		zap.ByteString(errorLogCollectionName, docJSON), zap.Int64("count", count),
	)
	return errorLogList, &count, nil

}

func (e *ErrorLogDaoImpl) InsertErrorLog(
	ctx context.Context,
	UserID primitive.ObjectID,
	Username, IPAddress, UserAgent, RequestURL, RequestMethod, RequestPayload, ErrorCode, ErrorMsg, Stack string,
) (primitive.ObjectID, error) {
	collection := e.Dao.Mongo.MongoDatabase.Collection(errorLogCollectionName)
	doc := bson.M{
		"user_id":         UserID,
		"username":        Username,
		"ip_address":      IPAddress,
		"user_agent":      UserAgent,
		"request_url":     RequestURL,
		"request_method":  RequestMethod,
		"request_payload": RequestPayload,
		"error_code":      ErrorCode,
		"error_msg":       ErrorMsg,
		"stack":           Stack,
		"created_at":      time.Now(),
	}
	docJSON, _ := json.Marshal(doc)
	result, err := collection.InsertOne(ctx, doc)
	if err != nil {
		e.Dao.Zap.Logger.Error(
			"ErrorLogDaoImpl.InsertErrorLog",
			zap.ByteString(errorLogCollectionName, docJSON), zap.Error(err),
			zap.Error(err),
		)
	} else {
		e.Dao.Zap.Logger.Info(
			"ErrorLogDaoImpl.InsertErrorLog",
			zap.ByteString(errorLogCollectionName, docJSON), zap.Error(err),
			zap.String("errorLogID", result.InsertedID.(primitive.ObjectID).Hex()),
		)
	}
	return result.InsertedID.(primitive.ObjectID), err
}

func (e *ErrorLogDaoImpl) DeleteErrorLog(ctx context.Context, errorLogID primitive.ObjectID) error {
	collection := e.Dao.Mongo.MongoDatabase.Collection(errorLogCollectionName)
	err := collection.RemoveId(ctx, errorLogID)
	if err != nil {
		e.Dao.Zap.Logger.Error(
			"ErrorLogDaoImpl.DeleteErrorLog",
			zap.String("errorLogID", errorLogID.Hex()), zap.Error(err),
		)
	} else {
		e.Dao.Zap.Logger.Info(
			"ErrorLogDaoImpl.DeleteErrorLog",
			zap.String("errorLogID", errorLogID.Hex()),
		)
	}
	return err
}

func (e *ErrorLogDaoImpl) DeleteErrorLogList(
	ctx context.Context, startTime, endTime *time.Time, userID *primitive.ObjectID,
	ipAddress, requestURL, errorCode, query *string,
) (*int64, error) {
	collection := e.Dao.Mongo.MongoDatabase.Collection(errorLogCollectionName)
	doc := bson.M{}
	if userID != nil {
		doc["user_id"] = *userID
	}
	if ipAddress != nil {
		doc["ip_address"] = *ipAddress
	}
	if requestURL != nil {
		doc["request_url"] = *requestURL
	}
	if errorCode != nil {
		doc["error_code"] = *errorCode
	}
	if query != nil {
		doc["$or"] = []bson.M{
			{"user_id": bson.M{"$regex": *query}},
			{"username": bson.M{"$regex": *query}},
			{"ip_address": bson.M{"$regex": *query}},
			{"request_url": bson.M{"$regex": *query}},
			{"error_code": bson.M{"$regex": *query}},
			{"error_msg": bson.M{"$regex": *query}},
			{"stack": bson.M{"$regex": *query}},
		}
	}
	if startTime != nil && endTime != nil {
		doc["created_at"] = bson.M{"$gte": *startTime, "$lte": *endTime}
	}
	docJSON, _ := json.Marshal(doc)
	result, err := collection.RemoveAll(ctx, doc)
	if err != nil {
		e.Dao.Zap.Logger.Error(
			"ErrorLogDaoImpl.DeleteErrorLogList",
			zap.ByteString(errorLogCollectionName, docJSON), zap.Error(err),
		)
		return nil, err
	}
	e.Dao.Zap.Logger.Info(
		"ErrorLogDaoImpl.DeleteErrorLogList",
		zap.ByteString(errorLogCollectionName, docJSON),
		zap.Int64("deletedCount", result.DeletedCount),
	)
	return &result.DeletedCount, nil
}
