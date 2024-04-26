package admin

import (
	dao "data-collection-hub-server/dal/modules"
	"data-collection-hub-server/service"
)

type LogsService interface {
}

type LogsServiceImpl struct {
	*service.Service
	dao.LoginLogDao
	dao.OperationLogDao
	dao.ErrorLogDao
}

func NewLogsService(s *service.Service, loginLogDaoImpl *dao.LoginLogDaoImpl, operationLogDaoImpl *dao.OperationLogDaoImpl, errorLogDaoImpl *dao.ErrorLogDaoImpl) LogsService {
	return &LogsServiceImpl{
		Service:         s,
		LoginLogDao:     loginLogDaoImpl,
		OperationLogDao: operationLogDaoImpl,
		ErrorLogDao:     errorLogDaoImpl,
	}
}
