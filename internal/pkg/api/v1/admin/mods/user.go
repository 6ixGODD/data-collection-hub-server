package mods

import (
	"time"

	"data-collection-hub-server/internal/pkg/domain/vo"
	"data-collection-hub-server/internal/pkg/domain/vo/admin"
	adminservice "data-collection-hub-server/internal/pkg/service/admin/mods"
	"data-collection-hub-server/pkg/errors"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserApi struct {
	UserService adminservice.UserService
}

func (u *UserApi) InsertUser(c *fiber.Ctx) error {
	req := new(admin.InsertUserRequest)

	if err := c.BodyParser(req); err != nil {
		return errors.InvalidRequest(err)
	}

	err := u.UserService.InsertUser(c.UserContext(), req.Username, req.Email, req.Password, req.Role, req.Organization)
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

func (u *UserApi) GetUser(c *fiber.Ctx) error {
	req := new(admin.GetUserRequest)

	if err := c.QueryParser(req); err != nil {
		return errors.InvalidRequest(err)
	}

	userID, err := primitive.ObjectIDFromHex(*req.UserID)
	if err != nil {
		return errors.InvalidRequest(err)
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

func (u *UserApi) GetUserList(c *fiber.Ctx) error {
	req := new(admin.GetUserListRequest)

	if err := c.QueryParser(req); err != nil {
		return errors.InvalidRequest(err)
	}

	var (
		lastLoginTimeStart, lastLoginTimeEnd,
		createdTimeStart, createdTimeEnd *time.Time
		err error
	)

	if req.LastLoginTimeStart != nil {
		*lastLoginTimeStart, err = time.Parse(time.RFC3339, *req.LastLoginTimeStart)
		if err != nil {
			return errors.InvalidRequest(err)
		}
	}
	if req.LastLoginTimeEnd != nil {
		*lastLoginTimeEnd, err = time.Parse(time.RFC3339, *req.LastLoginTimeEnd)
		if err != nil {
			return errors.InvalidRequest(err)
		}
	}
	if req.CreateTimeStart != nil {
		*createdTimeStart, err = time.Parse(time.RFC3339, *req.CreateTimeStart)
		if err != nil {
			return errors.InvalidRequest(err)
		}
	}
	if req.CreateTimeEnd != nil {
		*createdTimeEnd, err = time.Parse(time.RFC3339, *req.CreateTimeEnd)
		if err != nil {
			return errors.InvalidRequest(err)
		}
	}

	resp, err := u.UserService.GetUserList(
		c.UserContext(), req.Page, req.PageSize, req.Desc, req.Role, lastLoginTimeStart, lastLoginTimeEnd,
		createdTimeStart, createdTimeEnd, req.Query,
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

func (u *UserApi) UpdateUser(c *fiber.Ctx) error {
	req := new(admin.UpdateUserRequest)

	if err := c.BodyParser(req); err != nil {
		return errors.InvalidRequest(err)
	}

	userID, err := primitive.ObjectIDFromHex(*req.UserID)
	if err != nil {
		return errors.InvalidRequest(err)
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

func (u *UserApi) DeleteUser(c *fiber.Ctx) error {
	req := new(admin.DeleteUserRequest)

	if err := c.QueryParser(req); err != nil {
		return errors.InvalidRequest(err)
	}

	userID, err := primitive.ObjectIDFromHex(*req.UserID)
	if err != nil {
		return errors.InvalidRequest(err)
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

func (u *UserApi) ChangeUserPassword(c *fiber.Ctx) error {
	req := new(admin.ChangeUserPasswordRequest)

	if err := c.BodyParser(req); err != nil {
		return errors.InvalidRequest(err)
	}

	userID, err := primitive.ObjectIDFromHex(*req.UserID)
	if err != nil {
		return errors.InvalidRequest(err)
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
