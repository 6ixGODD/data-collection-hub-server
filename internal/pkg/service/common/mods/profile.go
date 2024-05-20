package mods

import (
	"context"
	"time"

	"data-collection-hub-server/internal/pkg/config"
	dao "data-collection-hub-server/internal/pkg/dao/mods"
	"data-collection-hub-server/internal/pkg/schema/common"
	"data-collection-hub-server/internal/pkg/service"
	"data-collection-hub-server/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProfileService interface {
	GetProfile(ctx context.Context) (*common.GetProfileResponse, error)
}

type ProfileServiceImpl struct {
	core    *service.Core
	userDao dao.UserDao
}

func NewProfileService(core *service.Core, userDao dao.UserDao) ProfileService {
	return &ProfileServiceImpl{
		core:    core,
		userDao: userDao,
	}
}

func (p ProfileServiceImpl) GetProfile(ctx context.Context) (*common.GetProfileResponse, error) {
	userID, err := primitive.ObjectIDFromHex(ctx.Value(config.UserIDKey).(string))
	if err != nil {
		return nil, errors.UserNotFound(err) // TODO: change error type
	}
	user, err := p.userDao.GetUserByID(ctx, userID)
	if err != nil {
		return nil, errors.DBError(errors.ReadError(err))
	}
	return &common.GetProfileResponse{
		UserID:       user.UserID.Hex(),
		Username:     user.Username,
		Email:        user.Email,
		Role:         user.Role,
		Organization: user.Organization,
		LastLogin:    user.LastLogin.Format(time.RFC3339),
		CreatedAt:    user.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    user.UpdatedAt.Format(time.RFC3339),
	}, nil
}
