package mods

import (
	"context"
	"time"

	dao "data-collection-hub-server/internal/pkg/dao/mods"
	"data-collection-hub-server/internal/pkg/schema/admin"
	"data-collection-hub-server/internal/pkg/service"
	"data-collection-hub-server/pkg/errors"
	"data-collection-hub-server/pkg/utils/crypt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService interface {
	InsertUser(ctx context.Context, username, email, password, role, organization *string) error
	GetUser(ctx context.Context, userID *primitive.ObjectID) (*admin.GetUserResponse, error)
	GetUserList(
		ctx context.Context, page, pageSize *int64, desc *bool, role *string,
		lastLoginBefore, lastLoginAfter, createdBefore, createdAfter *time.Time, query *string,
	) (*admin.GetUserListResponse, error)
	UpdateUser(ctx context.Context, userID *primitive.ObjectID, username, email, role, organization *string) error
	DeleteUser(ctx context.Context, userID *primitive.ObjectID) error
	ChangeUserPassword(ctx context.Context, userID *primitive.ObjectID, newPassword *string) error
}

type UserServiceImpl struct {
	service *service.Core
	userDao dao.UserDao
}

func NewUserService(s *service.Core, userDao dao.UserDao) UserService {
	return &UserServiceImpl{
		service: s,
		userDao: userDao,
	}
}

func (u UserServiceImpl) InsertUser(ctx context.Context, username, email, password, role, organization *string) error {
	passwordHash, err := crypt.PasswordHash(*password)
	_, err = u.userDao.InsertUser(ctx, *username, *email, passwordHash, *role, *organization)
	if err != nil {
		return errors.MongoError(errors.WriteError(err))
	}
	return nil
}

func (u UserServiceImpl) GetUser(ctx context.Context, userID *primitive.ObjectID) (*admin.GetUserResponse, error) {
	user, err := u.userDao.GetUserById(ctx, *userID)
	if err != nil {
		return nil, errors.MongoError(errors.ReadError(err))
	}
	return &admin.GetUserResponse{
		UserID:       user.UserID.Hex(),
		Username:     user.Username,
		Email:        user.Email,
		Password:     user.Password,
		Role:         user.Role,
		Organization: user.Organization,
		LastLogin:    user.LastLogin.Format(time.RFC3339),
		CreatedAt:    user.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    user.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (u UserServiceImpl) GetUserList(
	ctx context.Context, page, pageSize *int64, desc *bool, role *string,
	lastLoginBefore, lastLoginAfter, createdBefore, createdAfter *time.Time, query *string,
) (*admin.GetUserListResponse, error) {
	users, count, err := u.userDao.GetUserList(
		ctx, *page, *pageSize, *desc, nil, role, createdBefore, createdAfter,
		nil, nil, lastLoginBefore, lastLoginAfter, query,
	)
	if err != nil {
		return nil, errors.MongoError(errors.ReadError(err))
	}
	resp := make([]*admin.GetUserResponse, 0, len(users))
	for _, user := range users {
		resp = append(
			resp, &admin.GetUserResponse{
				UserID:       user.UserID.Hex(),
				Username:     user.Username,
				Email:        user.Email,
				Password:     user.Password,
				Role:         user.Role,
				Organization: user.Organization,
				LastLogin:    user.LastLogin.Format(time.RFC3339),
				CreatedAt:    user.CreatedAt.Format(time.RFC3339),
				UpdatedAt:    user.UpdatedAt.Format(time.RFC3339),
			},
		)
	}
	return &admin.GetUserListResponse{
		Total:    *count,
		UserList: resp,
	}, nil
}

func (u UserServiceImpl) UpdateUser(
	ctx context.Context, userID *primitive.ObjectID, username, email, role, organization *string,
) error {
	err := u.userDao.UpdateUser(ctx, *userID, username, email, nil, role, organization)
	if err != nil {
		return errors.MongoError(errors.WriteError(err))
	}
	return nil
}

func (u UserServiceImpl) DeleteUser(ctx context.Context, userID *primitive.ObjectID) error {
	err := u.userDao.DeleteUser(ctx, *userID)
	if err != nil {
		return errors.MongoError(errors.WriteError(err))
	}
	return nil
}

func (u UserServiceImpl) ChangeUserPassword(
	ctx context.Context, userID *primitive.ObjectID, newPassword *string,
) error {
	err := u.userDao.UpdateUser(ctx, *userID, nil, nil, newPassword, nil, nil)
	if err != nil {
		return errors.MongoError(errors.WriteError(err))
	}
	return nil
}
