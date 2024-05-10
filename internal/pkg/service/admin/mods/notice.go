package mods

import (
	"context"

	dao "data-collection-hub-server/internal/pkg/dal/mods"
	"data-collection-hub-server/internal/pkg/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NoticeService interface {
	InsertNotice(ctx context.Context, title, content *string) error
	UpdateNotice(ctx context.Context, noticeID *primitive.ObjectID, title, content *string) error
	DeleteNotice(ctx context.Context, noticeID *primitive.ObjectID) error
}

type NoticeServiceImpl struct {
	service   *service.Service
	noticeDao dao.NoticeDao
}

func NewNoticeService(s *service.Service, noticeDao dao.NoticeDao) NoticeService {
	return &NoticeServiceImpl{
		service:   s,
		noticeDao: noticeDao,
	}
}

func (n NoticeServiceImpl) InsertNotice(ctx context.Context, title, content *string) error {
	// TODO implement me
	panic("implement me")
}

func (n NoticeServiceImpl) UpdateNotice(
	ctx context.Context, noticeID *primitive.ObjectID, title, content *string,
) error {
	// TODO implement me
	panic("implement me")
}

func (n NoticeServiceImpl) DeleteNotice(ctx context.Context, noticeID *primitive.ObjectID) error {
	// TODO implement me
	panic("implement me")
}
