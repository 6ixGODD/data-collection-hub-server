package mods

import (
	"time"

	"data-collection-hub-server/internal/pkg/schema"
	"data-collection-hub-server/internal/pkg/schema/common"
	commonservice "data-collection-hub-server/internal/pkg/service/common/mods"
	"data-collection-hub-server/pkg/errors"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DocumentationApi struct {
	commonservice.DocumentationService
}

func NewDocumentationApi(documentationService commonservice.DocumentationService) DocumentationApi {
	return DocumentationApi{documentationService}
}

func (d DocumentationApi) GetDocumentation(c *fiber.Ctx) error {
	req := new(common.GetDocumentationRequest)

	if err := c.QueryParser(req); err != nil {
		return errors.InvalidRequest(err)
	}

	documentationID, err := primitive.ObjectIDFromHex(*req.DocumentationID)
	if err != nil {
		return errors.InvalidRequest(err)
	}
	resp, err := d.DocumentationService.GetDocumentation(c.Context(), &documentationID)
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

func (d DocumentationApi) GetDocumentationList(c *fiber.Ctx) error {
	req := new(common.GetDocumentationListRequest)

	if err := c.QueryParser(req); err != nil {
		return errors.InvalidRequest(err)
	}

	var (
		updateBefore, updateAfter *time.Time
		err                       error
	)

	if req.UpdateBefore != nil {
		*updateBefore, err = time.Parse(time.RFC3339, *req.UpdateBefore)
		if err != nil {
			return errors.InvalidRequest(err)
		}
	}
	if req.UpdateAfter != nil {
		*updateAfter, err = time.Parse(time.RFC3339, *req.UpdateAfter)
		if err != nil {
			return errors.InvalidRequest(err)
		}
	}

	resp, err := d.DocumentationService.GetDocumentationList(c.Context(), req.Page, updateBefore, updateAfter)
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
