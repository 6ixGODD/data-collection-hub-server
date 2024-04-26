package admin

import (
	dao "data-collection-hub-server/dal/modules"
	"data-collection-hub-server/service"
)

type NoticeService interface {
}

type NoticeServiceImpl struct {
	*service.Service
	dao.NoticeDao
}

func NewNoticeService(s *service.Service, noticeDaoImpl *dao.NoticeDaoImpl) NoticeService {
	return &NoticeServiceImpl{
		Service:   s,
		NoticeDao: noticeDaoImpl,
	}
}
