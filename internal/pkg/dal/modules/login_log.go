package modules

import (
	"context"

	"data-collection-hub-server/internal/pkg/dal"
	"data-collection-hub-server/internal/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var loginLogCollectionName = "login_log"

type LoginLogDao interface {
	GetLoginLogById(loginLogId string, ctx context.Context) (*models.LoginLogModel, error)
	GetLoginLogList(offset, limit int64, ctx context.Context) ([]models.LoginLogModel, error)
	GetLoginLogListByCreatedTime(startTime, endTime string, offset, limit int64, ctx context.Context) ([]models.LoginLogModel, error)
	GetLoginLogListByUserUUID(userUUID string, offset, limit int64, ctx context.Context) ([]models.LoginLogModel, error)
	GetLoginLogListByIPAddress(ipAddress string, offset, limit int64, ctx context.Context) ([]models.LoginLogModel, error)
	InsertLoginLog(loginLog *models.LoginLogModel, ctx context.Context) error
	DeleteLoginLog(loginLog *models.LoginLogModel, ctx context.Context) error
}

type LoginLogDaoImpl struct{ *dal.Core }

func NewLoginLogDao(core *dal.Core) LoginLogDao {
	var _ LoginLogDao = (*LoginLogDaoImpl)(nil) // Ensure that the interface is implemented
	return &LoginLogDaoImpl{core}
}

func (l *LoginLogDaoImpl) GetLoginLogById(loginLogId string, ctx context.Context) (*models.LoginLogModel, error) {
	var loginLog models.LoginLogModel
	collection := l.Core.Mongo.MongoDatabase.Collection(loginLogCollectionName)
	err := collection.Find(ctx, bson.M{"_id": loginLogId}).One(&loginLog)
	if err != nil {
		l.Core.Zap.Logger.Error(
			"LoginLogDaoImpl.GetLoginLogById",
			zap.Field{Key: "loginLogId", Type: zapcore.StringType, String: loginLogId},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		l.Core.Zap.Logger.Info(
			"LoginLogDaoImpl.GetLoginLogById",
			zap.Field{Key: "loginLogId", Type: zapcore.StringType, String: loginLogId},
		)
		return &loginLog, nil
	}
}
func (l *LoginLogDaoImpl) GetLoginLogList(offset, limit int64, ctx context.Context) ([]models.LoginLogModel, error) {
	var loginLogList []models.LoginLogModel
	collection := l.Core.Mongo.MongoDatabase.Collection(loginLogCollectionName)
	err := collection.Find(ctx, bson.M{}).Skip(offset).Limit(limit).All(&loginLogList)
	if err != nil {
		l.Core.Zap.Logger.Error(
			"LoginLogDaoImpl.GetLoginLogList",
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		l.Core.Zap.Logger.Info(
			"LoginLogDaoImpl.GetLoginLogList",
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return loginLogList, nil
	}
}
func (l *LoginLogDaoImpl) GetLoginLogListByCreatedTime(startTime, endTime string, offset, limit int64, ctx context.Context) ([]models.LoginLogModel, error) {
	var loginLogList []models.LoginLogModel
	collection := l.Core.Mongo.MongoDatabase.Collection(loginLogCollectionName)
	err := collection.Find(
		ctx, bson.M{"created_at": bson.M{"$gte": startTime, "$lte": endTime}},
	).Skip(offset).Limit(limit).All(&loginLogList)
	if err != nil {
		l.Core.Zap.Logger.Error(
			"LoginLogDaoImpl.GetLoginLogListByCreatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		l.Core.Zap.Logger.Info(
			"LoginLogDaoImpl.GetLoginLogListByCreatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return loginLogList, nil
	}
}
func (l *LoginLogDaoImpl) GetLoginLogListByUserUUID(userUUID string, offset, limit int64, ctx context.Context) ([]models.LoginLogModel, error) {
	var loginLogList []models.LoginLogModel
	collection := l.Core.Mongo.MongoDatabase.Collection(loginLogCollectionName)
	err := collection.Find(ctx, bson.M{"user_uuid": userUUID}).Skip(offset).Limit(limit).All(&loginLogList)
	if err != nil {
		l.Core.Zap.Logger.Error(
			"LoginLogDaoImpl.GetLoginLogListByUserUUID",
			zap.Field{Key: "userUUID", Type: zapcore.StringType, String: userUUID},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		l.Core.Zap.Logger.Info(
			"LoginLogDaoImpl.GetLoginLogListByUserUUID",
			zap.Field{Key: "userUUID", Type: zapcore.StringType, String: userUUID},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return loginLogList, nil
	}
}
func (l *LoginLogDaoImpl) GetLoginLogListByIPAddress(ipAddress string, offset, limit int64, ctx context.Context) ([]models.LoginLogModel, error) {
	var loginLogList []models.LoginLogModel
	collection := l.Core.Mongo.MongoDatabase.Collection(loginLogCollectionName)
	err := collection.Find(ctx, bson.M{"ip_address": ipAddress}).Skip(offset).Limit(limit).All(&loginLogList)
	if err != nil {
		l.Core.Zap.Logger.Error(
			"LoginLogDaoImpl.GetLoginLogListByIPAddress",
			zap.Field{Key: "ipAddress", Type: zapcore.StringType, String: ipAddress},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		l.Core.Zap.Logger.Info(
			"LoginLogDaoImpl.GetLoginLogListByIPAddress",
			zap.Field{Key: "ipAddress", Type: zapcore.StringType, String: ipAddress},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return loginLogList, nil
	}
}
func (l *LoginLogDaoImpl) InsertLoginLog(loginLog *models.LoginLogModel, ctx context.Context) error {
	collection := l.Core.Mongo.MongoDatabase.Collection(loginLogCollectionName)
	result, err := collection.InsertOne(ctx, loginLog)
	if err != nil {
		l.Core.Zap.Logger.Error(
			"LoginLogDaoImpl.InsertLoginLog",
			zap.Field{Key: "loginLog", Type: zapcore.ObjectMarshalerType, Interface: loginLog},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
	} else {
		l.Core.Zap.Logger.Info(
			"LoginLogDaoImpl.InsertLoginLog",
			zap.Field{Key: "loginLog", Type: zapcore.ObjectMarshalerType, Interface: loginLog},
			zap.Field{Key: "result", Type: zapcore.ObjectMarshalerType, Interface: result},
		)
	}
	return err
}
func (l *LoginLogDaoImpl) DeleteLoginLog(loginLog *models.LoginLogModel, ctx context.Context) error {
	collection := l.Core.Mongo.MongoDatabase.Collection(loginLogCollectionName)
	err := collection.RemoveId(ctx, loginLog.LoginLogID)
	if err != nil {
		l.Core.Zap.Logger.Error(
			"LoginLogDaoImpl.DeleteLoginLog",
			zap.Field{Key: "loginLog", Type: zapcore.ObjectMarshalerType, Interface: loginLog},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
	} else {
		l.Core.Zap.Logger.Info(
			"LoginLogDaoImpl.DeleteLoginLog",
			zap.Field{Key: "loginLog", Type: zapcore.ObjectMarshalerType, Interface: loginLog},
		)
	}
	return err
}
