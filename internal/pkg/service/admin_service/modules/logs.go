package modules

import (
	dao "data-collection-hub-server/internal/pkg/dal/modules"
	"data-collection-hub-server/internal/pkg/service"
)

type LogsService interface {
}

type LogsServiceImpl struct {
	core            *service.Core
	loginLogDao     dao.LoginLogDao
	operationLogDao dao.OperationLogDao
	errorLogDao     dao.ErrorLogDao
}

func NewLogsService(s *service.Core, loginLogDao dao.LoginLogDao, operationLogDao dao.OperationLogDao, errorLogDao dao.ErrorLogDao) LogsService {
	return &LogsServiceImpl{
		core:            s,
		loginLogDao:     loginLogDao,
		operationLogDao: operationLogDao,
		errorLogDao:     errorLogDao,
	}
}
