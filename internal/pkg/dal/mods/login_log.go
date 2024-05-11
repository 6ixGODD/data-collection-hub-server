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
)

const loginLogCollectionName = "login_log"

type LoginLogDao interface {
	GetLoginLogById(ctx context.Context, loginLogID primitive.ObjectID) (*models.LoginLogModel, error)
	GetLoginLogList(
		ctx context.Context,
		offset, limit int64, desc bool, startTime, endTime *time.Time, userID *primitive.ObjectID,
		ipAddress, userAgent, query *string,
	) ([]models.LoginLogModel, *int64, error)
	InsertLoginLog(
		ctx context.Context,
		UserID primitive.ObjectID, Username, Email, IPAddress, UserAgent string,
	) (primitive.ObjectID, error)
	DeleteLoginLog(LoginLogID primitive.ObjectID, ctx context.Context) error
	DeleteLoginLogList(
		ctx context.Context, startTime, endTime *time.Time, userID *primitive.ObjectID,
		ipAddress, userAgent *string,
	) (*int64, error)
}

type LoginLogDaoImpl struct{ *dal.Dao }

func NewLoginLogDao(dao *dal.Dao) LoginLogDao {
	var _ LoginLogDao = (*LoginLogDaoImpl)(nil) // Ensure that the interface is implemented
	return &LoginLogDaoImpl{dao}
}

func (l *LoginLogDaoImpl) GetLoginLogById(
	ctx context.Context, loginLogID primitive.ObjectID,
) (*models.LoginLogModel, error) {
	collection := l.Dao.Mongo.MongoDatabase.Collection(loginLogCollectionName)
	var loginLog models.LoginLogModel
	err := collection.Find(ctx, bson.M{"_id": loginLogID}).One(&loginLog)
	if err != nil {
		l.Dao.Zap.Logger.Error(
			"LoginLogDaoImpl.GetLoginLogById",
			zap.String("loginLogID", loginLogID.Hex()),
			zap.Error(err),
		)
		return nil, err
	} else {
		l.Dao.Zap.Logger.Info(
			"LoginLogDaoImpl.GetLoginLogById",
			zap.String("loginLogID", loginLogID.Hex()),
		)
		return &loginLog, nil
	}
}

func (l *LoginLogDaoImpl) GetLoginLogList(
	ctx context.Context,
	offset, limit int64, desc bool, startTime, endTime *time.Time, userID *primitive.ObjectID,
	ipAddress, userAgent, query *string,
) ([]models.LoginLogModel, *int64, error) {
	collection := l.Dao.Mongo.MongoDatabase.Collection(loginLogCollectionName)
	var loginLogList []models.LoginLogModel
	var err error
	doc := bson.M{}
	if startTime != nil && endTime != nil {
		doc["created_at"] = bson.M{"$gte": startTime, "$lte": endTime}
	}
	if userID != nil {
		doc["user_id"] = *userID
	}
	if ipAddress != nil {
		doc["ip_address"] = *ipAddress
	}
	if userAgent != nil {
		doc["user_agent"] = *userAgent
	}
	if query != nil {
		doc["$or"] = []bson.M{
			{"username": bson.M{"$regex": *query, "$options": "i"}},
			{"email": bson.M{"$regex": *query, "$options": "i"}},
			{"ip_address": bson.M{"$regex": *query, "$options": "i"}},
			{"user_agent": bson.M{"$regex": *query, "$options": "i"}},
		}
	}
	docJSON, _ := json.Marshal(doc)
	if desc {
		err = collection.Find(ctx, doc).Sort("-created_at").Skip(offset).Limit(limit).All(&loginLogList)
	} else {
		err = collection.Find(ctx, doc).Skip(offset).Limit(limit).All(&loginLogList)
	}
	if err != nil {
		l.Dao.Zap.Logger.Error(
			"LoginLogDaoImpl.GetLoginLogList",
			zap.Int64("offset", offset), zap.Int64("limit", limit),
			zap.ByteString(loginLogCollectionName, docJSON), zap.Error(err), zap.Error(err),
		)
		return nil, nil, err
	}
	count, err := collection.Find(ctx, doc).Count()
	if err != nil {
		l.Dao.Zap.Logger.Error(
			"LoginLogDaoImpl.GetLoginLogList",
			zap.Int64("offset", offset), zap.Int64("limit", limit),
			zap.ByteString(loginLogCollectionName, docJSON), zap.Error(err),
		)
		return nil, nil, err
	}
	l.Dao.Zap.Logger.Info(
		"LoginLogDaoImpl.GetLoginLogList",
		zap.Int64("offset", offset), zap.Int64("limit", limit),
		zap.ByteString(loginLogCollectionName, docJSON), zap.Int64("count", count),
	)
	return loginLogList, &count, nil
}

func (l *LoginLogDaoImpl) InsertLoginLog(
	ctx context.Context, UserID primitive.ObjectID, Username, Email, IPAddress, UserAgent string,
) (primitive.ObjectID, error) {
	collection := l.Dao.Mongo.MongoDatabase.Collection(loginLogCollectionName)
	doc := bson.M{
		"user_id":    UserID,
		"username":   Username,
		"email":      Email,
		"ip_address": IPAddress,
		"user_agent": UserAgent,
		"created_at": time.Now(),
	}
	docJSON, _ := json.Marshal(doc)
	result, err := collection.InsertOne(ctx, doc)
	if err != nil {
		l.Dao.Zap.Logger.Error(
			"LoginLogDaoImpl.InsertLoginLog",
			zap.ByteString(loginLogCollectionName, docJSON),
			zap.Error(err),
		)
	} else {
		l.Dao.Zap.Logger.Info(
			"LoginLogDaoImpl.InsertLoginLog",
			zap.ByteString(loginLogCollectionName, docJSON),
			zap.String("loginLogID", result.InsertedID.(primitive.ObjectID).Hex()),
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
			zap.String("loginLogID", loginLogID.Hex()),
			zap.Error(err),
		)
	} else {
		l.Dao.Zap.Logger.Info(
			"LoginLogDaoImpl.DeleteLoginLog",
			zap.String("loginLogID", loginLogID.Hex()),
		)
	}
	return err
}

func (l *LoginLogDaoImpl) DeleteLoginLogList(
	ctx context.Context, startTime, endTime *time.Time, userID *primitive.ObjectID,
	ipAddress, userAgent *string,
) (*int64, error) {
	collection := l.Dao.Mongo.MongoDatabase.Collection(loginLogCollectionName)
	doc := bson.M{}
	if startTime != nil && endTime != nil {
		doc["created_at"] = bson.M{"$gte": startTime, "$lte": endTime}
	}
	if userID != nil {
		doc["user_id"] = *userID
	}
	if ipAddress != nil {
		doc["ip_address"] = *ipAddress
	}
	if userAgent != nil {
		doc["user_agent"] = *userAgent
	}
	docJSON, _ := json.Marshal(doc)
	result, err := collection.RemoveAll(ctx, doc)
	if err != nil {
		l.Dao.Zap.Logger.Error(
			"LoginLogDaoImpl.DeleteLoginLogList",
			zap.ByteString(loginLogCollectionName, docJSON),
			zap.Error(err),
		)
	} else {
		l.Dao.Zap.Logger.Info(
			"LoginLogDaoImpl.DeleteLoginLogList",
			zap.ByteString(loginLogCollectionName, docJSON),
			zap.Int64("count", result.DeletedCount),
		)
	}
	return &result.DeletedCount, err
}
