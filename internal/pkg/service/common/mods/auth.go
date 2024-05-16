package mods

import (
	"context"

	"data-collection-hub-server/internal/pkg/config"
	dao "data-collection-hub-server/internal/pkg/dao/mods"
	"data-collection-hub-server/internal/pkg/schema/common"
	"data-collection-hub-server/internal/pkg/service"
	"data-collection-hub-server/pkg/errors"
	"data-collection-hub-server/pkg/jwt"
	"data-collection-hub-server/pkg/utils/crypt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthService interface {
	Login(ctx context.Context, email, password *string) (*common.LoginResponse, error)
	RefreshToken(ctx context.Context, refreshToken *string) (*common.RefreshTokenResponse, error)
	Logout(ctx context.Context) error
	ChangePassword(ctx context.Context, oldPassword, newPassword *string) error
}

type AuthServiceImpl struct {
	service     *service.Service
	userDao     dao.UserDao
	loginLogDao dao.LoginLogDao
	Jwt         *jwt.Jwt
}

func (a AuthServiceImpl) Login(ctx context.Context, email, password *string) (*common.LoginResponse, error) {
	user, err := a.userDao.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, errors.MongoError(errors.ReadError(err))
	}
	if !crypt.VerifyPassword(*password, user.Password) {
		return nil, errors.PasswordWrong(err)
	}
	accessToken, err := a.Jwt.GenerateAccessToken(user.UserID.Hex())
	if err != nil {
		return nil, errors.TokenGenerateFailed(err)
	}
	refreshToken, err := a.Jwt.GenerateRefreshToken(user.UserID.Hex())
	if err != nil {
		return nil, errors.TokenGenerateFailed(err)
	}
	return &common.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    a.service.Config.JWTConfig.TokenDuration.Seconds(),
		Meta: struct {
			UserID   string `json:"user_id"`
			Username string `json:"username"`
			Email    string `json:"email"`
			Role     string `json:"role"`
		}(struct {
			UserID   string
			Username string
			Email    string
			Role     string
		}{
			UserID:   user.UserID.Hex(),
			Username: user.Username,
			Email:    user.Email,
			Role:     user.Role,
		}),
	}, nil
}

func (a AuthServiceImpl) RefreshToken(ctx context.Context, refreshToken *string) (*common.RefreshTokenResponse, error) {
	userIDStr, err := a.Jwt.VerifyToken(*refreshToken)
	if err != nil {
		return nil, errors.InvalidToken(err) // TODO: Change error type
	}
	userID, err := primitive.ObjectIDFromHex(userIDStr)
	if err != nil {
		return nil, errors.InvalidToken(err) // TODO: Change error type
	}
	user, err := a.userDao.GetUserById(ctx, userID)
	if err != nil {
		return nil, errors.UserNotFound(err)
	}
	accessToken, err := a.Jwt.GenerateAccessToken(userID.Hex())
	if err != nil {
		return nil, errors.TokenGenerateFailed(err)
	}
	newRefreshToken, err := a.Jwt.GenerateRefreshToken(userID.Hex())
	if err != nil {
		return nil, errors.TokenGenerateFailed(err)
	}
	return &common.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
		ExpiresIn:    a.service.Config.JWTConfig.TokenDuration.Seconds(),
		Meta: struct {
			UserID   string `json:"user_id"`
			Username string `json:"username"`
			Email    string `json:"email"`
			Role     string `json:"role"`
		}(struct {
			UserID   string
			Username string
			Email    string
			Role     string
		}{
			UserID:   user.UserID.Hex(),
			Username: user.Username,
			Email:    user.Email,
			Role:     user.Role,
		}),
	}, nil
}

func (a AuthServiceImpl) Logout(ctx context.Context) error {
	_, err := primitive.ObjectIDFromHex(ctx.Value(config.KeyUserID).(string))
	if err != nil {
		return errors.InvalidToken(err)
	}
	// TODO: Implement logout, use redis to store token
	return nil
}

func (a AuthServiceImpl) ChangePassword(ctx context.Context, oldPassword, newPassword *string) error {
	userID, err := primitive.ObjectIDFromHex(ctx.Value(config.KeyUserID).(string))
	if err != nil {
		return errors.InvalidToken(err) // TODO: Change error type
	}
	user, err := a.userDao.GetUserById(ctx, userID)
	if err != nil {
		return errors.UserNotFound(err)
	}
	if !crypt.VerifyPassword(*oldPassword, user.Password) {
		return errors.PasswordWrong(err)
	}
	hashedPassword, err := crypt.PasswordHash(*newPassword)
	if err != nil {
		return errors.ServiceError(err)
	}
	err = a.userDao.UpdateUser(ctx, userID, nil, nil, nil, &hashedPassword, nil)
	if err != nil {
		return errors.MongoError(errors.WriteError(err))
	}
	return nil
}

func NewAuthService(s *service.Service, userDao dao.UserDao, loginLogDao dao.LoginLogDao, jwt *jwt.Jwt) AuthService {
	return &AuthServiceImpl{
		service:     s,
		userDao:     userDao,
		loginLogDao: loginLogDao,
		Jwt:         jwt,
	}
}
