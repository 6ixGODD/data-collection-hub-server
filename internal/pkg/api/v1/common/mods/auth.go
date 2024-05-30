package mods

import (
	"fmt"

	"data-collection-hub-server/internal/pkg/domain/vo"
	"data-collection-hub-server/internal/pkg/domain/vo/common"
	commonservice "data-collection-hub-server/internal/pkg/service/common/mods"
	sysservice "data-collection-hub-server/internal/pkg/service/sys/mods"
	"data-collection-hub-server/pkg/errors"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthApi struct {
	AuthService commonservice.AuthService
	LogsService sysservice.LogsService
}

func (api *AuthApi) Login(c *fiber.Ctx) error {
	ctx := c.UserContext()
	req := new(common.LoginRequest)

	if err := c.BodyParser(req); err != nil {
		return errors.InvalidRequest(err)
	}

	resp, err := api.AuthService.Login(ctx, req.Email, req.Password)
	if err != nil {
		return err
	}

	userID, _ := primitive.ObjectIDFromHex(resp.Meta.UserID)
	ipAddr := c.IP()
	userAgent := c.Get(fiber.HeaderUserAgent)
	_ = api.LogsService.CacheLoginLog(ctx, &userID, &ipAddr, &userAgent)
	return c.JSON(
		vo.Response{
			Code:    errors.CodeSuccess,
			Message: errors.MessageSuccess,
			Data:    resp,
		},
	)
}

func (api *AuthApi) Logout(c *fiber.Ctx) error {
	token := c.Get(fiber.HeaderAuthorization)
	if token == "" {
		return errors.InvalidRequest(errors.TokenMissed(fmt.Errorf("token is missed")))
	}
	err := api.AuthService.Logout(c.UserContext(), &token)
	if err != nil {
		return err
	}

	return c.JSON(
		vo.Response{
			Code:    errors.CodeSuccess,
			Message: errors.MessageSuccess,
			Data:    nil,
		},
	)
}

func (api *AuthApi) RefreshToken(c *fiber.Ctx) error {
	req := new(common.RefreshTokenRequest)

	if err := c.BodyParser(req); err != nil {
		return errors.InvalidRequest(err)
	}

	resp, err := api.AuthService.RefreshToken(c.UserContext(), req.RefreshToken)
	if err != nil {
		return err
	}

	return c.JSON(
		vo.Response{
			Code:    errors.CodeSuccess,
			Message: errors.MessageSuccess,
			Data:    resp,
		},
	)
}

func (api *AuthApi) ChangePassword(c *fiber.Ctx) error {
	req := new(common.ChangePasswordRequest)

	if err := c.BodyParser(req); err != nil {
		return errors.InvalidRequest(err)
	}

	err := api.AuthService.ChangePassword(c.UserContext(), req.OldPassword, req.NewPassword)
	if err != nil {
		return err
	}

	return c.JSON(
		vo.Response{
			Code:    errors.CodeSuccess,
			Message: errors.MessageSuccess,
			Data:    nil,
		},
	)
}
