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

const noticeCollectionName = "notice"

type NoticeDao interface {
	GetNoticeById(ctx context.Context, noticeID primitive.ObjectID) (*models.NoticeModel, error)
	GetNoticeList(
		ctx context.Context,
		offset, limit int64, desc bool, createStartTime, createEndTime, updateStartTime, updateEndTime *time.Time,
		noticeType *string,
	) ([]models.NoticeModel, *int64, error)
	InsertNotice(ctx context.Context, title, content, noticeType string) (primitive.ObjectID, error)
	UpdateNotice(ctx context.Context, noticeID primitive.ObjectID, title, content, noticeType *string) error
	DeleteNotice(ctx context.Context, noticeID primitive.ObjectID) error
	DeleteNoticeList(
		ctx context.Context,
		createStartTime, createEndTime, updateStartTime, updateEndTime *time.Time,
		title, content, noticeType *string,
	) (*int64, error)
}

type NoticeDaoImpl struct{ *dal.Dao }

func NewNoticeDao(dao *dal.Dao) NoticeDao {
	var _ NoticeDao = (*NoticeDaoImpl)(nil) // Ensure that the interface is implemented
	return &NoticeDaoImpl{dao}
}

func (n *NoticeDaoImpl) GetNoticeById(ctx context.Context, noticeID primitive.ObjectID) (*models.NoticeModel, error) {
	collection := n.Dao.Mongo.MongoDatabase.Collection(noticeCollectionName)
	var notice models.NoticeModel
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

func (n *NoticeDaoImpl) GetNoticeList(
	ctx context.Context,
	offset, limit int64, desc bool, createStartTime, createEndTime, updateStartTime, updateEndTime *time.Time,
	noticeType *string,
) ([]models.NoticeModel, *int64, error) {
	collection := n.Dao.Mongo.MongoDatabase.Collection(noticeCollectionName)
	var noticeList []models.NoticeModel
	var err error
	doc := bson.M{}
	if createStartTime != nil && createEndTime != nil {
		doc["created_at"] = bson.M{"$gte": createStartTime, "$lte": createEndTime}
	}
	if updateStartTime != nil && updateEndTime != nil {
		doc["updated_at"] = bson.M{"$gte": updateStartTime, "$lte": updateEndTime}
	}
	if noticeType != nil {
		doc["notice_type"] = *noticeType
	}
	docJSON, _ := json.Marshal(doc)

	if desc {
		err = collection.Find(ctx, doc).Sort("-created_at").Skip(offset).Limit(limit).All(&noticeList)
	} else {
		err = collection.Find(ctx, doc).Skip(offset).Limit(limit).All(&noticeList)
	}
	if err != nil {
		n.Dao.Zap.Logger.Error(
			"NoticeDaoImpl.GetNoticeList",
			zap.Int64("offset", offset), zap.Int64("limit", limit),
			zap.Bool("desc", desc),
			zap.ByteString(noticeCollectionName, docJSON),
			zap.Error(err),
		)
		return nil, nil, err
	}
	count, err := collection.Find(ctx, doc).Count()
	if err != nil {
		n.Dao.Zap.Logger.Error(
			"NoticeDaoImpl.GetNoticeList",
			zap.Int64("offset", offset), zap.Int64("limit", limit),
			zap.Bool("desc", desc),
			zap.ByteString(noticeCollectionName, docJSON),
			zap.Error(err),
		)
		return nil, nil, err
	}
	n.Dao.Zap.Logger.Info(
		"NoticeDaoImpl.GetNoticeList",
		zap.Int64("offset", offset), zap.Int64("limit", limit),
		zap.Bool("desc", desc),
		zap.ByteString(noticeCollectionName, docJSON),
		zap.Int64("count", count),
	)
	return noticeList, &count, nil
}

func (n *NoticeDaoImpl) InsertNotice(
	ctx context.Context, title, content, noticeType string,
) (primitive.ObjectID, error) {
	collection := n.Dao.Mongo.MongoDatabase.Collection(noticeCollectionName)
	doc := bson.M{
		"title":       title,
		"content":     content,
		"notice_type": noticeType,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
	}
	docJSON, _ := json.Marshal(doc)
	result, err := collection.InsertOne(ctx, doc)
	if err != nil {
		n.Dao.Zap.Logger.Error(
			"NoticeDaoImpl.InsertNotice",
			zap.ByteString(noticeCollectionName, docJSON),
			zap.Error(err),
		)
	} else {
		n.Dao.Zap.Logger.Info(
			"NoticeDaoImpl.InsertNotice",
			zap.ByteString(noticeCollectionName, docJSON),
			zap.String("noticeID", result.InsertedID.(primitive.ObjectID).Hex()),
		)
	}
	return result.InsertedID.(primitive.ObjectID), err
}

func (n *NoticeDaoImpl) UpdateNotice(
	ctx context.Context, noticeID primitive.ObjectID, title, content, noticeType *string,
) error {
	collection := n.Dao.Mongo.MongoDatabase.Collection(noticeCollectionName)
	doc := bson.M{"updated_at": time.Now()}
	if title != nil {
		doc["title"] = *title
	}
	if content != nil {
		doc["content"] = *content
	}
	if noticeType != nil {
		doc["notice_type"] = *noticeType
	}
	docJSON, _ := json.Marshal(doc)
	err := collection.UpdateId(ctx, noticeID, bson.M{"$set": doc})

	if err != nil {
		n.Dao.Zap.Logger.Error(
			"NoticeDaoImpl.UpdateNotice",
			zap.String("noticeID", noticeID.Hex()),
			zap.ByteString(noticeCollectionName, docJSON),
			zap.Error(err),
		)
	} else {
		n.Dao.Zap.Logger.Info(
			"NoticeDaoImpl.UpdateNotice",
			zap.String("noticeID", noticeID.Hex()),
			zap.ByteString(noticeCollectionName, docJSON),
		)
	}
	return err
}

func (n *NoticeDaoImpl) DeleteNotice(ctx context.Context, noticeID primitive.ObjectID) error {
	collection := n.Dao.Mongo.MongoDatabase.Collection(noticeCollectionName)
	err := collection.RemoveId(ctx, noticeID)
	if err != nil {
		n.Dao.Zap.Logger.Error(
			"NoticeDaoImpl.DeleteNotice",
			zap.String("noticeID", noticeID.Hex()),
			zap.Error(err),
		)
	} else {
		n.Dao.Zap.Logger.Info(
			"NoticeDaoImpl.DeleteNotice",
			zap.String("noticeID", noticeID.Hex()),
		)
	}
	return err
}

func (n *NoticeDaoImpl) DeleteNoticeList(
	ctx context.Context,
	createStartTime, createEndTime, updateStartTime, updateEndTime *time.Time,
	title, content, noticeType *string,
) (*int64, error) {
	collection := n.Dao.Mongo.MongoDatabase.Collection(noticeCollectionName)
	doc := bson.M{}
	if createStartTime != nil && createEndTime != nil {
		doc["created_at"] = bson.M{"$gte": createStartTime, "$lte": createEndTime}
	}
	if updateStartTime != nil && updateEndTime != nil {
		doc["updated_at"] = bson.M{"$gte": updateStartTime, "$lte": updateEndTime}
	}
	if title != nil {
		doc["title"] = *title
	}
	if content != nil {
		doc["content"] = *content
	}
	if noticeType != nil {
		doc["notice_type"] = *noticeType
	}
	docJSON, _ := json.Marshal(doc)
	result, err := collection.RemoveAll(ctx, doc)
	if err != nil {
		n.Dao.Zap.Logger.Error(
			"NoticeDaoImpl.DeleteNoticeList",
			zap.ByteString(noticeCollectionName, docJSON),
			zap.Error(err),
		)
	} else {
		n.Dao.Zap.Logger.Info(
			"NoticeDaoImpl.DeleteNoticeList",
			zap.ByteString(noticeCollectionName, docJSON),
		)
	}
	return &result.DeletedCount, err
}
