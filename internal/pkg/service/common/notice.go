package common

import (
	dao "data-collection-hub-server/internal/pkg/dal/modules"
	"data-collection-hub-server/internal/pkg/service"
)

type NoticeService interface {
}

type noticeServiceImpl struct {
	service   *service.Service
	noticeDao dao.NoticeDao
}

func NewNoticeService(s *service.Service, noticeDao dao.NoticeDao) NoticeService {
	return &noticeServiceImpl{
		service:   s,
		noticeDao: noticeDao,
	}
}
