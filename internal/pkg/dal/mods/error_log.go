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

const errorLogCollectionName = "error_log"

type ErrorLogDao interface {
	GetErrorLogById(errorLogID primitive.ObjectID, ctx context.Context) (*models.ErrorLogModel, error)
	GetErrorLogList(offset, limit int64, desc bool, ctx context.Context) ([]models.ErrorLogModel, error)
	GetErrorLogListByCreatedTime(
		startTime, endTime time.Time, offset, limit int64, desc bool, ctx context.Context,
	) ([]models.ErrorLogModel, error)
	GetErrorLogListByUserID(
		userID primitive.ObjectID, offset, limit int64, desc bool, ctx context.Context,
	) ([]models.ErrorLogModel, error)
	GetErrorLogListByIPAddress(ipAddress string, offset, limit int64, ctx context.Context) (
		[]models.ErrorLogModel, error,
	)
	GetErrorLogListByRequestURL(requestURL string, offset, limit int64, ctx context.Context) (
		[]models.ErrorLogModel, error,
	)
	GetErrorLogListByErrorCode(errorCode string, offset, limit int64, ctx context.Context) (
		[]models.ErrorLogModel, error,
	)
	GetErrorLogListByFuzzyQuery(query string, offset, limit int64, ctx context.Context) ([]models.ErrorLogModel, error)
	InsertErrorLog(
		userID primitive.ObjectID,
		Username, IPAddress, UserAgent, RequestURL, RequestMethod, RequestPayload, ErrorCode, ErrorMsg, Stack string,
		ctx context.Context,
	) (primitive.ObjectID, error)
	DeleteErrorLog(errorLogID primitive.ObjectID, ctx context.Context) error
}

type ErrorLogDaoImpl struct{ *dal.Dao }

func NewErrorLogDao(dao *dal.Dao) ErrorLogDao {
	var _ ErrorLogDao = (*ErrorLogDaoImpl)(nil) // Ensure that the interface is implemented
	return &ErrorLogDaoImpl{dao}
}

func (e *ErrorLogDaoImpl) GetErrorLogById(errorLogID primitive.ObjectID, ctx context.Context) (
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

func (e *ErrorLogDaoImpl) GetErrorLogList(offset, limit int64, desc bool, ctx context.Context) (
	[]models.ErrorLogModel, error,
) {
	var errorLogList []models.ErrorLogModel
	var err error
	collection := e.Dao.Mongo.MongoDatabase.Collection(errorLogCollectionName)
	if desc {
		err = collection.Find(ctx, bson.M{}).Sort("-created_at").Skip(offset).Limit(limit).All(&errorLogList)
	} else {
		err = collection.Find(ctx, bson.M{}).Skip(offset).Limit(limit).All(&errorLogList)
	}
	if err != nil {
		e.Dao.Zap.Logger.Error(
			"ErrorLogDaoImpl.GetErrorLogList",
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		e.Dao.Zap.Logger.Info(
			"ErrorLogDaoImpl.GetErrorLogList",
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Integer: limit},
		)
		return errorLogList, nil
	}
}

func (e *ErrorLogDaoImpl) GetErrorLogListByCreatedTime(
	startTime, endTime time.Time, offset, limit int64, desc bool, ctx context.Context,
) ([]models.ErrorLogModel, error) {
	var errorLogList []models.ErrorLogModel
	var err error
	collection := e.Dao.Mongo.MongoDatabase.Collection(errorLogCollectionName)
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
		e.Dao.Zap.Logger.Error(
			"ErrorLogDaoImpl.GetErrorLogListByCreatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime.Format(time.RFC3339)},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime.Format(time.RFC3339)},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		e.Dao.Zap.Logger.Info(
			"ErrorLogDaoImpl.GetErrorLogListByCreatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime.Format(time.RFC3339)},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime.Format(time.RFC3339)},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Integer: limit},
		)
		return errorLogList, nil
	}
}

func (e *ErrorLogDaoImpl) GetErrorLogListByUserID(
	userID primitive.ObjectID, offset, limit int64, desc bool, ctx context.Context,
) ([]models.ErrorLogModel, error) {
	var errorLogList []models.ErrorLogModel
	var err error
	collection := e.Dao.Mongo.MongoDatabase.Collection(errorLogCollectionName)
	if desc {
		err = collection.Find(
			ctx, bson.M{"user_id": userID},
		).Sort("-created_at").Skip(offset).Limit(limit).All(&errorLogList)
	} else {
		err = collection.Find(ctx, bson.M{"user_id": userID}).Skip(offset).Limit(limit).All(&errorLogList)
	}
	if err != nil {
		e.Dao.Zap.Logger.Error(
			"ErrorLogDaoImpl.GetErrorLogListByUserID",
			zap.Field{Key: "userID", Type: zapcore.StringType, String: userID.Hex()},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		e.Dao.Zap.Logger.Info(
			"ErrorLogDaoImpl.GetErrorLogListByUserID",
			zap.Field{Key: "userID", Type: zapcore.StringType, String: userID.Hex()},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "desc", Type: zapcore.BoolType, Integer: limit},
		)
		return errorLogList, nil
	}
}

func (e *ErrorLogDaoImpl) GetErrorLogListByIPAddress(
	ipAddress string, offset, limit int64, ctx context.Context,
) ([]models.ErrorLogModel, error) {
	var errorLogList []models.ErrorLogModel
	collection := e.Dao.Mongo.MongoDatabase.Collection(errorLogCollectionName)
	err := collection.Find(ctx, bson.M{"ip_address": ipAddress}).Skip(offset).Limit(limit).All(&errorLogList)
	if err != nil {
		e.Dao.Zap.Logger.Error(
			"ErrorLogDaoImpl.GetErrorLogListByIPAddress",
			zap.Field{Key: "ipAddress", Type: zapcore.StringType, String: ipAddress},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		e.Dao.Zap.Logger.Info(
			"ErrorLogDaoImpl.GetErrorLogListByIPAddress",
			zap.Field{Key: "ipAddress", Type: zapcore.StringType, String: ipAddress},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return errorLogList, nil
	}
}

func (e *ErrorLogDaoImpl) GetErrorLogListByRequestURL(
	requestURL string, offset, limit int64, ctx context.Context,
) ([]models.ErrorLogModel, error) {
	var errorLogList []models.ErrorLogModel
	collection := e.Dao.Mongo.MongoDatabase.Collection(errorLogCollectionName)
	err := collection.Find(ctx, bson.M{"request_url": requestURL}).Skip(offset).Limit(limit).All(&errorLogList)
	if err != nil {
		e.Dao.Zap.Logger.Error(
			"ErrorLogDaoImpl.GetErrorLogListByRequestURL",
			zap.Field{Key: "requestURL", Type: zapcore.StringType, String: requestURL},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		e.Dao.Zap.Logger.Info(
			"ErrorLogDaoImpl.GetErrorLogListByRequestURL",
			zap.Field{Key: "requestURL", Type: zapcore.StringType, String: requestURL},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return errorLogList, nil
	}
}

func (e *ErrorLogDaoImpl) GetErrorLogListByErrorCode(
	errorCode string, offset, limit int64, ctx context.Context,
) ([]models.ErrorLogModel, error) {
	var errorLogList []models.ErrorLogModel
	collection := e.Dao.Mongo.MongoDatabase.Collection(errorLogCollectionName)
	err := collection.Find(ctx, bson.M{"error_code": errorCode}).Skip(offset).Limit(limit).All(&errorLogList)
	if err != nil {
		e.Dao.Zap.Logger.Error(
			"ErrorLogDaoImpl.GetErrorLogListByErrorCode",
			zap.Field{Key: "errorCode", Type: zapcore.StringType, String: errorCode},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		e.Dao.Zap.Logger.Info(
			"ErrorLogDaoImpl.GetErrorLogListByErrorCode",
			zap.Field{Key: "errorCode", Type: zapcore.StringType, String: errorCode},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return errorLogList, nil
	}
}

func (e *ErrorLogDaoImpl) GetErrorLogListByFuzzyQuery(
	query string, offset, limit int64, ctx context.Context,
) ([]models.ErrorLogModel, error) {
	var errorLogList []models.ErrorLogModel
	collection := e.Dao.Mongo.MongoDatabase.Collection(errorLogCollectionName)
	err := collection.Find(
		ctx, bson.M{
			"$or": []bson.M{
				{"user_id": bson.M{"$regex": query}},
				{"username": bson.M{"$regex": query}},
				{"ip_address": bson.M{"$regex": query}},
				{"request_url": bson.M{"$regex": query}},
				{"error_code": bson.M{"$regex": query}},
				{"error_msg": bson.M{"$regex": query}},
				{"stack": bson.M{"$regex": query}},
			},
		},
	).Skip(offset).Limit(limit).All(&errorLogList)
	if err != nil {
		e.Dao.Zap.Logger.Error(
			"ErrorLogDaoImpl.GetErrorLogListByFuzzyQuery",
			zap.Field{Key: "query", Type: zapcore.StringType, String: query},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		e.Dao.Zap.Logger.Info(
			"ErrorLogDaoImpl.GetErrorLogListByFuzzyQuery",
			zap.Field{Key: "query", Type: zapcore.StringType, String: query},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return errorLogList, nil
	}
}

func (e *ErrorLogDaoImpl) InsertErrorLog(
	UserID primitive.ObjectID,
	Username, IPAddress, UserAgent, RequestURL, RequestMethod, RequestPayload, ErrorCode, ErrorMsg, Stack string,
	ctx context.Context,
) (primitive.ObjectID, error) {
	collection := e.Dao.Mongo.MongoDatabase.Collection(errorLogCollectionName)
	errorLog := models.ErrorLogModel{
		UserID:         UserID,
		Username:       Username,
		IPAddress:      IPAddress,
		UserAgent:      UserAgent,
		RequestURL:     RequestURL,
		RequestMethod:  RequestMethod,
		RequestPayload: RequestPayload,
		ErrorCode:      ErrorCode,
		ErrorMsg:       ErrorMsg,
		Stack:          Stack,
		CreatedAt:      time.Now(),
	}
	result, err := collection.InsertOne(ctx, errorLog)
	if err != nil {
		e.Dao.Zap.Logger.Error(
			"ErrorLogDaoImpl.InsertErrorLog",
			zap.Field{Key: "errorLog", Type: zapcore.ObjectMarshalerType, Interface: errorLog},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
	} else {
		e.Dao.Zap.Logger.Info(
			"ErrorLogDaoImpl.InsertErrorLog",
			zap.Field{
				Key: "errorLogID", Type: zapcore.StringType, String: result.InsertedID.(primitive.ObjectID).Hex(),
			},
			zap.Field{Key: "errorLog", Type: zapcore.ObjectMarshalerType, Interface: errorLog},
		)
	}
	return result.InsertedID.(primitive.ObjectID), err
}

func (e *ErrorLogDaoImpl) DeleteErrorLog(errorLogID primitive.ObjectID, ctx context.Context) error {
	collection := e.Dao.Mongo.MongoDatabase.Collection(errorLogCollectionName)
	err := collection.RemoveId(ctx, errorLogID)
	if err != nil {
		e.Dao.Zap.Logger.Error(
			"ErrorLogDaoImpl.DeleteErrorLog",
			zap.Field{Key: "errorLogID", Type: zapcore.StringType, String: errorLogID.Hex()},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
	} else {
		e.Dao.Zap.Logger.Info(
			"ErrorLogDaoImpl.DeleteErrorLog",
			zap.Field{Key: "errorLogID", Type: zapcore.StringType, String: errorLogID.Hex()},
		)
	}
	return err
}
