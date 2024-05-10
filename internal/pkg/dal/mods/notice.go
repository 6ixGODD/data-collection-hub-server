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

const noticeCollectionName = "notice"

type NoticeDao interface {
	GetNoticeById(noticeID primitive.ObjectID, ctx context.Context) (*models.NoticeModel, error)
	GetNoticeList(offset, limit int64, ctx context.Context) ([]models.NoticeModel, error)
	GetNoticeListByNoticeType(noticeType string, offset, limit int64, ctx context.Context) ([]models.NoticeModel, error)
	GetNoticeListByNoticeTypeAndCreatedTime(
		noticeType string, startTime, endTime time.Time, offset, limit int64, ctx context.Context,
	) ([]models.NoticeModel, error)
	GetNoticeListByNoticeTypeAndUpdatedTime(
		noticeType string, startTime, endTime time.Time, offset, limit int64, ctx context.Context,
	) ([]models.NoticeModel, error)
	GetNoticeListByCreatedTime(
		startTime, endTime time.Time, offset, limit int64, ctx context.Context,
	) ([]models.NoticeModel, error)
	GetNoticeListByUpdatedTime(
		startTime, endTime time.Time, offset, limit int64, ctx context.Context,
	) ([]models.NoticeModel, error)
	InsertNotice(title, content, noticeType string, ctx context.Context) (primitive.ObjectID, error)
	UpdateNotice(noticeID primitive.ObjectID, title, content, noticeType string, ctx context.Context) error
	DeleteNotice(noticeID primitive.ObjectID, ctx context.Context) error
}

type NoticeDaoImpl struct{ *dal.Dao }

func NewNoticeDao(dao *dal.Dao) NoticeDao {
	var _ NoticeDao = (*NoticeDaoImpl)(nil) // Ensure that the interface is implemented
	return &NoticeDaoImpl{dao}
}

func (n NoticeDaoImpl) GetNoticeById(noticeID primitive.ObjectID, ctx context.Context) (*models.NoticeModel, error) {
	var notice models.NoticeModel
	collection := n.Dao.Mongo.MongoDatabase.Collection(noticeCollectionName)
	err := collection.Find(ctx, bson.M{"_id": noticeID}).One(&notice)
	if err != nil {
		n.Dao.Zap.Logger.Error(
			"NoticeDaoImpl.GetNoticeById",
			zap.Field{Key: "noticeID", Type: zapcore.StringType, String: noticeID.Hex()},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		n.Dao.Zap.Logger.Info(
			"NoticeDaoImpl.GetNoticeById",
			zap.Field{Key: "noticeID", Type: zapcore.StringType, String: noticeID.Hex()},
		)
		return &notice, nil
	}
}

func (n NoticeDaoImpl) GetNoticeList(offset, limit int64, ctx context.Context) ([]models.NoticeModel, error) {
	var noticeList []models.NoticeModel
	collection := n.Dao.Mongo.MongoDatabase.Collection(noticeCollectionName)
	err := collection.Find(ctx, bson.M{}).Skip(offset).Limit(limit).All(&noticeList)
	if err != nil {
		n.Dao.Zap.Logger.Error(
			"NoticeDaoImpl.GetNoticeList",
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		n.Dao.Zap.Logger.Info(
			"NoticeDaoImpl.GetNoticeList",
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return noticeList, nil
	}
}

func (n NoticeDaoImpl) GetNoticeListByNoticeType(
	noticeType string, offset, limit int64, ctx context.Context,
) ([]models.NoticeModel, error) {
	var noticeList []models.NoticeModel
	collection := n.Dao.Mongo.MongoDatabase.Collection(noticeCollectionName)
	err := collection.Find(ctx, bson.M{"notice_type": noticeType}).Skip(offset).Limit(limit).All(&noticeList)
	if err != nil {
		n.Dao.Zap.Logger.Error(
			"NoticeDaoImpl.GetNoticeListByNoticeType",
			zap.Field{Key: "noticeType", Type: zapcore.StringType, String: noticeType},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		n.Dao.Zap.Logger.Info(
			"NoticeDaoImpl.GetNoticeListByNoticeType",
			zap.Field{Key: "noticeType", Type: zapcore.StringType, String: noticeType},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return noticeList, nil
	}
}

func (n NoticeDaoImpl) GetNoticeListByNoticeTypeAndCreatedTime(
	noticeType string, startTime, endTime time.Time, offset, limit int64, ctx context.Context,
) ([]models.NoticeModel, error) {
	var noticeList []models.NoticeModel
	collection := n.Dao.Mongo.MongoDatabase.Collection(noticeCollectionName)
	err := collection.Find(
		ctx, bson.M{"notice_type": noticeType, "created_time": bson.M{"$gte": startTime, "$lte": endTime}},
	).Skip(offset).Limit(limit).All(&noticeList)
	if err != nil {
		n.Dao.Zap.Logger.Error(
			"NoticeDaoImpl.GetNoticeListByNoticeTypeAndCreatedTime",
			zap.Field{Key: "noticeType", Type: zapcore.StringType, String: noticeType},
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime.Format(time.RFC3339)},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime.Format(time.RFC3339)},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		n.Dao.Zap.Logger.Info(
			"NoticeDaoImpl.GetNoticeListByNoticeTypeAndCreatedTime",
			zap.Field{Key: "noticeType", Type: zapcore.StringType, String: noticeType},
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime.Format(time.RFC3339)},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime.Format(time.RFC3339)},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return noticeList, nil
	}
}

func (n NoticeDaoImpl) GetNoticeListByNoticeTypeAndUpdatedTime(
	noticeType string, startTime, endTime time.Time, offset, limit int64, ctx context.Context,
) ([]models.NoticeModel, error) {
	var noticeList []models.NoticeModel
	collection := n.Dao.Mongo.MongoDatabase.Collection(noticeCollectionName)
	err := collection.Find(
		ctx, bson.M{"notice_type": noticeType, "updated_time": bson.M{"$gte": startTime, "$lte": endTime}},
	).Skip(offset).Limit(limit).All(&noticeList)
	if err != nil {
		n.Dao.Zap.Logger.Error(
			"NoticeDaoImpl.GetNoticeListByNoticeTypeAndUpdatedTime",
			zap.Field{Key: "noticeType", Type: zapcore.StringType, String: noticeType},
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime.Format(time.RFC3339)},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime.Format(time.RFC3339)},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		n.Dao.Zap.Logger.Info(
			"NoticeDaoImpl.GetNoticeListByNoticeTypeAndUpdatedTime",
			zap.Field{Key: "noticeType", Type: zapcore.StringType, String: noticeType},
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime.Format(time.RFC3339)},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime.Format(time.RFC3339)},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return noticeList, nil
	}
}

func (n NoticeDaoImpl) GetNoticeListByCreatedTime(
	startTime, endTime time.Time, offset, limit int64, ctx context.Context,
) ([]models.NoticeModel, error) {
	var noticeList []models.NoticeModel
	collection := n.Dao.Mongo.MongoDatabase.Collection(noticeCollectionName)
	err := collection.Find(
		ctx, bson.M{"created_time": bson.M{"$gte": startTime, "$lte": endTime}},
	).Skip(offset).Limit(limit).All(&noticeList)
	if err != nil {
		n.Dao.Zap.Logger.Error(
			"NoticeDaoImpl.GetNoticeListByCreatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime.Format(time.RFC3339)},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime.Format(time.RFC3339)},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		n.Dao.Zap.Logger.Info(
			"NoticeDaoImpl.GetNoticeListByCreatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime.Format(time.RFC3339)},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime.Format(time.RFC3339)},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return noticeList, nil
	}
}

func (n NoticeDaoImpl) GetNoticeListByUpdatedTime(
	startTime, endTime time.Time, offset, limit int64, ctx context.Context,
) ([]models.NoticeModel, error) {
	var noticeList []models.NoticeModel
	collection := n.Dao.Mongo.MongoDatabase.Collection(noticeCollectionName)
	err := collection.Find(
		ctx, bson.M{"updated_time": bson.M{"$gte": startTime, "$lte": endTime}},
	).Skip(offset).Limit(limit).All(&noticeList)
	if err != nil {
		n.Dao.Zap.Logger.Error(
			"NoticeDaoImpl.GetNoticeListByUpdatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime.Format(time.RFC3339)},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime.Format(time.RFC3339)},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		n.Dao.Zap.Logger.Info(
			"NoticeDaoImpl.GetNoticeListByUpdatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime.Format(time.RFC3339)},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime.Format(time.RFC3339)},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return noticeList, nil
	}
}

func (n NoticeDaoImpl) InsertNotice(title, content, noticeType string, ctx context.Context) (
	primitive.ObjectID, error,
) {
	notice := models.NoticeModel{
		NoticeID:   primitive.NewObjectID(),
		Title:      title,
		Content:    content,
		NoticeType: noticeType,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	collection := n.Dao.Mongo.MongoDatabase.Collection(noticeCollectionName)
	_, err := collection.InsertOne(ctx, notice)
	if err != nil {
		n.Dao.Zap.Logger.Error(
			"NoticeDaoImpl.InsertNotice",
			zap.Field{Key: "notice", Type: zapcore.ReflectType, Interface: notice},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
	} else {
		n.Dao.Zap.Logger.Info(
			"NoticeDaoImpl.InsertNotice",
			zap.Field{Key: "notice", Type: zapcore.ReflectType, Interface: notice},
		)
	}
	return notice.NoticeID, err
}

func (n NoticeDaoImpl) UpdateNotice(
	noticeID primitive.ObjectID, title, content, noticeType string, ctx context.Context,
) error {
	collection := n.Dao.Mongo.MongoDatabase.Collection(noticeCollectionName)
	err := collection.UpdateOne(
		ctx,
		bson.M{"_id": noticeID},
		bson.M{
			"$set": bson.M{
				"title":       title,
				"content":     content,
				"notice_type": noticeType,
				"updated_at":  time.Now(),
			},
		},
	)
	if err != nil {
		n.Dao.Zap.Logger.Error(
			"NoticeDaoImpl.UpdateNotice",
			zap.Field{Key: "noticeID", Type: zapcore.StringType, String: noticeID.Hex()},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
	} else {
		n.Dao.Zap.Logger.Info(
			"NoticeDaoImpl.UpdateNotice",
			zap.Field{Key: "noticeID", Type: zapcore.StringType, String: noticeID.Hex()},
		)
	}
	return err
}

func (n NoticeDaoImpl) DeleteNotice(noticeID primitive.ObjectID, ctx context.Context) error {
	collection := n.Dao.Mongo.MongoDatabase.Collection(noticeCollectionName)
	err := collection.RemoveId(ctx, noticeID)
	if err != nil {
		n.Dao.Zap.Logger.Error(
			"NoticeDaoImpl.DeleteNotice",
			zap.Field{Key: "noticeID", Type: zapcore.StringType, String: noticeID.Hex()},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
	} else {
		n.Dao.Zap.Logger.Info(
			"NoticeDaoImpl.DeleteNotice",
			zap.Field{Key: "noticeID", Type: zapcore.StringType, String: noticeID.Hex()},
		)
	}
	return err
}
