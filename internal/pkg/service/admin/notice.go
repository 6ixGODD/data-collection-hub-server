package admin

import (
	dao "data-collection-hub-server/internal/pkg/dal/modules"
	"data-collection-hub-server/internal/pkg/service"
)

type NoticeService interface {
}

type NoticeServiceImpl struct {
	Service   *service.Service
	NoticeDao dao.NoticeDao
}

func NewNoticeService(s *service.Service, noticeDaoImpl *dao.NoticeDaoImpl) NoticeService {
	return &NoticeServiceImpl{
		Service:   s,
		NoticeDao: noticeDaoImpl,
	}
}
