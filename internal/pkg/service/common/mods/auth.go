package mods

import (
	"context"
	"fmt"

	"data-collection-hub-server/internal/pkg/config"
	"data-collection-hub-server/internal/pkg/dao"
	daos "data-collection-hub-server/internal/pkg/dao/mods"
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
	Logout(ctx context.Context, accessToken *string) error
	ChangePassword(ctx context.Context, oldPassword, newPassword *string) error
}

type authServiceImpl struct {
	core    *service.Core
	cache   *dao.Cache
	userDao daos.UserDao
	jwt     *jwt.Jwt
}

func NewAuthService(core *service.Core, userDao daos.UserDao, cache *dao.Cache, jwt *jwt.Jwt) AuthService {
	return &authServiceImpl{
		core:    core,
		cache:   cache,
		userDao: userDao,
		jwt:     jwt,
	}
}

func (a authServiceImpl) Login(ctx context.Context, email, password *string) (*common.LoginResponse, error) {
	user, err := a.userDao.GetUserByEmail(ctx, *email)
	if err != nil {
		return nil, errors.DBError(errors.ReadError(err))
	}
	if !crypt.Compare(*password, user.Password) {
		return nil, errors.PasswordWrong(err)
	}
	accessToken, err := a.jwt.GenerateAccessToken(user.UserID.Hex())
	if err != nil {
		return nil, errors.TokenGenerateFailed(err)
	}
	refreshToken, err := a.jwt.GenerateRefreshToken(user.UserID.Hex())
	if err != nil {
		return nil, errors.TokenGenerateFailed(err)
	}
	err = a.userDao.UpdateUserLastLogin(ctx, user.UserID)
	if err != nil {
		return nil, errors.DBError(errors.WriteError(err))
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

func (a authServiceImpl) RefreshToken(ctx context.Context, refreshToken *string) (*common.RefreshTokenResponse, error) {
	userIDHex, err := a.jwt.VerifyToken(*refreshToken)
	if err != nil {
		return nil, errors.InvalidToken(err) // TODO: Change error type
	}
	userID, err := primitive.ObjectIDFromHex(userIDHex)
	if err != nil {
		return nil, errors.InvalidToken(err) // TODO: Change error type
	}
	user, err := a.userDao.GetUserByID(ctx, userID)
	if err != nil {
		return nil, errors.UserNotFound(err)
	}
	accessToken, err := a.jwt.GenerateAccessToken(userID.Hex())
	if err != nil {
		return nil, errors.TokenGenerateFailed(err)
	}
	newRefreshToken, err := a.jwt.GenerateRefreshToken(userID.Hex())
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

func (a authServiceImpl) Logout(ctx context.Context, accessToken *string) error {
	if err := a.cache.Set(
		ctx, fmt.Sprintf("%s:%s", config.TokenBlacklistCachePrefix, crypt.MD5(*accessToken)), config.CacheTrue,
		&a.core.Config.CacheConfig.TokenBlacklistTTL,
	); err != nil {
		return errors.CacheError(err)
	}
	return nil
}

func (a authServiceImpl) ChangePassword(ctx context.Context, oldPassword, newPassword *string) error {
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
	if err = a.userDao.UpdateUser(ctx, userID, nil, nil, nil, &hashedPassword, nil); err != nil {
		return errors.DBError(errors.WriteError(err))
	}
	return nil
}
