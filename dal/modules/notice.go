package modules

import (
	"context"

	"data-collection-hub-server/models"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
)

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

type NoticeDaoImpl struct {
	noticeClient *qmgo.Collection
}

func NewNoticeDao(mongoDatabase *qmgo.Database) NoticeDao {
	var _ NoticeDao = new(NoticeDaoImpl)
	return &NoticeDaoImpl{noticeClient: mongoDatabase.Collection("notice")}
}

func (n NoticeDaoImpl) GetNoticeById(noticeId string, ctx context.Context) (*models.NoticeModel, error) {
	var notice models.NoticeModel
	err := n.noticeClient.Find(ctx, bson.M{"_id": noticeId}).One(&notice)
	if err != nil {
		return nil, err
	} else {
		return &notice, nil
	}
}

func (n NoticeDaoImpl) GetNoticeList(offset, limit int64, ctx context.Context) ([]models.NoticeModel, error) {
	var noticeList []models.NoticeModel
	err := n.noticeClient.Find(ctx, bson.M{}).Skip(offset).Limit(limit).All(&noticeList)
	if err != nil {
		return nil, err
	} else {
		return noticeList, nil
	}
}

func (n NoticeDaoImpl) GetNoticeListByType(noticeType string, offset, limit int64, ctx context.Context) ([]models.NoticeModel, error) {
	var noticeList []models.NoticeModel
	err := n.noticeClient.Find(ctx, bson.M{"notice_type": noticeType}).Skip(offset).Limit(limit).All(&noticeList)
	if err != nil {
		return nil, err
	} else {
		return noticeList, nil
	}
}

func (n NoticeDaoImpl) GetNoticeListByTypeAndCreatedTime(noticeType, startTime, endTime string, offset, limit int64, ctx context.Context) ([]models.NoticeModel, error) {
	var noticeList []models.NoticeModel
	err := n.noticeClient.Find(ctx, bson.M{"notice_type": noticeType, "created_time": bson.M{"$gte": startTime, "$lte": endTime}}).Skip(offset).Limit(limit).All(&noticeList)
	if err != nil {
		return nil, err
	} else {
		return noticeList, nil
	}
}

func (n NoticeDaoImpl) GetNoticeListByTypeAndUpdatedTime(noticeType, startTime, endTime string, offset, limit int64, ctx context.Context) ([]models.NoticeModel, error) {
	var noticeList []models.NoticeModel
	err := n.noticeClient.Find(ctx, bson.M{"notice_type": noticeType, "updated_time": bson.M{"$gte": startTime, "$lte": endTime}}).Skip(offset).Limit(limit).All(&noticeList)
	if err != nil {
		return nil, err
	} else {
		return noticeList, nil
	}
}

func (n NoticeDaoImpl) GetNoticeListByCreatedTime(startTime, endTime string, offset, limit int64, ctx context.Context) ([]models.NoticeModel, error) {
	var noticeList []models.NoticeModel
	err := n.noticeClient.Find(ctx, bson.M{"created_time": bson.M{"$gte": startTime, "$lte": endTime}}).Skip(offset).Limit(limit).All(&noticeList)
	if err != nil {
		return nil, err
	} else {
		return noticeList, nil
	}
}

func (n NoticeDaoImpl) GetNoticeListByUpdatedTime(startTime, endTime string, offset, limit int64, ctx context.Context) ([]models.NoticeModel, error) {
	var noticeList []models.NoticeModel
	err := n.noticeClient.Find(ctx, bson.M{"updated_time": bson.M{"$gte": startTime, "$lte": endTime}}).Skip(offset).Limit(limit).All(&noticeList)
	if err != nil {
		return nil, err
	} else {
		return noticeList, nil
	}
}

func (n NoticeDaoImpl) InsertNotice(notice *models.NoticeModel, ctx context.Context) error {
	_, err := n.noticeClient.InsertOne(ctx, notice)
	return err
}

func (n NoticeDaoImpl) UpdateNotice(notice *models.NoticeModel, ctx context.Context) error {
	err := n.noticeClient.UpdateOne(ctx, bson.M{"_id": notice.NoticeID}, bson.M{"$set": notice})
	return err
}

func (n NoticeDaoImpl) DeleteNotice(notice *models.NoticeModel, ctx context.Context) error {
	err := n.noticeClient.RemoveId(ctx, notice.NoticeID)
	return err
}
