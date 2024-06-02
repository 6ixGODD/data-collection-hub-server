package mods

import (
	"fmt"
	"time"

	"data-collection-hub-server/internal/pkg/domain/vo"
	"data-collection-hub-server/internal/pkg/domain/vo/admin"
	adminservice "data-collection-hub-server/internal/pkg/service/admin/mods"
	"data-collection-hub-server/pkg/errors"
	"data-collection-hub-server/pkg/utils/common"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserApi struct {
	UserService adminservice.UserService
	Validator   *validator.Validate
}

// InsertUser inserts a new user.
//
//	@description	Insert a new user.
//	@id				admin-insert-user
//	@summary		insert user
//	@tags			Admin API
//	@accept			json
//	@produce		json
//	@param			admin.InsertUserRequest	body	admin.InsertUserRequest	true	"Insert user request"
//	@security		Bearer
//	@success		200			{object}	vo.Response{data=nil}	"Success"
//	@failure		400			{object}	vo.Response{data=nil}	"Invalid request"
//	@failure		401			{object}	vo.Response{data=nil}	"Unauthorized"
//	@failure		403			{object}	vo.Response{data=nil}	"Forbidden"
//	@failure		500			{object}	vo.Response{data=nil}	"Internal server error"
//	@router			/admin/user																																																																		[post]
func (u *UserApi) InsertUser(c *fiber.Ctx) error {
	req := new(admin.InsertUserRequest)

	if err := c.BodyParser(req); err != nil {
		return errors.InvalidRequest(fmt.Errorf("failed to parse request"))
	}
	if errs := u.Validator.Struct(req); errs != nil {
		return errors.InvalidRequest(common.FormatValidateError(errs))
	}

	_, err := u.UserService.InsertUser(
		c.UserContext(), req.Username, req.Email, req.Password, req.Organization,
	)
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

// GetUser returns the user by ID.
//
//	@description	Get the user by ID.
//	@id				admin-get-user
//	@summary		get user by ID
//	@tags			Admin API
//	@accept			json
//	@produce		json
//	@param			admin.GetUserRequest	query	admin.GetUserRequest	true	"Get user request"
//	@security		Bearer
//	@success		200			{object}	vo.Response{data=admin.GetUserResponse}	"Success"
//	@failure		400			{object}	vo.Response{data=nil}					"Invalid request"
//	@failure		401			{object}	vo.Response{data=nil}					"Unauthorized"
//	@failure		403			{object}	vo.Response{data=nil}					"Forbidden"
//	@failure		404			{object}	vo.Response{data=nil}					"User not found"
//	@failure		500			{object}	vo.Response{data=nil}					"Internal server error"
//	@router			/admin/user																																																																																													[get]
func (u *UserApi) GetUser(c *fiber.Ctx) error {
	req := new(admin.GetUserRequest)

	if err := c.QueryParser(req); err != nil {
		return errors.InvalidRequest(fmt.Errorf("failed to parse request"))
	}
	if errs := u.Validator.Struct(req); errs != nil {
		return errors.InvalidRequest(common.FormatValidateError(errs))
	}

	userID, err := primitive.ObjectIDFromHex(*req.UserID)
	if err != nil {
		return errors.InvalidRequest(fmt.Errorf("invalid user id"))
	}

	resp, err := u.UserService.GetUser(c.UserContext(), &userID)
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

// GetUserList returns the list of users based on the query parameters.
//
//	@description	Get the list of users based on the query parameters.
//	@id				admin-get-user-list
//	@summary		get user list
//	@tags			Admin API
//	@accept			json
//	@produce		json
//	@param			admin.GetUserListRequest	query	admin.GetUserListRequest	true	"Get user list request"
//	@security		Bearer
//	@success		200					{object}	vo.Response{data=admin.GetUserListResponse}	"Success"
//	@failure		400					{object}	vo.Response{data=nil}						"Invalid request"
//	@failure		401					{object}	vo.Response{data=nil}						"Unauthorized"
//	@failure		403					{object}	vo.Response{data=nil}						"Forbidden"
//	@failure		500					{object}	vo.Response{data=nil}						"Internal server error"
//	@router			/admin/user/list																																																																																																				[get]
func (u *UserApi) GetUserList(c *fiber.Ctx) error {
	req := new(admin.GetUserListRequest)

	if err := c.QueryParser(req); err != nil {
		return errors.InvalidRequest(fmt.Errorf("failed to parse request"))
	}
	if errs := u.Validator.Struct(req); errs != nil {
		return errors.InvalidRequest(common.FormatValidateError(errs))
	}

	var (
		lastLoginStartTime, lastLoginEndTime,
		createdStartTime, createdEndTime time.Time
		lastLoginStartTimePtr, lastLoginEndTimePtr,
		createdStartTimePtr, createdEndTimePtr *time.Time
		err error
	)

	if req.LastLoginTimeStart != nil && req.LastLoginTimeEnd != nil {
		lastLoginStartTime, err = time.Parse(time.RFC3339, *req.LastLoginTimeStart)
		if err != nil {
			return errors.InvalidRequest(
				fmt.Errorf(
					"invalid last login time start %s (should be in RFC3339 format)", *req.LastLoginTimeStart,
				),
			)
		}
		lastLoginEndTime, err = time.Parse(time.RFC3339, *req.LastLoginTimeEnd)
		if err != nil {
			return errors.InvalidRequest(
				fmt.Errorf(
					"invalid last login time end %s (should be in RFC3339 format)", *req.LastLoginTimeEnd,
				),
			)
		}
		lastLoginStartTimePtr = &lastLoginStartTime
		lastLoginEndTimePtr = &lastLoginEndTime
	}
	if req.CreateTimeStart != nil && req.CreateTimeEnd != nil {
		createdStartTime, err = time.Parse(time.RFC3339, *req.CreateTimeStart)
		if err != nil {
			return errors.InvalidRequest(
				fmt.Errorf(
					"invalid create time start %s (should be in RFC3339 format)", *req.CreateTimeStart,
				),
			)
		}
		createdEndTime, err = time.Parse(time.RFC3339, *req.CreateTimeEnd)
		if err != nil {
			return errors.InvalidRequest(
				fmt.Errorf(
					"invalid create time end %s (should be in RFC3339 format)", *req.CreateTimeEnd,
				),
			)
		}
		createdStartTimePtr = &createdStartTime
		createdEndTimePtr = &createdEndTime
	}

	resp, err := u.UserService.GetUserList(
		c.UserContext(), req.Page, req.PageSize, req.Desc, req.Role, lastLoginStartTimePtr, lastLoginEndTimePtr,
		createdStartTimePtr, createdEndTimePtr, req.Query,
	)
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

// UpdateUser updates the user.
//
//	@description	Update the user.
//	@id				admin-update-user
//	@summary		update user
//	@tags			Admin API
//	@accept			json
//	@produce		json
//	@param			admin.UpdateUserRequest	body	admin.UpdateUserRequest	true	"Update user request"
//	@security		Bearer
//	@success		200			{object}	vo.Response{data=nil}	"Success"
//	@failure		400			{object}	vo.Response{data=nil}	"Invalid request"
//	@failure		401			{object}	vo.Response{data=nil}	"Unauthorized"
//	@failure		403			{object}	vo.Response{data=nil}	"Forbidden"
//	@failure		500			{object}	vo.Response{data=nil}	"Internal server error"
//	@router			/admin/user																																																																		[put]
func (u *UserApi) UpdateUser(c *fiber.Ctx) error {
	req := new(admin.UpdateUserRequest)

	if err := c.BodyParser(req); err != nil {
		return errors.InvalidRequest(fmt.Errorf("failed to parse request"))
	}
	if errs := u.Validator.Struct(req); errs != nil {
		return errors.InvalidRequest(common.FormatValidateError(errs))
	}

	userID, err := primitive.ObjectIDFromHex(*req.UserID)
	if err != nil {
		return errors.InvalidRequest(fmt.Errorf("invalid user id"))
	}

	err = u.UserService.UpdateUser(c.UserContext(), &userID, req.Username, req.Email, req.Role, req.Organization)
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

// DeleteUser deletes a user by user ID.
//
//	@description	Delete the user by ID.
//	@id				admin-delete-user
//	@summary		delete user by ID
//	@tags			Admin API
//	@accept			json
//	@produce		json
//	@param			admin.DeleteUserRequest	query	admin.DeleteUserRequest	true	"Delete user request"
//	@security		Bearer
//	@success		200			{object}	vo.Response{data=nil}	"Success"
//	@failure		400			{object}	vo.Response{data=nil}	"Invalid request"
//	@failure		401			{object}	vo.Response{data=nil}	"Unauthorized"
//	@failure		403			{object}	vo.Response{data=nil}	"Forbidden"
//	@failure		500			{object}	vo.Response{data=nil}	"Internal server error"
//	@router			/admin/user																																																																					[delete]
func (u *UserApi) DeleteUser(c *fiber.Ctx) error {
	req := new(admin.DeleteUserRequest)

	if err := c.QueryParser(req); err != nil {
		return errors.InvalidRequest(fmt.Errorf("failed to parse request"))
	}
	if errs := u.Validator.Struct(req); errs != nil {
		return errors.InvalidRequest(common.FormatValidateError(errs))
	}

	userID, err := primitive.ObjectIDFromHex(*req.UserID)
	if err != nil {
		return errors.InvalidRequest(fmt.Errorf("invalid user id"))
	}
	err = u.UserService.DeleteUser(c.UserContext(), &userID)
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

// ChangeUserPassword changes a user's password.
//
//	@description	Change the user's password.
//	@id				admin-change-user-password
//	@summary		change user password
//	@tags			Admin API
//	@accept			json
//	@produce		json
//	@param			admin.ChangeUserPasswordRequest	body	admin.ChangeUserPasswordRequest	true	"Change user password request"
//	@security		Bearer
//	@success		200						{object}	vo.Response{data=nil}	"Success"
//	@failure		400						{object}	vo.Response{data=nil}	"Invalid request"
//	@failure		401						{object}	vo.Response{data=nil}	"Unauthorized"
//	@failure		403						{object}	vo.Response{data=nil}	"Forbidden"
//	@failure		500						{object}	vo.Response{data=nil}	"Internal server error"
//	@router			/admin/user/password																																																																[put]
func (u *UserApi) ChangeUserPassword(c *fiber.Ctx) error {
	req := new(admin.ChangeUserPasswordRequest)

	if err := c.BodyParser(req); err != nil {
		return errors.InvalidRequest(fmt.Errorf("failed to parse request"))
	}
	if errs := u.Validator.Struct(req); errs != nil {
		return errors.InvalidRequest(common.FormatValidateError(errs))
	}

	userID, err := primitive.ObjectIDFromHex(*req.UserID)
	if err != nil {
		return errors.InvalidRequest(fmt.Errorf("invalid user id"))
	}
	err = u.UserService.ChangeUserPassword(c.UserContext(), &userID, req.NewPassword)
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
