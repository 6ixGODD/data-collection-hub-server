package mods

import (
	"context"
	"time"

	dao "data-collection-hub-server/internal/pkg/dal/mods"
	"data-collection-hub-server/internal/pkg/schema/common"
	"data-collection-hub-server/internal/pkg/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NoticeService interface {
	GetNotice(ctx context.Context, noticeID *primitive.ObjectID) (*common.GetNoticeResponse, error)
	GetNoticeList(
		ctx context.Context, page *int, noticeType *string, updateBefore, updateAfter *time.Time,
	) (*common.GetNoticeListResponse, error)
}

type NoticeServiceImpl struct {
	service   *service.Service
	noticeDao dao.NoticeDao
}

func (n NoticeServiceImpl) GetNotice(ctx context.Context, noticeID *primitive.ObjectID) (
	*common.GetNoticeResponse, error,
) {
	// TODO implement me
	panic("implement me")
}

func (n NoticeServiceImpl) GetNoticeList(
	ctx context.Context, page *int, noticeType *string, updateBefore, updateAfter *time.Time,
) (*common.GetNoticeListResponse, error) {
	// TODO implement me
	panic("implement me")
}

func NewNoticeService(s *service.Service, noticeDao dao.NoticeDao) NoticeService {
	return &NoticeServiceImpl{
		service:   s,
		noticeDao: noticeDao,
	}
}
