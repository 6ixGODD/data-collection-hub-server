package mods

import (
	"data-collection-hub-server/internal/pkg/schema"
	"data-collection-hub-server/internal/pkg/schema/admin"
	adminservice "data-collection-hub-server/internal/pkg/service/admin/mods"
	"data-collection-hub-server/pkg/errors"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DocumentationApi struct {
	adminservice.DocumentationService
}

func NewDocumentationApi(documentationService adminservice.DocumentationService) DocumentationApi {
	return DocumentationApi{documentationService}
}

func (d DocumentationApi) InsertDocumentation(c *fiber.Ctx) error {
	req := new(admin.InsertDocumentationRequest)

	if err := c.BodyParser(req); err != nil {
		return errors.InvalidRequest(err)
	}

	err := d.DocumentationService.InsertDocumentation(c.Context(), req.Title, req.Content)
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

func (d DocumentationApi) UpdateDocumentation(c *fiber.Ctx) error {
	req := new(admin.UpdateDocumentationRequest)

	if err := c.BodyParser(req); err != nil {
		return errors.InvalidRequest(err)
	}

	documentationID, err := primitive.ObjectIDFromHex(*req.DocumentationID)
	if err != nil {
		return errors.InvalidRequest(err)
	}
	err = d.DocumentationService.UpdateDocumentation(c.Context(), &documentationID, req.Title, req.Content)
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

func (d DocumentationApi) DeleteDocumentation(c *fiber.Ctx) error {
	req := new(admin.DeleteDocumentationRequest)

	if err := c.BodyParser(req); err != nil {
		return errors.InvalidRequest(err)
	}

	documentationID, err := primitive.ObjectIDFromHex(*req.DocumentationID)
	if err != nil {
		return errors.InvalidRequest(err)
	}
	err = d.DocumentationService.DeleteDocumentation(c.Context(), &documentationID)
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
