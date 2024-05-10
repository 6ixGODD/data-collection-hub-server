package mods

import (
	"context"

	dao "data-collection-hub-server/internal/pkg/dal/mods"
	"data-collection-hub-server/internal/pkg/schema/common"
	"data-collection-hub-server/internal/pkg/service"
)

type ProfileService interface {
	GetProfile(ctx context.Context) (*common.GetProfileResponse, error)
}

type ProfileServiceImpl struct {
	service *service.Service
	userDao dao.UserDao
}

func NewProfileService(s *service.Service, userDao dao.UserDao) ProfileService {
	var _ ProfileService = (*ProfileServiceImpl)(nil)
	return &ProfileServiceImpl{
		service: s,
		userDao: userDao,
	}
}

func (p ProfileServiceImpl) GetProfile(ctx context.Context) (*common.GetProfileResponse, error) {
	// TODO implement me
	panic("implement me")
}
