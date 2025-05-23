package mods

import (
	"fmt"

	"data-collection-hub-server/internal/pkg/config"
	"data-collection-hub-server/internal/pkg/domain/vo"
	"data-collection-hub-server/internal/pkg/domain/vo/admin"
	adminservice "data-collection-hub-server/internal/pkg/service/admin/mods"
	sysservice "data-collection-hub-server/internal/pkg/service/sys/mods"
	"data-collection-hub-server/pkg/errors"
	"data-collection-hub-server/pkg/utils/common"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DocumentationApi struct {
	DocumentationService adminservice.DocumentationService
	LogsService          sysservice.LogsService
	Validator            *validator.Validate
}

// InsertDocumentation Insert a new documentation.
//
//	@description	Insert a new documentation.
//	@id				admin-insert-documentation
//	@summary		insert documentation
//	@tags			Admin API
//	@accept			json
//	@produce		json
//	@param			admin.InsertDocumentationRequest	body	admin.InsertDocumentationRequest	true	"Insert documentation request"
//	@security		Bearer
//	@success		200						{object}	vo.Response{data=nil}	"Success"
//	@failure		400						{object}	vo.Response{data=nil}	"Invalid request"
//	@failure		401						{object}	vo.Response{data=nil}	"Unauthorized"
//	@failure		403						{object}	vo.Response{data=nil}	"Forbidden"
//	@failure		500						{object}	vo.Response{data=nil}	"Internal server error"
//	@router			/admin/documentation	[post]
func (d DocumentationApi) InsertDocumentation(c *fiber.Ctx) error {
	ctx := c.UserContext()
	req := new(admin.InsertDocumentationRequest)

	if err := c.BodyParser(req); err != nil {
		return errors.InvalidRequest(fmt.Errorf("failed to parse request"))
	}

	if errs := d.Validator.Struct(req); errs != nil {
		return errors.InvalidRequest(common.FormatValidateError(errs))
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

// UpdateDocumentation Update a documentation.
//
//	@description	Update a documentation.
//	@id				admin-update-documentation
//	@summary		update documentation
//	@tags			Admin API
//	@accept			json
//	@produce		json
//	@param			admin.UpdateDocumentationRequest	body	admin.UpdateDocumentationRequest	true	"Update documentation request"
//	@security		Bearer
//	@success		200						{object}	vo.Response{data=nil}	"Success"
//	@failure		400						{object}	vo.Response{data=nil}	"Invalid request"
//	@failure		401						{object}	vo.Response{data=nil}	"Unauthorized"
//	@failure		403						{object}	vo.Response{data=nil}	"Forbidden"
//	@failure		404						{object}	vo.Response{data=nil}	"Not found"
//	@failure		500						{object}	vo.Response{data=nil}	"Internal server error"
//	@router			/admin/documentation	[put]
func (d DocumentationApi) UpdateDocumentation(c *fiber.Ctx) error {
	ctx := c.UserContext()
	req := new(admin.UpdateDocumentationRequest)

	if err := c.BodyParser(req); err != nil {
		return errors.InvalidRequest(fmt.Errorf("failed to parse request"))
	}

	if errs := d.Validator.Struct(req); errs != nil {
		return errors.InvalidRequest(common.FormatValidateError(errs))
	}
	documentationID, err := primitive.ObjectIDFromHex(*req.DocumentationID)
	if err != nil {
		return errors.InvalidRequest(fmt.Errorf("invalid documentation id"))
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

// DeleteDocumentation Delete a documentation.
//
//	@description	Delete a documentation.
//	@id				admin-delete-documentation
//	@summary		delete documentation
//	@tags			Admin API
//	@accept			json
//	@produce		json
//	@param			admin.DeleteDocumentationRequest	query	admin.DeleteDocumentationRequest	true	"Delete documentation request"
//	@security		Bearer
//	@success		200						{object}	vo.Response{data=nil}	"Success"
//	@failure		400						{object}	vo.Response{data=nil}	"Invalid request"
//	@failure		401						{object}	vo.Response{data=nil}	"Unauthorized"
//	@failure		403						{object}	vo.Response{data=nil}	"Forbidden"
//	@failure		404						{object}	vo.Response{data=nil}	"Not found"
//	@failure		500						{object}	vo.Response{data=nil}	"Internal server error"
//	@router			/admin/documentation	[delete]
func (d DocumentationApi) DeleteDocumentation(c *fiber.Ctx) error {
	ctx := c.UserContext()
	req := new(admin.DeleteDocumentationRequest)

	if err := c.QueryParser(req); err != nil {
		return errors.InvalidRequest(fmt.Errorf("failed to parse request"))
	}

	if errs := d.Validator.Struct(req); errs != nil {
		return errors.InvalidRequest(common.FormatValidateError(errs))
	}
	documentationID, err := primitive.ObjectIDFromHex(*req.DocumentationID)
	if err != nil {
		return errors.InvalidRequest(fmt.Errorf("invalid documentation id"))
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
