package mods

import (
	"context"

	"data-collection-hub-server/internal/pkg/dao/mods"
	"data-collection-hub-server/internal/pkg/service"
	"data-collection-hub-server/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LogsService interface {
	InsertLoginLog(
		ctx context.Context, userID *primitive.ObjectID, username, email, ipAddress, userAgent *string,
	) error
	CacheLoginLog(ctx context.Context, username, ipAddress, userAgent *string) error
	InsertOperationLog(
		ctx context.Context, userID, entityID *primitive.ObjectID,
		username, email, ipAddress, userAgent, operation, entityType, description, status *string,
	) error
	CacheOperationLog(
		ctx context.Context, userID, entityID *primitive.ObjectID, ipAddress, userAgent *string,
		operation, entityType, description, status *string,
	) error
}

type logsServiceImpl struct {
	core            *service.Core
	loginLogDao     mods.LoginLogDao
	operationLogDao mods.OperationLogDao
	userDao         mods.UserDao
}

func NewLogsService(
	core *service.Core, loginLogDao mods.LoginLogDao, operationLogDao mods.OperationLogDao,
) LogsService {
	return &logsServiceImpl{
		core:            core,
		loginLogDao:     loginLogDao,
		operationLogDao: operationLogDao,
	}
}

func (l logsServiceImpl) InsertLoginLog(
	ctx context.Context, userID *primitive.ObjectID, username, email, ipAddress, userAgent *string,
) error {
	_, err := l.loginLogDao.InsertLoginLog(ctx, *userID, *username, *email, *ipAddress, *userAgent)
	if err != nil {
		return errors.DBError(errors.WriteError(err))
	}
	return nil
}

func (l logsServiceImpl) CacheLoginLog(
	ctx context.Context, username, ipAddress, userAgent *string,
) error {
	err := l.loginLogDao.CacheLoginLog(ctx, *username, *ipAddress, *userAgent)
	if err != nil {
		return errors.CacheError(errors.WriteError(err))
	}
	return nil
}

func (l logsServiceImpl) InsertOperationLog(
	ctx context.Context, userID, entityID *primitive.ObjectID,
	username, email, ipAddress, userAgent, operation, entityType, description, status *string,
) error {
	_, err := l.operationLogDao.InsertOperationLog(
		ctx, *userID, *entityID, *username, *email, *ipAddress, *userAgent, *operation, *entityType, *description,
		*status,
	)
	if err != nil {
		return errors.DBError(errors.WriteError(err))
	}
	return nil
}

func (l logsServiceImpl) CacheOperationLog(
	ctx context.Context, userID, entityID *primitive.ObjectID, ipAddress, userAgent *string,
	operation, entityType, description, status *string,
) error {
	err := l.operationLogDao.CacheOperationLog(
		ctx, *userID, *entityID, *ipAddress, *userAgent, *operation, *entityType, *description, *status,
	)
	if err != nil {
		return errors.CacheError(errors.WriteError(err))
	}
	return nil
}
