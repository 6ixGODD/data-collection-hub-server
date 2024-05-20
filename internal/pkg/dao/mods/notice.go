package mods

import (
	"context"
	"fmt"
	"time"

	"data-collection-hub-server/internal/pkg/config"
	"data-collection-hub-server/internal/pkg/dao"
	"data-collection-hub-server/internal/pkg/models"
	"github.com/goccy/go-json"
	"github.com/qiniu/qmgo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	opt "go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

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

type NoticeDaoImpl struct {
	*dao.Core
	*dao.Cache
}

func NewNoticeDao(ctx context.Context, dao *dao.Core, cache *dao.Cache) (NoticeDao, error) {
	var _ NoticeDao = (*NoticeDaoImpl)(nil)
	collection := dao.Mongo.MongoClient.Database(dao.Mongo.DatabaseName).Collection(config.NoticeCollectionName)
	err := collection.CreateIndexes(
		ctx, []options.IndexModel{
			{
				Key:          []string{"title"},
				IndexOptions: opt.Index().SetUnique(true),
			},
			{Key: []string{"created_at"}}, {Key: []string{"updated_at"}},
		},
	)
	if err != nil {
		dao.Logger.Error(fmt.Sprintf("Failed to create indexes for %s", config.NoticeCollectionName), zap.Error(err))
		return nil, err
	}
	return &NoticeDaoImpl{dao, cache}, nil
}

func (n *NoticeDaoImpl) GetNoticeById(ctx context.Context, noticeID primitive.ObjectID) (*models.NoticeModel, error) {
	collection := n.Core.Mongo.MongoClient.Database(n.Core.Mongo.DatabaseName).Collection(config.NoticeCollectionName)
	var notice models.NoticeModel
	err := collection.Find(ctx, bson.M{"_id": noticeID}).One(&notice)
	if err != nil {
		n.Core.Logger.Error(
			"NoticeDaoImpl.GetNoticeById: failed to find notice",
			zap.Error(err), zap.String("noticeID", noticeID.Hex()),
		)
		return nil, err
	} else {
		n.Core.Logger.Info("NoticeDaoImpl.GetNoticeById: success", zap.String("noticeID", noticeID.Hex()))
		return &notice, nil
	}
}

func (n *NoticeDaoImpl) GetNoticeList(
	ctx context.Context,
	offset, limit int64, desc bool, createStartTime, createEndTime, updateStartTime, updateEndTime *time.Time,
	noticeType *string,
) ([]models.NoticeModel, *int64, error) {
	var noticeList []models.NoticeModel
	var err error
	doc := bson.M{}
	key := fmt.Sprintf("%s:offset:%d:limit:%d", config.NoticeCachePrefix, offset, limit)
	if createStartTime != nil && createEndTime != nil {
		doc["created_at"] = bson.M{"$gte": createStartTime, "$lte": createEndTime}
		key += fmt.Sprintf(
			":createStartTime:%s:createEndTime:%s", createStartTime.String(), createEndTime.String(),
		)
	}
	if updateStartTime != nil && updateEndTime != nil {
		doc["updated_at"] = bson.M{"$gte": updateStartTime, "$lte": updateEndTime}
		key += fmt.Sprintf(
			":updateStartTime:%s:updateEndTime:%s", updateStartTime.String(), updateEndTime.String(),
		)
	}
	if noticeType != nil {
		doc["notice_type"] = *noticeType
		key += fmt.Sprintf(":noticeType:%s", *noticeType)
	}
	docJSON, _ := json.Marshal(doc)

	if desc {
		key += ":desc"
	}
	cache, err := n.Cache.GetList(ctx, key)
	if err != nil {
		n.Core.Logger.Error("NoticeDaoImpl.GetNoticeList: failed to get cache", zap.Error(err), zap.String("key", key))
	} else if cache != nil {
		if noticeList, ok := cache.List.([]models.NoticeModel); ok {
			n.Core.Logger.Info("NoticeDaoImpl.GetNoticeList: cache hit", zap.String("key", key))
			return noticeList, &cache.Total, nil
		} else {
			n.Core.Logger.Error(
				"NoticeDaoImpl.GetNoticeList: failed to cast cache", zap.String("key", key),
			)
		}
	} else {
		n.Core.Logger.Info("NoticeDaoImpl.GetNoticeList: cache miss", zap.String("key", key))
	}

	collection := n.Core.Mongo.MongoClient.Database(n.Core.Mongo.DatabaseName).Collection(config.NoticeCollectionName)
	if desc {
		err = collection.Find(ctx, doc).Sort("-created_at").Skip(offset).Limit(limit).All(&noticeList)
	} else {
		err = collection.Find(ctx, doc).Skip(offset).Limit(limit).All(&noticeList)
	}
	if err != nil {
		n.Core.Logger.Error(
			"NoticeDaoImpl.GetNoticeList: failed to find notices",
			zap.Error(err), zap.ByteString(config.NoticeCollectionName, docJSON),
		)
		return nil, nil, err
	}
	count, err := collection.Find(ctx, doc).Count()
	if err != nil {
		n.Core.Logger.Error(
			"NoticeDaoImpl.GetNoticeList: count failed",
			zap.Error(err), zap.ByteString(config.NoticeCollectionName, docJSON),
		)
		return nil, nil, err
	}
	n.Core.Logger.Info(
		"NoticeDaoImpl.GetNoticeList: success",
		zap.Int64("count", count), zap.ByteString(config.NoticeCollectionName, docJSON),
	)

	cacheList := models.CacheList{Total: count, List: noticeList}
	err = n.Cache.SetList(ctx, key, &cacheList)
	if err != nil {
		n.Core.Logger.Error(
			"NoticeDaoImpl.GetNoticeList: failed to set cache",
			zap.Error(err), zap.String("key", key), zap.ByteString(config.NoticeCollectionName, docJSON),
		)
	} else {
		n.Core.Logger.Info(
			"NoticeDaoImpl.GetNoticeList: cache set", zap.String("key", key),
			zap.ByteString(config.NoticeCollectionName, docJSON),
		)
	}
	return noticeList, &count, nil
}

func (n *NoticeDaoImpl) InsertNotice(
	ctx context.Context, title, content, noticeType string,
) (primitive.ObjectID, error) {
	collection := n.Core.Mongo.MongoClient.Database(n.Core.Mongo.DatabaseName).Collection(config.NoticeCollectionName)
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
		n.Core.Logger.Error(
			"NoticeDaoImpl.InsertNotice: failed to insert notice", zap.Error(err),
			zap.ByteString(config.NoticeCollectionName, docJSON),
		)
	} else {
		n.Core.Logger.Info(
			"NoticeDaoImpl.InsertNotice: success",
			zap.String("noticeID", result.InsertedID.(primitive.ObjectID).Hex()),
			zap.ByteString(config.NoticeCollectionName, docJSON),
		)
		prefix := config.NoticeCachePrefix
		if err = n.Cache.Flush(ctx, &prefix); err != nil {
			n.Core.Logger.Error("NoticeDaoImpl.InsertNotice: failed to flush cache", zap.Error(err))
		} else {
			n.Core.Logger.Info("NoticeDaoImpl.InsertNotice: cache flush success")
		}
	}
	return result.InsertedID.(primitive.ObjectID), err
}

func (n *NoticeDaoImpl) UpdateNotice(
	ctx context.Context, noticeID primitive.ObjectID, title, content, noticeType *string,
) error {
	collection := n.Core.Mongo.MongoClient.Database(n.Core.Mongo.DatabaseName).Collection(config.NoticeCollectionName)
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
		n.Core.Logger.Error(
			"NoticeDaoImpl.UpdateNotice: failed to update notice",
			zap.Error(err), zap.String("noticeID", noticeID.Hex()),
			zap.ByteString(config.NoticeCollectionName, docJSON),
		)
	} else {
		n.Core.Logger.Info(
			"NoticeDaoImpl.UpdateNotice: success",
			zap.String("noticeID", noticeID.Hex()),
			zap.ByteString(config.NoticeCollectionName, docJSON),
		)
		prefix := config.NoticeCachePrefix
		if err = n.Cache.Flush(ctx, &prefix); err != nil {
			n.Core.Logger.Error("NoticeDaoImpl.UpdateNotice: failed to flush cache", zap.Error(err))
		} else {
			n.Core.Logger.Info("NoticeDaoImpl.UpdateNotice: cache flush success")
		}
	}
	return err
}

func (n *NoticeDaoImpl) DeleteNotice(ctx context.Context, noticeID primitive.ObjectID) error {
	collection := n.Core.Mongo.MongoClient.Database(n.Core.Mongo.DatabaseName).Collection(config.NoticeCollectionName)
	err := collection.RemoveId(ctx, noticeID)
	if err != nil {
		n.Core.Logger.Error(
			"NoticeDaoImpl.DeleteNotice: failed to delete notice",
			zap.Error(err), zap.String("noticeID", noticeID.Hex()),
		)
	} else {
		prefix := config.NoticeCachePrefix
		n.Core.Logger.Info("NoticeDaoImpl.DeleteNotice, success", zap.String("noticeID", noticeID.Hex()))
		if err = n.Cache.Flush(ctx, &prefix); err != nil {
			n.Core.Logger.Error("NoticeDaoImpl.DeleteNotice: failed to flush cache", zap.Error(err))
		} else {
			n.Core.Logger.Info("NoticeDaoImpl.DeleteNotice: cache flush success")
		}
	}
	return err
}

func (n *NoticeDaoImpl) DeleteNoticeList(
	ctx context.Context,
	createStartTime, createEndTime, updateStartTime, updateEndTime *time.Time,
	title, content, noticeType *string,
) (*int64, error) {
	collection := n.Core.Mongo.MongoClient.Database(n.Core.Mongo.DatabaseName).Collection(config.NoticeCollectionName)
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
		n.Core.Logger.Error(
			"NoticeDaoImpl.DeleteNoticeList: failed to delete notices",
			zap.Error(err), zap.ByteString(config.NoticeCollectionName, docJSON),
		)
	} else {
		n.Core.Logger.Info(
			"NoticeDaoImpl.DeleteNoticeList: success", zap.ByteString(config.NoticeCollectionName, docJSON),
		)
		prefix := config.NoticeCachePrefix
		if err = n.Cache.Flush(ctx, &prefix); err != nil {
			n.Core.Logger.Error("NoticeDaoImpl.DeleteNoticeList: failed to flush cache", zap.Error(err))
		} else {
			n.Core.Logger.Info("NoticeDaoImpl.DeleteNoticeList: cache flush success")
		}
	}
	return &result.DeletedCount, err
}
