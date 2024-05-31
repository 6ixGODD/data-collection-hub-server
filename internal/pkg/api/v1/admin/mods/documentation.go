package mods

import (
	"fmt"

	"data-collection-hub-server/internal/pkg/config"
	"data-collection-hub-server/internal/pkg/domain/vo"
	"data-collection-hub-server/internal/pkg/domain/vo/admin"
	adminservice "data-collection-hub-server/internal/pkg/service/admin/mods"
	sysservice "data-collection-hub-server/internal/pkg/service/sys/mods"
	"data-collection-hub-server/pkg/errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DocumentationApi struct {
	DocumentationService adminservice.DocumentationService
	LogsService          sysservice.LogsService
	Validator            *validator.Validate
}

func (d DocumentationApi) InsertDocumentation(c *fiber.Ctx) error {
	ctx := c.UserContext()
	req := new(admin.InsertDocumentationRequest)

	if err := c.BodyParser(req); err != nil {
		return errors.InvalidRequest(err)
	}

	if err := d.Validator.Struct(req); err != nil {
		return errors.InvalidParams(err) // Compare this line with the original one
	}

	documentationIDHex, err := d.DocumentationService.InsertDocumentation(ctx, req.Title, req.Content)
	var (
		userID, _  = primitive.ObjectIDFromHex(ctx.Value(config.UserIDKey).(string))
		ipAddr     = c.IP()
		userAgent  = c.Get(fiber.HeaderUserAgent)
		operation  = config.OperationTypeCreate
		entityType = config.EntityTypeDocumentation
	)
	if err != nil {
		var (
			description = fmt.Sprintf("Insert documentation failed: %s", err.Error())
			status      = config.OperationStatusFailure
		)
		_ = d.LogsService.CacheOperationLog(
			ctx, &userID, nil, &ipAddr, &userAgent, &operation, &entityType, &description, &status,
		)
		return err
	}

	var (
		documentationID, _ = primitive.ObjectIDFromHex(documentationIDHex)
		description        = fmt.Sprintf("Insert documentation: %s", documentationIDHex)
		status             = config.OperationStatusSuccess
	)

	_ = d.LogsService.CacheOperationLog(
		ctx, &userID, &documentationID, &ipAddr, &userAgent, &operation, &entityType, &description, &status,
	)

	return c.JSON(
		vo.Response{
			Code:    errors.CodeSuccess,
			Message: errors.MessageSuccess,
			Data:    nil,
		},
	)
}

func (d DocumentationApi) UpdateDocumentation(c *fiber.Ctx) error {
	ctx := c.UserContext()
	req := new(admin.UpdateDocumentationRequest)

	if err := c.BodyParser(req); err != nil {
		return errors.InvalidRequest(err)
	}

	if err := d.Validator.Struct(req); err != nil {
		return errors.InvalidParams(err) // Compare this line with the original one
	}
	documentationID, err := primitive.ObjectIDFromHex(*req.DocumentationID)
	if err != nil {
		return errors.InvalidRequest(err)
	}
	err = d.DocumentationService.UpdateDocumentation(ctx, &documentationID, req.Title, req.Content)
	var (
		userID, _  = primitive.ObjectIDFromHex(ctx.Value(config.UserIDKey).(string))
		ipAddr     = c.IP()
		userAgent  = c.Get(fiber.HeaderUserAgent)
		operation  = config.OperationTypeUpdate
		entityType = config.EntityTypeDocumentation
	)

	if err != nil {
		var (
			description = fmt.Sprintf("Update documentation failed: %s", err.Error())
			status      = config.OperationStatusFailure
		)
		_ = d.LogsService.CacheOperationLog(
			ctx, &userID, &documentationID, &ipAddr, &userAgent, &operation, &entityType, &description, &status,
		)
		return err
	}

	var (
		description = fmt.Sprintf("Update documentation: %s", *req.DocumentationID)
		status      = config.OperationStatusSuccess
	)
	_ = d.LogsService.CacheOperationLog(
		ctx, &userID, &documentationID, &ipAddr, &userAgent, &operation, &entityType, &description, &status,
	)
	return c.JSON(
		vo.Response{
			Code:    errors.CodeSuccess,
			Message: errors.MessageSuccess,
			Data:    nil,
		},
	)
}

func (d DocumentationApi) DeleteDocumentation(c *fiber.Ctx) error {
	ctx := c.UserContext()
	req := new(admin.DeleteDocumentationRequest)

	if err := c.BodyParser(req); err != nil {
		return errors.InvalidRequest(err)
	}

	if err := d.Validator.Struct(req); err != nil {
		return errors.InvalidParams(err) // Compare this line with the original one
	}
	documentationID, err := primitive.ObjectIDFromHex(*req.DocumentationID)
	if err != nil {
		return errors.InvalidRequest(err)
	}
	err = d.DocumentationService.DeleteDocumentation(ctx, &documentationID)
	var (
		userID, _  = primitive.ObjectIDFromHex(ctx.Value(config.UserIDKey).(string))
		ipAddr     = c.IP()
		userAgent  = c.Get(fiber.HeaderUserAgent)
		operation  = config.OperationTypeDelete
		entityType = config.EntityTypeDocumentation
	)
	if err != nil {
		var (
			description = fmt.Sprintf("Delete documentation failed: %s", err.Error())
			status      = config.OperationStatusFailure
		)
		_ = d.LogsService.CacheOperationLog(
			ctx, &userID, &documentationID, &ipAddr, &userAgent, &operation, &entityType, &description, &status,
		)
		return err
	}

	var (
		description = fmt.Sprintf("Delete documentation: %s", *req.DocumentationID)
		status      = config.OperationStatusSuccess
	)
	_ = d.LogsService.CacheOperationLog(
		ctx, &userID, &documentationID, &ipAddr, &userAgent, &operation, &entityType, &description, &status,
	)
	return c.JSON(
		vo.Response{
			Code:    errors.CodeSuccess,
			Message: errors.MessageSuccess,
			Data:    nil,
		},
	)
}
