package modules

import (
	"context"

	"data-collection-hub-server/internal/pkg/dal"
	"data-collection-hub-server/internal/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var errorLogCollectionName = "error_log"

type ErrorLogDao interface {
	GetErrorLogById(errorLogId string, ctx context.Context) (*models.ErrorLogModel, error)
	GetErrorLogList(offset, limit int64, desc bool, ctx context.Context) ([]models.ErrorLogModel, error)
	GetErrorLogListByCreatedTime(startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.ErrorLogModel, error)
	GetErrorLogListByUserUUID(userUUID string, offset, limit int64, desc bool, ctx context.Context) ([]models.ErrorLogModel, error)
	GetErrorLogListByIPAddress(ipAddress string, offset, limit int64, ctx context.Context) ([]models.ErrorLogModel, error)
	GetErrorLogListByRequestURL(requestURL string, offset, limit int64, ctx context.Context) ([]models.ErrorLogModel, error)
	GetErrorLogListByErrorCode(errorCode string, offset, limit int64, ctx context.Context) ([]models.ErrorLogModel, error)
	GetErrorLogListByFuzzyQuery(query string, offset, limit int64, ctx context.Context) ([]models.ErrorLogModel, error)
	InsertErrorLog(errorLog *models.ErrorLogModel, ctx context.Context) error
	DeleteErrorLog(errorLog *models.ErrorLogModel, ctx context.Context) error
}

type ErrorLogDaoImpl struct{ *dal.Core }

func NewErrorLogDao(core *dal.Core) ErrorLogDao {
	var _ ErrorLogDao = (*ErrorLogDaoImpl)(nil) // Ensure that the interface is implemented
	return &ErrorLogDaoImpl{core}
}

func (e *ErrorLogDaoImpl) GetErrorLogById(errorLogId string, ctx context.Context) (*models.ErrorLogModel, error) {
	var errorLog models.ErrorLogModel
	collection := e.Core.Mongo.MongoDatabase.Collection(errorLogCollectionName)
	err := collection.Find(ctx, bson.M{"_id": errorLogId}).One(&errorLog)
	if err != nil {
		e.Core.Zap.Logger.Error(
			"ErrorLogDaoImpl.GetErrorLogById",
			zap.Field{Key: "errorLogId", Type: zapcore.StringType, String: errorLogId},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		e.Core.Zap.Logger.Info(
			"ErrorLogDaoImpl.GetErrorLogById",
			zap.Field{Key: "errorLogId", Type: zapcore.StringType, String: errorLogId},
		)
		return &errorLog, nil
	}
}

func (e *ErrorLogDaoImpl) GetErrorLogList(offset, limit int64, desc bool, ctx context.Context) ([]models.ErrorLogModel, error) {
	var errorLogList []models.ErrorLogModel
	var err error
	collection := e.Core.Mongo.MongoDatabase.Collection(errorLogCollectionName)
	if desc {
		err = collection.Find(ctx, bson.M{}).Sort("-created_at").Skip(offset).Limit(limit).All(&errorLogList)
	} else {
		err = collection.Find(ctx, bson.M{}).Skip(offset).Limit(limit).All(&errorLogList)
	}
	if err != nil {
		e.Core.Zap.Logger.Error(
			"ErrorLogDaoImpl.GetErrorLogList",
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		e.Core.Zap.Logger.Info(
			"ErrorLogDaoImpl.GetErrorLogList",
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Integer: limit},
		)
		return errorLogList, nil
	}
}

func (e *ErrorLogDaoImpl) GetErrorLogListByCreatedTime(startTime, endTime string, offset, limit int64, desc bool, ctx context.Context) ([]models.ErrorLogModel, error) {
	var errorLogList []models.ErrorLogModel
	var err error
	collection := e.Core.Mongo.MongoDatabase.Collection(errorLogCollectionName)
	if desc {
		err = collection.Find(
			ctx, bson.M{"created_at": bson.M{"$gte": startTime, "$lte": endTime}},
		).Sort("-created_at").Skip(offset).Limit(limit).All(&errorLogList)
	} else {
		err = collection.Find(
			ctx, bson.M{"created_at": bson.M{"$gte": startTime, "$lte": endTime}},
		).Skip(offset).Limit(limit).All(&errorLogList)
	}
	if err != nil {
		e.Core.Zap.Logger.Error(
			"ErrorLogDaoImpl.GetErrorLogListByCreatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		e.Core.Zap.Logger.Info(
			"ErrorLogDaoImpl.GetErrorLogListByCreatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Integer: limit},
		)
		return errorLogList, nil
	}
}

func (e *ErrorLogDaoImpl) GetErrorLogListByUserUUID(userUUID string, offset, limit int64, desc bool, ctx context.Context) ([]models.ErrorLogModel, error) {
	var errorLogList []models.ErrorLogModel
	var err error
	collection := e.Core.Mongo.MongoDatabase.Collection(errorLogCollectionName)
	if desc {
		err = collection.Find(ctx, bson.M{"user_uuid": userUUID}).Sort("-created_at").Skip(offset).Limit(limit).All(&errorLogList)
	} else {
		err = collection.Find(ctx, bson.M{"user_uuid": userUUID}).Skip(offset).Limit(limit).All(&errorLogList)
	}
	if err != nil {
		e.Core.Zap.Logger.Error(
			"ErrorLogDaoImpl.GetErrorLogListByUserUUID",
			zap.Field{Key: "userUUID", Type: zapcore.StringType, String: userUUID},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		e.Core.Zap.Logger.Info(
			"ErrorLogDaoImpl.GetErrorLogListByUserUUID",
			zap.Field{Key: "userUUID", Type: zapcore.StringType, String: userUUID},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Integer: limit},
		)
		return errorLogList, nil
	}
}

func (e *ErrorLogDaoImpl) GetErrorLogListByIPAddress(ipAddress string, offset, limit int64, ctx context.Context) ([]models.ErrorLogModel, error) {
	var errorLogList []models.ErrorLogModel
	collection := e.Core.Mongo.MongoDatabase.Collection(errorLogCollectionName)
	err := collection.Find(ctx, bson.M{"ip_address": ipAddress}).Skip(offset).Limit(limit).All(&errorLogList)
	if err != nil {
		e.Core.Zap.Logger.Error(
			"ErrorLogDaoImpl.GetErrorLogListByIPAddress",
			zap.Field{Key: "ipAddress", Type: zapcore.StringType, String: ipAddress},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		e.Core.Zap.Logger.Info(
			"ErrorLogDaoImpl.GetErrorLogListByIPAddress",
			zap.Field{Key: "ipAddress", Type: zapcore.StringType, String: ipAddress},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return errorLogList, nil
	}
}

func (e *ErrorLogDaoImpl) GetErrorLogListByRequestURL(requestURL string, offset, limit int64, ctx context.Context) ([]models.ErrorLogModel, error) {
	var errorLogList []models.ErrorLogModel
	collection := e.Core.Mongo.MongoDatabase.Collection(errorLogCollectionName)
	err := collection.Find(ctx, bson.M{"request_url": requestURL}).Skip(offset).Limit(limit).All(&errorLogList)
	if err != nil {
		e.Core.Zap.Logger.Error(
			"ErrorLogDaoImpl.GetErrorLogListByRequestURL",
			zap.Field{Key: "requestURL", Type: zapcore.StringType, String: requestURL},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		e.Core.Zap.Logger.Info(
			"ErrorLogDaoImpl.GetErrorLogListByRequestURL",
			zap.Field{Key: "requestURL", Type: zapcore.StringType, String: requestURL},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return errorLogList, nil
	}
}

func (e *ErrorLogDaoImpl) GetErrorLogListByErrorCode(errorCode string, offset, limit int64, ctx context.Context) ([]models.ErrorLogModel, error) {
	var errorLogList []models.ErrorLogModel
	collection := e.Core.Mongo.MongoDatabase.Collection(errorLogCollectionName)
	err := collection.Find(ctx, bson.M{"error_code": errorCode}).Skip(offset).Limit(limit).All(&errorLogList)
	if err != nil {
		e.Core.Zap.Logger.Error(
			"ErrorLogDaoImpl.GetErrorLogListByErrorCode",
			zap.Field{Key: "errorCode", Type: zapcore.StringType, String: errorCode},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		e.Core.Zap.Logger.Info(
			"ErrorLogDaoImpl.GetErrorLogListByErrorCode",
			zap.Field{Key: "errorCode", Type: zapcore.StringType, String: errorCode},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return errorLogList, nil
	}
}

func (e *ErrorLogDaoImpl) GetErrorLogListByFuzzyQuery(query string, offset, limit int64, ctx context.Context) ([]models.ErrorLogModel, error) {
	var errorLogList []models.ErrorLogModel
	collection := e.Core.Mongo.MongoDatabase.Collection(errorLogCollectionName)
	err := collection.Find(ctx, bson.M{"$or": []bson.M{
		{"user_uuid": bson.M{"$regex": query}},
		{"username": bson.M{"$regex": query}},
		{"ip_address": bson.M{"$regex": query}},
		{"request_url": bson.M{"$regex": query}},
		{"error_code": bson.M{"$regex": query}},
		{"error_msg": bson.M{"$regex": query}},
		{"stack": bson.M{"$regex": query}},
	}}).Skip(offset).Limit(limit).All(&errorLogList)
	if err != nil {
		e.Core.Zap.Logger.Error(
			"ErrorLogDaoImpl.GetErrorLogListByFuzzyQuery",
			zap.Field{Key: "query", Type: zapcore.StringType, String: query},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		e.Core.Zap.Logger.Info(
			"ErrorLogDaoImpl.GetErrorLogListByFuzzyQuery",
			zap.Field{Key: "query", Type: zapcore.StringType, String: query},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return errorLogList, nil
	}
}

func (e *ErrorLogDaoImpl) InsertErrorLog(errorLog *models.ErrorLogModel, ctx context.Context) error {
	collection := e.Core.Mongo.MongoDatabase.Collection(errorLogCollectionName)
	result, err := collection.InsertOne(ctx, errorLog)
	if err != nil {
		e.Core.Zap.Logger.Error(
			"ErrorLogDaoImpl.InsertErrorLog",
			zap.Field{Key: "errorLog", Type: zapcore.ObjectMarshalerType, Interface: errorLog},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
	} else {
		e.Core.Zap.Logger.Info(
			"ErrorLogDaoImpl.InsertErrorLog",
			zap.Field{Key: "errorLog", Type: zapcore.ObjectMarshalerType, Interface: errorLog},
			zap.Field{Key: "result", Type: zapcore.ObjectMarshalerType, Interface: result},
		)
	}
	return err
}

func (e *ErrorLogDaoImpl) DeleteErrorLog(errorLog *models.ErrorLogModel, ctx context.Context) error {
	collection := e.Core.Mongo.MongoDatabase.Collection(errorLogCollectionName)
	err := collection.RemoveId(ctx, errorLog.ErrorLogID)
	if err != nil {
		e.Core.Zap.Logger.Error(
			"ErrorLogDaoImpl.DeleteErrorLog",
			zap.Field{Key: "errorLog", Type: zapcore.ObjectMarshalerType, Interface: errorLog},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
	} else {
		e.Core.Zap.Logger.Info(
			"ErrorLogDaoImpl.DeleteErrorLog",
			zap.Field{Key: "errorLog", Type: zapcore.ObjectMarshalerType, Interface: errorLog},
		)
	}
	return err
}
