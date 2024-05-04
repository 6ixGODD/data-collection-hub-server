package modules

import (
	dao "data-collection-hub-server/internal/pkg/dal/modules"
	"data-collection-hub-server/internal/pkg/service"
)

type NoticeService interface {
}

type NoticeServiceImpl struct {
	core      *service.Core
	noticeDao dao.NoticeDao
}

func NewNoticeService(s *service.Core, noticeDao dao.NoticeDao) NoticeService {
	return &NoticeServiceImpl{
		core:      s,
		noticeDao: noticeDao,
	}
}
