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

const loginLogCollectionName = "login_log"

type LoginLogDao interface {
	GetLoginLogById(loginLogID primitive.ObjectID, ctx context.Context) (*models.LoginLogModel, error)
	GetLoginLogList(offset, limit int64, ctx context.Context) ([]models.LoginLogModel, error)
	GetLoginLogListByCreatedTime(
		startTime, endTime time.Time, offset, limit int64, ctx context.Context,
	) ([]models.LoginLogModel, error)
	GetLoginLogListByUserID(
		userID primitive.ObjectID, offset, limit int64, ctx context.Context,
	) ([]models.LoginLogModel, error)
	GetLoginLogListByIPAddress(ipAddress string, offset, limit int64, ctx context.Context) (
		[]models.LoginLogModel, error,
	)
	InsertLoginLog(
		UserID primitive.ObjectID, Username, Email, IPAddress, UserAgent string, ctx context.Context,
	) (primitive.ObjectID, error)
	DeleteLoginLog(LoginLogID primitive.ObjectID, ctx context.Context) error
}

type LoginLogDaoImpl struct{ *dal.Dao }

func NewLoginLogDao(dao *dal.Dao) LoginLogDao {
	var _ LoginLogDao = (*LoginLogDaoImpl)(nil) // Ensure that the interface is implemented
	return &LoginLogDaoImpl{dao}
}

func (l *LoginLogDaoImpl) GetLoginLogById(loginLogID primitive.ObjectID, ctx context.Context) (
	*models.LoginLogModel, error,
) {
	var loginLog models.LoginLogModel
	collection := l.Dao.Mongo.MongoDatabase.Collection(loginLogCollectionName)
	err := collection.Find(ctx, bson.M{"_id": loginLogID}).One(&loginLog)
	if err != nil {
		l.Dao.Zap.Logger.Error(
			"LoginLogDaoImpl.GetLoginLogById",
			zap.Field{Key: "loginLogID", Type: zapcore.StringType, String: loginLogID.Hex()},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		l.Dao.Zap.Logger.Info(
			"LoginLogDaoImpl.GetLoginLogById",
			zap.Field{Key: "loginLogID", Type: zapcore.StringType, String: loginLogID.Hex()},
		)
		return &loginLog, nil
	}
}
func (l *LoginLogDaoImpl) GetLoginLogList(offset, limit int64, ctx context.Context) ([]models.LoginLogModel, error) {
	var loginLogList []models.LoginLogModel
	collection := l.Dao.Mongo.MongoDatabase.Collection(loginLogCollectionName)
	err := collection.Find(ctx, bson.M{}).Skip(offset).Limit(limit).All(&loginLogList)
	if err != nil {
		l.Dao.Zap.Logger.Error(
			"LoginLogDaoImpl.GetLoginLogList",
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		l.Dao.Zap.Logger.Info(
			"LoginLogDaoImpl.GetLoginLogList",
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return loginLogList, nil
	}
}
func (l *LoginLogDaoImpl) GetLoginLogListByCreatedTime(
	startTime, endTime time.Time, offset, limit int64, ctx context.Context,
) ([]models.LoginLogModel, error) {
	var loginLogList []models.LoginLogModel
	collection := l.Dao.Mongo.MongoDatabase.Collection(loginLogCollectionName)
	err := collection.Find(
		ctx, bson.M{"created_at": bson.M{"$gte": startTime, "$lte": endTime}},
	).Skip(offset).Limit(limit).All(&loginLogList)
	if err != nil {
		l.Dao.Zap.Logger.Error(
			"LoginLogDaoImpl.GetLoginLogListByCreatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime.Format(time.RFC3339)},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime.Format(time.RFC3339)},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		l.Dao.Zap.Logger.Info(
			"LoginLogDaoImpl.GetLoginLogListByCreatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime.Format(time.RFC3339)},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime.Format(time.RFC3339)},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return loginLogList, nil
	}
}
func (l *LoginLogDaoImpl) GetLoginLogListByUserID(
	userID primitive.ObjectID, offset, limit int64, ctx context.Context,
) ([]models.LoginLogModel, error) {
	var loginLogList []models.LoginLogModel
	collection := l.Dao.Mongo.MongoDatabase.Collection(loginLogCollectionName)
	err := collection.Find(ctx, bson.M{"user_id": userID}).Skip(offset).Limit(limit).All(&loginLogList)
	if err != nil {
		l.Dao.Zap.Logger.Error(
			"LoginLogDaoImpl.GetLoginLogListByUserID",
			zap.Field{Key: "userID", Type: zapcore.StringType, String: userID.Hex()},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		l.Dao.Zap.Logger.Info(
			"LoginLogDaoImpl.GetLoginLogListByUserID",
			zap.Field{Key: "userID", Type: zapcore.StringType, String: userID.Hex()},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return loginLogList, nil
	}
}
func (l *LoginLogDaoImpl) GetLoginLogListByIPAddress(
	ipAddress string, offset, limit int64, ctx context.Context,
) ([]models.LoginLogModel, error) {
	var loginLogList []models.LoginLogModel
	collection := l.Dao.Mongo.MongoDatabase.Collection(loginLogCollectionName)
	err := collection.Find(ctx, bson.M{"ip_address": ipAddress}).Skip(offset).Limit(limit).All(&loginLogList)
	if err != nil {
		l.Dao.Zap.Logger.Error(
			"LoginLogDaoImpl.GetLoginLogListByIPAddress",
			zap.Field{Key: "ipAddress", Type: zapcore.StringType, String: ipAddress},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		l.Dao.Zap.Logger.Info(
			"LoginLogDaoImpl.GetLoginLogListByIPAddress",
			zap.Field{Key: "ipAddress", Type: zapcore.StringType, String: ipAddress},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return loginLogList, nil
	}
}
func (l *LoginLogDaoImpl) InsertLoginLog(
	UserID primitive.ObjectID, Username, Email, IPAddress, UserAgent string, ctx context.Context,
) (primitive.ObjectID, error) {
	loginLog := models.LoginLogModel{
		UserID:    UserID,
		Username:  Username,
		Email:     Email,
		IPAddress: IPAddress,
		UserAgent: UserAgent,
		CreatedAt: time.Now(),
	}
	collection := l.Dao.Mongo.MongoDatabase.Collection(loginLogCollectionName)
	result, err := collection.InsertOne(ctx, loginLog)
	if err != nil {
		l.Dao.Zap.Logger.Error(
			"LoginLogDaoImpl.InsertLoginLog",
			zap.Field{Key: "loginLog", Type: zapcore.ObjectMarshalerType, Interface: loginLog},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
	} else {
		l.Dao.Zap.Logger.Info(
			"LoginLogDaoImpl.InsertLoginLog",
			zap.Field{Key: "loginLog", Type: zapcore.ObjectMarshalerType, Interface: loginLog},
		)
	}
	return result.InsertedID.(primitive.ObjectID), err
}

func (l *LoginLogDaoImpl) DeleteLoginLog(loginLogID primitive.ObjectID, ctx context.Context) error {
	collection := l.Dao.Mongo.MongoDatabase.Collection(loginLogCollectionName)
	err := collection.RemoveId(ctx, loginLogID)
	if err != nil {
		l.Dao.Zap.Logger.Error(
			"LoginLogDaoImpl.DeleteLoginLog",
			zap.Field{Key: "loginLogID", Type: zapcore.StringType, String: loginLogID.Hex()},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
	} else {
		l.Dao.Zap.Logger.Info(
			"LoginLogDaoImpl.DeleteLoginLog",
			zap.Field{Key: "loginLogID", Type: zapcore.StringType, String: loginLogID.Hex()},
		)
	}
	return err
}
