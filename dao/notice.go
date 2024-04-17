package dao

import (
	"context"
	"data-collection-hub-server/models"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
)

type NoticeDao interface {
	GetNoticeById(noticeId string, mongoClient *qmgo.QmgoClient, ctx context.Context) (*models.NoticeModel, error)
	GetNoticeList(mongoClient *qmgo.QmgoClient, offset, limit int64, ctx context.Context) ([]models.NoticeModel, error)
	GetNoticeListByType(mongoClient *qmgo.QmgoClient, noticeType string, offset, limit int64, ctx context.Context) ([]models.NoticeModel, error)
	GetNoticeListByTypeAndCreatedTime(mongoClient *qmgo.QmgoClient, noticeType, startTime, endTime string, offset, limit int64, ctx context.Context) ([]models.NoticeModel, error)
	GetNoticeListByTypeAndUpdatedTime(mongoClient *qmgo.QmgoClient, noticeType, startTime, endTime string, offset, limit int64, ctx context.Context) ([]models.NoticeModel, error)
	GetNoticeListByCreatedTime(mongoClient *qmgo.QmgoClient, startTime, endTime string, offset, limit int64, ctx context.Context) ([]models.NoticeModel, error)
	GetNoticeListByUpdatedTime(mongoClient *qmgo.QmgoClient, startTime, endTime string, offset, limit int64, ctx context.Context) ([]models.NoticeModel, error)
	InsertNotice(notice *models.NoticeModel, mongoClient *qmgo.QmgoClient, ctx context.Context) error
	UpdateNotice(notice *models.NoticeModel, mongoClient *qmgo.QmgoClient, ctx context.Context) error
	DeleteNotice(notice *models.NoticeModel, mongoClient *qmgo.QmgoClient, ctx context.Context) error
}

type NoticeDaoImpl struct{}

func (n NoticeDaoImpl) GetNoticeById(noticeId string, mongoClient *qmgo.QmgoClient, ctx context.Context) (*models.NoticeModel, error) {
	var notice models.NoticeModel
	err := mongoClient.Find(ctx, bson.M{"_id": noticeId}).One(&notice)
	if err != nil {
		return nil, err
	} else {
		return &notice, nil
	}
}

func (n NoticeDaoImpl) GetNoticeList(mongoClient *qmgo.QmgoClient, offset, limit int64, ctx context.Context) ([]models.NoticeModel, error) {
	//TODO implement me
	panic("implement me")
}

func (n NoticeDaoImpl) GetNoticeListByType(mongoClient *qmgo.QmgoClient, noticeType string, offset, limit int64, ctx context.Context) ([]models.NoticeModel, error) {
	//TODO implement me
	panic("implement me")
}

func (n NoticeDaoImpl) GetNoticeListByTypeAndCreatedTime(mongoClient *qmgo.QmgoClient, noticeType, startTime, endTime string, offset, limit int64, ctx context.Context) ([]models.NoticeModel, error) {
	//TODO implement me
	panic("implement me")
}

func (n NoticeDaoImpl) GetNoticeListByTypeAndUpdatedTime(mongoClient *qmgo.QmgoClient, noticeType, startTime, endTime string, offset, limit int64, ctx context.Context) ([]models.NoticeModel, error) {
	//TODO implement me
	panic("implement me")
}

func (n NoticeDaoImpl) GetNoticeListByCreatedTime(mongoClient *qmgo.QmgoClient, startTime, endTime string, offset, limit int64, ctx context.Context) ([]models.NoticeModel, error) {
	//TODO implement me
	panic("implement me")
}

func (n NoticeDaoImpl) GetNoticeListByUpdatedTime(mongoClient *qmgo.QmgoClient, startTime, endTime string, offset, limit int64, ctx context.Context) ([]models.NoticeModel, error) {
	//TODO implement me
	panic("implement me")
}

func (n NoticeDaoImpl) InsertNotice(notice *models.NoticeModel, mongoClient *qmgo.QmgoClient, ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (n NoticeDaoImpl) UpdateNotice(notice *models.NoticeModel, mongoClient *qmgo.QmgoClient, ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (n NoticeDaoImpl) DeleteNotice(notice *models.NoticeModel, mongoClient *qmgo.QmgoClient, ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func NewNoticeDao() NoticeDao {
	var _ NoticeDao = new(NoticeDaoImpl)
	return &NoticeDaoImpl{}
}
