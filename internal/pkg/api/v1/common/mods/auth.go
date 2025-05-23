package mods

import (
	"fmt"

	"data-collection-hub-server/internal/pkg/domain/vo"
	"data-collection-hub-server/internal/pkg/domain/vo/common"
	commonservice "data-collection-hub-server/internal/pkg/service/common/mods"
	sysservice "data-collection-hub-server/internal/pkg/service/sys/mods"
	"data-collection-hub-server/pkg/errors"
	"data-collection-hub-server/pkg/utils/check"
	utils "data-collection-hub-server/pkg/utils/common"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthApi struct {
	AuthService commonservice.AuthService
	LogsService sysservice.LogsService
	Validator   *validator.Validate
}

// Login logs in the user and returns a token.
//
//	@description	Log in the user and return a token.
//	@id				common-login
//	@summary		login
//	@tags			Auth API
//	@accept			json
//	@produce		json
//	@param			common.LoginRequest	body		common.LoginRequest						true	"Login request"
//	@success		200					{object}	vo.Response{data=common.LoginResponse}	"Success"
//	@failure		400					{object}	vo.Response{data=nil}					"Invalid request"
//	@failure		401					{object}	vo.Response{data=nil}					"Unauthorized"
//	@failure		500					{object}	vo.Response{data=nil}					"Internal server error"
//	@router			/login																																																																																																																				[post]
func (a *AuthApi) Login(c *fiber.Ctx) error {
	ctx := c.UserContext()
	req := new(common.LoginRequest)

	if err := c.BodyParser(req); err != nil {
		return errors.InvalidRequest(fmt.Errorf("failed to parse request"))
	}
	if errs := a.Validator.Struct(req); errs != nil {
		return errors.InvalidRequest(utils.FormatValidateError(errs))
	}

	resp, err := a.AuthService.Login(ctx, req.Email, req.Password)
	if err != nil {
		return err
	}

	userID, _ := primitive.ObjectIDFromHex(resp.Meta.UserID)
	ipAddr := c.IP()
	userAgent := c.Get(fiber.HeaderUserAgent)
	_ = a.LogsService.CacheLoginLog(ctx, &userID, &ipAddr, &userAgent)
	return c.JSON(
		vo.Response{
			Code:    errors.CodeSuccess,
			Message: errors.MessageSuccess,
			Data:    resp,
		},
	)
}

// Logout logs out the user.
//
//	@description	Log out the user.
//	@id				common-logout
//	@summary		logout
//	@tags			Auth API
//	@accept			json
//	@produce		json
//	@security		Bearer
//	@success		200		{object}	vo.Response{data=nil}	"Success"
//	@failure		401		{object}	vo.Response{data=nil}	"Unauthorized"
//	@failure		500		{object}	vo.Response{data=nil}	"Internal server error"
//	@router			/logout	[get]
func (a *AuthApi) Logout(c *fiber.Ctx) error {
	token := c.Get(fiber.HeaderAuthorization)
	if token == "" {
		return errors.TokenMissed(fmt.Errorf("token is missed"))
	}
	if !check.IsBearerToken(token) {
		return errors.TokenInvalid(fmt.Errorf("token is invalid"))
	}
	token = token[7:]
	err := a.AuthService.Logout(c.UserContext(), &token)
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

// RefreshToken refreshes the user's token.
//
//	@description	Refresh the user's token.
//	@id				common-refresh-token
//	@summary		refresh token
//	@tags			Auth API
//	@accept			json
//	@produce		json
//	@param			common.RefreshTokenRequest	body	common.RefreshTokenRequest	true	"Refresh token request"
//	@security		Bearer
//	@success		200				{object}	vo.Response{data=common.RefreshTokenResponse}	"Success"
//	@failure		400				{object}	vo.Response{data=nil}							"Invalid request"
//	@failure		401				{object}	vo.Response{data=nil}							"Unauthorized"
//	@failure		500				{object}	vo.Response{data=nil}							"Internal server error"
//	@router			/refresh-token	[post]
func (a *AuthApi) RefreshToken(c *fiber.Ctx) error {
	req := new(common.RefreshTokenRequest)

	if err := c.BodyParser(req); err != nil {
		return errors.InvalidRequest(fmt.Errorf("failed to parse request"))
	}
	if errs := a.Validator.Struct(req); errs != nil {
		return errors.InvalidRequest(utils.FormatValidateError(errs))
	}

	resp, err := a.AuthService.RefreshToken(c.UserContext(), req.RefreshToken)
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

// ChangePassword changes the user's password.
//
//	@description	Change the user's password.
//	@id				common-change-password
//	@summary		change password
//	@tags			Auth API
//	@accept			json
//	@produce		json
//	@param			common.ChangePasswordRequest	body	common.ChangePasswordRequest	true	"Change password request"
//	@security		Bearer
//	@success		200					{object}	vo.Response{data=nil}	"Success"
//	@failure		400					{object}	vo.Response{data=nil}	"Invalid request"
//	@failure		401					{object}	vo.Response{data=nil}	"Unauthorized"
//	@failure		500					{object}	vo.Response{data=nil}	"Internal server error"
//	@router			/change-password	[post]
func (a *AuthApi) ChangePassword(c *fiber.Ctx) error {
	req := new(common.ChangePasswordRequest)

	if err := c.BodyParser(req); err != nil {
		return errors.InvalidRequest(fmt.Errorf("failed to parse request"))
	}
	if errs := a.Validator.Struct(req); errs != nil {
		return errors.InvalidRequest(utils.FormatValidateError(errs))
	}

	err := a.AuthService.ChangePassword(c.UserContext(), req.OldPassword, req.NewPassword)
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
