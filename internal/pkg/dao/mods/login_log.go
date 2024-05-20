package mods

import (
	"context"
	"errors"
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
)

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
	CacheLoginLog(
		ctx context.Context, username, IPAddress, UserAgent string,
	) error
	SyncLoginLog(ctx context.Context)
	DeleteLoginLog(LoginLogID primitive.ObjectID, ctx context.Context) error
	DeleteLoginLogList(
		ctx context.Context, startTime, endTime *time.Time, userID *primitive.ObjectID,
		ipAddress, userAgent *string,
	) (*int64, error)
}

type LoginLogDaoImpl struct {
	core    *dao.Core
	cache   *dao.Cache
	userDao UserDao
}

func NewLoginLogDao(ctx context.Context, dao *dao.Core, cache *dao.Cache, userDao UserDao) (LoginLogDao, error) {
	var _ LoginLogDao = (*LoginLogDaoImpl)(nil) // Ensure that the interface is implemented
	coll := dao.Mongo.MongoClient.Database(dao.Mongo.DatabaseName).Collection(config.LoginLogCollectionName)
	err := coll.CreateIndexes(
		ctx, []options.IndexModel{{Key: []string{"created_at"}}, {Key: []string{"user_id"}}},
	)
	if err != nil {
		dao.Logger.Error(
			fmt.Sprintf("Failed to create index for %s", config.LoginLogCollectionName),
			zap.Error(err),
		)
		return nil, err
	}
	return &LoginLogDaoImpl{
		core:    dao,
		userDao: userDao,
		cache:   cache,
	}, nil
}

func (l *LoginLogDaoImpl) GetLoginLogById(
	ctx context.Context, loginLogID primitive.ObjectID,
) (*models.LoginLogModel, error) {
	coll := l.core.Mongo.MongoClient.Database(l.core.Mongo.DatabaseName).Collection(config.LoginLogCollectionName)
	var loginLog models.LoginLogModel
	err := coll.Find(ctx, bson.M{"_id": loginLogID}).One(&loginLog)
	if err != nil {
		l.core.Logger.Error(
			"LoginLogDaoImpl.GetLoginLogById: failed to find login log",
			zap.Error(err), zap.String("loginLogID", loginLogID.Hex()),
		)
		return nil, err
	} else {
		l.core.Logger.Info("LoginLogDaoImpl.GetLoginLogById: success", zap.String("loginLogID", loginLogID.Hex()))
		return &loginLog, nil
	}
}

func (l *LoginLogDaoImpl) GetLoginLogList(
	ctx context.Context,
	offset, limit int64, desc bool, startTime, endTime *time.Time, userID *primitive.ObjectID,
	ipAddress, userAgent, query *string,
) ([]models.LoginLogModel, *int64, error) {
	coll := l.core.Mongo.MongoClient.Database(l.core.Mongo.DatabaseName).Collection(config.LoginLogCollectionName)
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
		err = coll.Find(ctx, doc).Sort("-created_at").Skip(offset).Limit(limit).All(&loginLogList)
	} else {
		err = coll.Find(ctx, doc).Skip(offset).Limit(limit).All(&loginLogList)
	}
	if err != nil {
		l.core.Logger.Error(
			"LoginLogDaoImpl.GetLoginLogList: failed to find login logs",
			zap.Error(err), zap.ByteString(config.LoginLogCollectionName, docJSON),
		)
		return nil, nil, err
	}
	count, err := coll.Find(ctx, doc).Count()
	if err != nil {
		l.core.Logger.Error(
			"LoginLogDaoImpl.GetLoginLogList: failed to count login logs",
			zap.Error(err), zap.ByteString(config.LoginLogCollectionName, docJSON),
		)
		return nil, nil, err
	}
	l.core.Logger.Info(
		"LoginLogDaoImpl.GetLoginLogList: success",
		zap.Int64("count", count), zap.ByteString(config.LoginLogCollectionName, docJSON),
	)
	return loginLogList, &count, nil
}

func (l *LoginLogDaoImpl) InsertLoginLog(
	ctx context.Context, UserID primitive.ObjectID, Username, Email, IPAddress, UserAgent string,
) (primitive.ObjectID, error) {
	coll := l.core.Mongo.MongoClient.Database(l.core.Mongo.DatabaseName).Collection(config.LoginLogCollectionName)
	doc := bson.M{
		"user_id":    UserID,
		"username":   Username,
		"email":      Email,
		"ip_address": IPAddress,
		"user_agent": UserAgent,
		"created_at": time.Now(),
	}
	docJSON, _ := json.Marshal(doc)
	result, err := coll.InsertOne(ctx, doc)
	if err != nil {
		l.core.Logger.Error(
			"LoginLogDaoImpl.InsertLoginLog: failed to insert login log",
			zap.Error(err), zap.ByteString(config.LoginLogCollectionName, docJSON),
		)
	} else {
		l.core.Logger.Info(
			"LoginLogDaoImpl.InsertLoginLog: success",
			zap.String("loginLogID", result.InsertedID.(primitive.ObjectID).Hex()),
			zap.ByteString(config.LoginLogCollectionName, docJSON),
		)
	}
	return result.InsertedID.(primitive.ObjectID), err
}

// CacheLoginLog caches login logs in cache
func (l *LoginLogDaoImpl) CacheLoginLog(
	ctx context.Context, username, IPAddress, UserAgent string,
) error {
	user, err := l.userDao.GetUserByUsername(ctx, username)
	if err != nil {
		l.core.Logger.Error(
			"LoginLogDaoImpl.CacheLoginLog: failed to get user",
			zap.Error(err), zap.String("username", username),
		)
		return err
	}
	loginLog := models.LoginLogCache{
		UserIDHex: user.UserID.Hex(),
		Username:  user.Username,
		Email:     user.Email,
		IPAddress: IPAddress,
		UserAgent: UserAgent,
		CreatedAt: time.Now(),
	}
	loginLogJSON, err := json.Marshal(loginLog)
	if err != nil {
		l.core.Logger.Error(
			"LoginLogDaoImpl.CacheLoginLog: failed to marshal login log",
			zap.Error(err), zap.String("username", username),
		)
		return err
	}
	return l.cache.RightPush(ctx, config.LoginLogCacheKey, string(loginLogJSON))
}

// SyncLoginLog syncs login logs from cache to database
func (l *LoginLogDaoImpl) SyncLoginLog(ctx context.Context) {
	for {
		loginLogJSON, err := l.cache.LeftPop(ctx, config.LoginLogCacheKey)
		if err != nil {
			if errors.Is(err, dao.CacheNil{}) {
				break
			}
			l.core.Logger.Error(
				"LoginLogDaoImpl.SyncLoginLog: failed to pop login log from cache", zap.Error(err),
			)
			return
		}
		var loginLog models.LoginLogCache
		if err := json.Unmarshal([]byte(*loginLogJSON), &loginLog); err != nil {
			l.core.Logger.Error(
				"LoginLogDaoImpl.SyncLoginLog: failed to unmarshal login log",
				zap.Error(err), zap.String("loginLogJSON", *loginLogJSON),
			)
			continue
		}
		userID, err := primitive.ObjectIDFromHex(loginLog.UserIDHex)
		if err != nil {
			l.core.Logger.Error(
				"LoginLogDaoImpl.SyncLoginLog: failed to convert user ID",
				zap.Error(err), zap.String("loginLogJSON", *loginLogJSON),
			)
			continue
		}
		if _, err := l.InsertLoginLog(
			ctx, userID, loginLog.Username, loginLog.Email, loginLog.IPAddress, loginLog.UserAgent,
		); err != nil {
			l.core.Logger.Error(
				"LoginLogDaoImpl.SyncLoginLog: failed to insert login log",
				zap.Error(err), zap.String("loginLogJSON", *loginLogJSON),
			)
		}
	}
}

func (l *LoginLogDaoImpl) DeleteLoginLog(loginLogID primitive.ObjectID, ctx context.Context) error {
	coll := l.core.Mongo.MongoClient.Database(l.core.Mongo.DatabaseName).Collection(config.LoginLogCollectionName)
	err := coll.RemoveId(ctx, loginLogID)
	if err != nil {
		l.core.Logger.Error(
			"LoginLogDaoImpl.DeleteLoginLog: failed to delete login log",
			zap.Error(err), zap.String("loginLogID", loginLogID.Hex()),
		)
	} else {
		l.core.Logger.Info("LoginLogDaoImpl.DeleteLoginLog: success", zap.String("loginLogID", loginLogID.Hex()))
	}
	return err
}

func (l *LoginLogDaoImpl) DeleteLoginLogList(
	ctx context.Context, startTime, endTime *time.Time, userID *primitive.ObjectID,
	ipAddress, userAgent *string,
) (*int64, error) {
	coll := l.core.Mongo.MongoClient.Database(l.core.Mongo.DatabaseName).Collection(config.LoginLogCollectionName)
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
	result, err := coll.RemoveAll(ctx, doc)
	if err != nil {
		l.core.Logger.Error(
			"LoginLogDaoImpl.DeleteLoginLogList: failed to delete login logs",
			zap.Error(err), zap.ByteString(config.LoginLogCollectionName, docJSON),
		)
	} else {
		l.core.Logger.Info(
			"LoginLogDaoImpl.DeleteLoginLogList: success",
			zap.Int64("count", result.DeletedCount), zap.ByteString(config.LoginLogCollectionName, docJSON),
		)
	}
	return &result.DeletedCount, err
}
