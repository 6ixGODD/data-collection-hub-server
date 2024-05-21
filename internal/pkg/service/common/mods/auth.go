package mods

import (
	"context"

	"data-collection-hub-server/internal/pkg/config"
	dao "data-collection-hub-server/internal/pkg/dao/mods"
	"data-collection-hub-server/internal/pkg/domain/vo/common"
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

type AuthDOImpl struct {
	core        *service.Core
	userDao     dao.UserDao
	loginLogDao dao.LoginLogDao
	Jwt         *jwt.Jwt
}

func NewAuthDO(core *service.Core, userDao dao.UserDao, loginLogDao dao.LoginLogDao, jwt *jwt.Jwt) AuthService {
	return &AuthDOImpl{
		core:        core,
		userDao:     userDao,
		loginLogDao: loginLogDao,
		Jwt:         jwt,
	}
}

func (a AuthDOImpl) Login(ctx context.Context, email, password *string) (*common.LoginResponse, error) {
	user, err := a.userDao.GetUserByEmail(ctx, *email)
	if err != nil {
		return nil, errors.DBError(errors.ReadError(err))
	}
	if !crypt.Compare(*password, user.Password) {
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
		ExpiresIn:    a.core.Config.JWTConfig.TokenDuration.Seconds(),
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

func (a AuthDOImpl) RefreshToken(ctx context.Context, refreshToken *string) (*common.RefreshTokenResponse, error) {
	userIDStr, err := a.Jwt.VerifyToken(*refreshToken)
	if err != nil {
		return nil, errors.InvalidToken(err) // TODO: Change error type
	}
	userID, err := primitive.ObjectIDFromHex(userIDStr)
	if err != nil {
		return nil, errors.InvalidToken(err) // TODO: Change error type
	}
	user, err := a.userDao.GetUserByID(ctx, userID)
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
		ExpiresIn:    a.core.Config.JWTConfig.TokenDuration.Seconds(),
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

func (a AuthDOImpl) Logout(ctx context.Context) error {
	_, err := primitive.ObjectIDFromHex(ctx.Value(config.UserIDKey).(string))
	if err != nil {
		return errors.InvalidToken(err)
	}
	// TODO: Implement logout, use redis to store token
	return nil
}

func (a AuthDOImpl) ChangePassword(ctx context.Context, oldPassword, newPassword *string) error {
	userID, err := primitive.ObjectIDFromHex(ctx.Value(config.UserIDKey).(string))
	if err != nil {
		return errors.InvalidToken(err) // TODO: Change error type
	}
	user, err := a.userDao.GetUserByID(ctx, userID)
	if err != nil {
		return errors.UserNotFound(err)
	}
	if !crypt.Compare(*oldPassword, user.Password) {
		return errors.PasswordWrong(err)
	}
	hashedPassword, err := crypt.Hash(*newPassword)
	if err != nil {
		return errors.ServiceError(err)
	}
	err = a.userDao.UpdateUser(ctx, userID, nil, nil, nil, &hashedPassword, nil)
	if err != nil {
		return errors.DBError(errors.WriteError(err))
	}
	return nil
}
