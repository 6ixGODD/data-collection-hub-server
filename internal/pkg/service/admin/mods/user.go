package mods

import (
	"context"
	"time"

	dao "data-collection-hub-server/internal/pkg/dal/mods"
	"data-collection-hub-server/internal/pkg/schema/admin"
	"data-collection-hub-server/internal/pkg/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService interface {
	InsertUser(ctx context.Context, username, email, password, role, organization *string) error
	GetUser(ctx context.Context, userID *primitive.ObjectID) (*admin.GetUserResponse, error)
	GetUserList(
		ctx context.Context, page *int, role *string,
		lastLoginBefore, lastLoginAfter, createdBefore, createdAfter *time.Time,
	) (*admin.GetUserListResponse, error)
	UpdateUser(ctx context.Context, userID *primitive.ObjectID, username, email, role, organization *string) error
	DeleteUser(ctx context.Context, userID *primitive.ObjectID) error
	ChangeUserPassword(ctx context.Context, userID *primitive.ObjectID, newPassword *string) error
}

type UserServiceImpl struct {
	service *service.Service
	userDao dao.UserDao
}

func NewUserService(s *service.Service, userDao dao.UserDao) UserService {
	return &UserServiceImpl{
		service: s,
		userDao: userDao,
	}
}

func (u UserServiceImpl) InsertUser(ctx context.Context, username, email, password, role, organization *string) error {
	// TODO implement me
	panic("implement me")
}

func (u UserServiceImpl) GetUser(ctx context.Context, userID *primitive.ObjectID) (*admin.GetUserResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (u UserServiceImpl) GetUserList(
	ctx context.Context, page *int, role *string,
	lastLoginBefore, lastLoginAfter, createdBefore, createdAfter *time.Time,
) (*admin.GetUserListResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (u UserServiceImpl) UpdateUser(
	ctx context.Context, userID *primitive.ObjectID, username, email, role, organization *string,
) error {
	// TODO implement me
	panic("implement me")
}

func (u UserServiceImpl) DeleteUser(ctx context.Context, userID *primitive.ObjectID) error {
	// TODO implement me
	panic("implement me")
}

func (u UserServiceImpl) ChangeUserPassword(
	ctx context.Context, userID *primitive.ObjectID, newPassword *string,
) error {
	// TODO implement me
	panic("implement me")
}
