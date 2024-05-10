package mods

import (
	"time"

	"data-collection-hub-server/internal/pkg/schema"
	"data-collection-hub-server/internal/pkg/schema/admin"
	adminservice "data-collection-hub-server/internal/pkg/service/admin/mods"
	"data-collection-hub-server/pkg/errors"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserApi struct {
	adminservice.UserService
}

func NewUserApi(userService adminservice.UserService) UserApi {
	return UserApi{userService}
}

func (u *UserApi) InsertUser(c *fiber.Ctx) error {
	req := new(admin.InsertUserRequest)

	if err := c.BodyParser(req); err != nil {
		return errors.InvalidRequest(err)
	}

	err := u.UserService.InsertUser(c.Context(), req.Username, req.Email, req.Password, req.Role, req.Organization)
	if err != nil {
		return err
	}

	return c.JSON(
		schema.Response{
			Code:    errors.CodeSuccess,
			Message: errors.MessageSuccess,
			Data:    nil,
		},
	)
}

func (u *UserApi) GetUser(c *fiber.Ctx) error {
	req := new(admin.GetUserRequest)

	if err := c.QueryParser(req); err != nil {
		return errors.InvalidRequest(err)
	}

	userID, err := primitive.ObjectIDFromHex(*req.UserID)
	if err != nil {
		return errors.InvalidRequest(err)
	}

	resp, err := u.UserService.GetUser(c.Context(), &userID)
	if err != nil {
		return err
	}

	return c.JSON(
		schema.Response{
			Code:    errors.CodeSuccess,
			Message: errors.MessageSuccess,
			Data:    resp,
		},
	)
}

func (u *UserApi) GetUserList(c *fiber.Ctx) error {
	req := new(admin.GetUserListRequest)

	if err := c.QueryParser(req); err != nil {
		return errors.InvalidRequest(err)
	}

	var (
		page                                                         *int
		role                                                         *string
		lastLoginBefore, lastLoginAfter, createdBefore, createdAfter *time.Time
		err                                                          error
	)

	if req.Page != nil {
		page = req.Page
	}
	if req.Role != nil {
		role = req.Role
	}
	if req.LastLoginBefore != nil {
		*lastLoginBefore, err = time.Parse(time.RFC3339, *req.LastLoginBefore)
		if err != nil {
			return errors.InvalidRequest(err)
		}
	}
	if req.LastLoginAfter != nil {
		*lastLoginAfter, err = time.Parse(time.RFC3339, *req.LastLoginAfter)
		if err != nil {
			return errors.InvalidRequest(err)
		}
	}
	if req.CreatedBefore != nil {
		*createdBefore, err = time.Parse(time.RFC3339, *req.CreatedBefore)
		if err != nil {
			return errors.InvalidRequest(err)
		}
	}
	if req.CreatedAfter != nil {
		*createdAfter, err = time.Parse(time.RFC3339, *req.CreatedAfter)
		if err != nil {
			return errors.InvalidRequest(err)
		}
	}

	resp, err := u.UserService.GetUserList(
		c.Context(), page, role, lastLoginBefore, lastLoginAfter, createdBefore, createdAfter,
	)
	if err != nil {
		return err
	}

	return c.JSON(
		schema.Response{
			Code:    errors.CodeSuccess,
			Message: errors.MessageSuccess,
			Data:    resp,
		},
	)
}

func (u *UserApi) UpdateUser(c *fiber.Ctx) error {
	req := new(admin.UpdateUserRequest)

	if err := c.BodyParser(req); err != nil {
		return errors.InvalidRequest(err)
	}

	userID, err := primitive.ObjectIDFromHex(*req.UserID)
	if err != nil {
		return errors.InvalidRequest(err)
	}

	err = u.UserService.UpdateUser(c.Context(), &userID, req.Username, req.Email, req.Role, req.Organization)
	if err != nil {
		return err
	}

	return c.JSON(
		schema.Response{
			Code:    errors.CodeSuccess,
			Message: errors.MessageSuccess,
			Data:    nil,
		},
	)
}

func (u *UserApi) DeleteUser(c *fiber.Ctx) error {
	req := new(admin.DeleteUserRequest)

	if err := c.QueryParser(req); err != nil {
		return errors.InvalidRequest(err)
	}

	userID, err := primitive.ObjectIDFromHex(*req.UserID)
	if err != nil {
		return errors.InvalidRequest(err)
	}
	err = u.UserService.DeleteUser(c.Context(), &userID)
	if err != nil {
		return err
	}

	return c.JSON(
		schema.Response{
			Code:    errors.CodeSuccess,
			Message: errors.MessageSuccess,
			Data:    nil,
		},
	)
}

func (u *UserApi) ChangeUserPassword(c *fiber.Ctx) error {
	req := new(admin.ChangeUserPasswordRequest)

	if err := c.BodyParser(req); err != nil {
		return errors.InvalidRequest(err)
	}

	userID, err := primitive.ObjectIDFromHex(*req.UserID)
	if err != nil {
		return errors.InvalidRequest(err)
	}
	err = u.UserService.ChangeUserPassword(c.Context(), &userID, req.NewPassword)
	if err != nil {
		return err
	}

	return c.JSON(
		schema.Response{
			Code:    errors.CodeSuccess,
			Message: errors.MessageSuccess,
			Data:    nil,
		},
	)
}
