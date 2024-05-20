package mods

import (
	"context"

	dao "data-collection-hub-server/internal/pkg/dao/mods"
	"data-collection-hub-server/internal/pkg/service"
	"data-collection-hub-server/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NoticeService interface {
	InsertNotice(ctx context.Context, title, content, noticeType *string) (string, error)
	UpdateNotice(ctx context.Context, noticeID *primitive.ObjectID, title, content, noticeType *string) error
	DeleteNotice(ctx context.Context, noticeID *primitive.ObjectID) error
}

type NoticeServiceImpl struct {
	core      *service.Core
	noticeDao dao.NoticeDao
}

func NewNoticeService(core *service.Core, noticeDao dao.NoticeDao) NoticeService {
	return &NoticeServiceImpl{
		core:      core,
		noticeDao: noticeDao,
	}
}

func (n NoticeServiceImpl) InsertNotice(ctx context.Context, title, content, noticeType *string) (string, error) {
	noticeID, err := n.noticeDao.InsertNotice(ctx, *title, *content, *noticeType)
	if err != nil {
		return "", errors.DBError(errors.WriteError(err))
	}
	return noticeID.Hex(), nil
}

func (n NoticeServiceImpl) UpdateNotice(
	ctx context.Context, noticeID *primitive.ObjectID, title, content, noticeType *string,
) error {
	err := n.noticeDao.UpdateNotice(ctx, *noticeID, title, content, noticeType)
	if err != nil {
		return errors.DBError(errors.WriteError(err))
	}
	return nil
}

func (n NoticeServiceImpl) DeleteNotice(ctx context.Context, noticeID *primitive.ObjectID) error {
	err := n.noticeDao.DeleteNotice(ctx, *noticeID)
	if err != nil {
		return errors.DBError(errors.WriteError(err))
	}
	return nil
}
