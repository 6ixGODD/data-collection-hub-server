package modules

import (
	"context"

	"data-collection-hub-server/internal/pkg/dal"
	"data-collection-hub-server/internal/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var NoticeCollectionName = "notice"

type NoticeDao interface {
	GetNoticeById(noticeId string, ctx context.Context) (*models.NoticeModel, error)
	GetNoticeList(offset, limit int64, ctx context.Context) ([]models.NoticeModel, error)
	GetNoticeListByType(noticeType string, offset, limit int64, ctx context.Context) ([]models.NoticeModel, error)
	GetNoticeListByTypeAndCreatedTime(noticeType, startTime, endTime string, offset, limit int64, ctx context.Context) ([]models.NoticeModel, error)
	GetNoticeListByTypeAndUpdatedTime(noticeType, startTime, endTime string, offset, limit int64, ctx context.Context) ([]models.NoticeModel, error)
	GetNoticeListByCreatedTime(startTime, endTime string, offset, limit int64, ctx context.Context) ([]models.NoticeModel, error)
	GetNoticeListByUpdatedTime(startTime, endTime string, offset, limit int64, ctx context.Context) ([]models.NoticeModel, error)
	InsertNotice(notice *models.NoticeModel, ctx context.Context) error
	UpdateNotice(notice *models.NoticeModel, ctx context.Context) error
	DeleteNotice(notice *models.NoticeModel, ctx context.Context) error
}

type NoticeDaoImpl struct{ *dal.Dao }

func NewNoticeDao(dao *dal.Dao) NoticeDao {
	var _ NoticeDao = new(NoticeDaoImpl)
	return &NoticeDaoImpl{dao}
}

func (n NoticeDaoImpl) GetNoticeById(noticeId string, ctx context.Context) (*models.NoticeModel, error) {
	var notice models.NoticeModel
	collection := n.Dao.MongoDB.Collection(NoticeCollectionName)
	err := collection.Find(ctx, bson.M{"_id": noticeId}).One(&notice)
	if err != nil {
		n.Dao.Logger.Error(
			"NoticeDaoImpl.GetNoticeById",
			zap.Field{Key: "noticeId", Type: zapcore.StringType, String: noticeId},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		n.Dao.Logger.Info(
			"NoticeDaoImpl.GetNoticeById",
			zap.Field{Key: "noticeId", Type: zapcore.StringType, String: noticeId},
		)
		return &notice, nil
	}
}

func (n NoticeDaoImpl) GetNoticeList(offset, limit int64, ctx context.Context) ([]models.NoticeModel, error) {
	var noticeList []models.NoticeModel
	collection := n.Dao.MongoDB.Collection(NoticeCollectionName)
	err := collection.Find(ctx, bson.M{}).Skip(offset).Limit(limit).All(&noticeList)
	if err != nil {
		n.Dao.Logger.Error(
			"NoticeDaoImpl.GetNoticeList",
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		n.Dao.Logger.Info(
			"NoticeDaoImpl.GetNoticeList",
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return noticeList, nil
	}
}

func (n NoticeDaoImpl) GetNoticeListByType(noticeType string, offset, limit int64, ctx context.Context) ([]models.NoticeModel, error) {
	var noticeList []models.NoticeModel
	collection := n.Dao.MongoDB.Collection(NoticeCollectionName)
	err := collection.Find(ctx, bson.M{"notice_type": noticeType}).Skip(offset).Limit(limit).All(&noticeList)
	if err != nil {
		n.Dao.Logger.Error(
			"NoticeDaoImpl.GetNoticeListByType",
			zap.Field{Key: "noticeType", Type: zapcore.StringType, String: noticeType},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		n.Dao.Logger.Info(
			"NoticeDaoImpl.GetNoticeListByType",
			zap.Field{Key: "noticeType", Type: zapcore.StringType, String: noticeType},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return noticeList, nil
	}
}

func (n NoticeDaoImpl) GetNoticeListByTypeAndCreatedTime(noticeType, startTime, endTime string, offset, limit int64, ctx context.Context) ([]models.NoticeModel, error) {
	var noticeList []models.NoticeModel
	collection := n.Dao.MongoDB.Collection(NoticeCollectionName)
	err := collection.Find(ctx, bson.M{"notice_type": noticeType, "created_time": bson.M{"$gte": startTime, "$lte": endTime}}).Skip(offset).Limit(limit).All(&noticeList)
	if err != nil {
		n.Dao.Logger.Error(
			"NoticeDaoImpl.GetNoticeListByTypeAndCreatedTime",
			zap.Field{Key: "noticeType", Type: zapcore.StringType, String: noticeType},
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		n.Dao.Logger.Info(
			"NoticeDaoImpl.GetNoticeListByTypeAndCreatedTime",
			zap.Field{Key: "noticeType", Type: zapcore.StringType, String: noticeType},
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return noticeList, nil
	}
}

func (n NoticeDaoImpl) GetNoticeListByTypeAndUpdatedTime(noticeType, startTime, endTime string, offset, limit int64, ctx context.Context) ([]models.NoticeModel, error) {
	var noticeList []models.NoticeModel
	collection := n.Dao.MongoDB.Collection(NoticeCollectionName)
	err := collection.Find(ctx, bson.M{"notice_type": noticeType, "updated_time": bson.M{"$gte": startTime, "$lte": endTime}}).Skip(offset).Limit(limit).All(&noticeList)
	if err != nil {
		n.Dao.Logger.Error(
			"NoticeDaoImpl.GetNoticeListByTypeAndUpdatedTime",
			zap.Field{Key: "noticeType", Type: zapcore.StringType, String: noticeType},
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		n.Dao.Logger.Info(
			"NoticeDaoImpl.GetNoticeListByTypeAndUpdatedTime",
			zap.Field{Key: "noticeType", Type: zapcore.StringType, String: noticeType},
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return noticeList, nil
	}
}

func (n NoticeDaoImpl) GetNoticeListByCreatedTime(startTime, endTime string, offset, limit int64, ctx context.Context) ([]models.NoticeModel, error) {
	var noticeList []models.NoticeModel
	collection := n.Dao.MongoDB.Collection(NoticeCollectionName)
	err := collection.Find(ctx, bson.M{"created_time": bson.M{"$gte": startTime, "$lte": endTime}}).Skip(offset).Limit(limit).All(&noticeList)
	if err != nil {
		n.Dao.Logger.Error(
			"NoticeDaoImpl.GetNoticeListByCreatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		n.Dao.Logger.Info(
			"NoticeDaoImpl.GetNoticeListByCreatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return noticeList, nil
	}
}

func (n NoticeDaoImpl) GetNoticeListByUpdatedTime(startTime, endTime string, offset, limit int64, ctx context.Context) ([]models.NoticeModel, error) {
	var noticeList []models.NoticeModel
	collection := n.Dao.MongoDB.Collection(NoticeCollectionName)
	err := collection.Find(ctx, bson.M{"updated_time": bson.M{"$gte": startTime, "$lte": endTime}}).Skip(offset).Limit(limit).All(&noticeList)
	if err != nil {
		n.Dao.Logger.Error(
			"NoticeDaoImpl.GetNoticeListByUpdatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		n.Dao.Logger.Info(
			"NoticeDaoImpl.GetNoticeListByUpdatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return noticeList, nil
	}
}

func (n NoticeDaoImpl) InsertNotice(notice *models.NoticeModel, ctx context.Context) error {
	collection := n.Dao.MongoDB.Collection(NoticeCollectionName)
	result, err := collection.InsertOne(ctx, notice)
	if err != nil {
		n.Dao.Logger.Error(
			"NoticeDaoImpl.InsertNotice",
			zap.Field{Key: "notice", Type: zapcore.ReflectType, Interface: notice},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
	} else {
		n.Dao.Logger.Info(
			"NoticeDaoImpl.InsertNotice",
			zap.Field{Key: "notice", Type: zapcore.ReflectType, Interface: notice},
			zap.Field{Key: "result", Type: zapcore.ReflectType, Interface: result},
		)
	}
	return err
}

func (n NoticeDaoImpl) UpdateNotice(notice *models.NoticeModel, ctx context.Context) error {
	collection := n.Dao.MongoDB.Collection(NoticeCollectionName)
	err := collection.UpdateOne(ctx, bson.M{"_id": notice.NoticeID}, bson.M{"$set": notice})
	if err != nil {
		n.Dao.Logger.Error(
			"NoticeDaoImpl.UpdateNotice",
			zap.Field{Key: "notice", Type: zapcore.ReflectType, Interface: notice},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
	} else {
		n.Dao.Logger.Info(
			"NoticeDaoImpl.UpdateNotice",
			zap.Field{Key: "notice", Type: zapcore.ReflectType, Interface: notice},
		)
	}
	return err
}

func (n NoticeDaoImpl) DeleteNotice(notice *models.NoticeModel, ctx context.Context) error {
	collection := n.Dao.MongoDB.Collection(NoticeCollectionName)
	err := collection.RemoveId(ctx, notice.NoticeID)
	if err != nil {
		n.Dao.Logger.Error(
			"NoticeDaoImpl.DeleteNotice",
			zap.Field{Key: "notice", Type: zapcore.ReflectType, Interface: notice},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
	} else {
		n.Dao.Logger.Info(
			"NoticeDaoImpl.DeleteNotice",
			zap.Field{Key: "notice", Type: zapcore.ReflectType, Interface: notice},
		)
	}
	return err
}
