package mods

import (
	"context"
	"time"

	dao "data-collection-hub-server/internal/pkg/dal/mods"
	"data-collection-hub-server/internal/pkg/schema/common"
	"data-collection-hub-server/internal/pkg/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DocumentationService interface {
	GetDocumentation(ctx context.Context, documentationID *primitive.ObjectID) (*common.GetDocumentationResponse, error)
	GetDocumentationList(
		ctx context.Context, page *int, updateBefore, updateAfter *time.Time,
	) (*common.GetDocumentationListResponse, error)
}

type DocumentationServiceImpl struct {
	service          *service.Service
	documentationDao dao.DocumentationDao
}

func NewDocumentationService(s *service.Service, documentationDao dao.DocumentationDao) DocumentationService {
	return &DocumentationServiceImpl{
		service:          s,
		documentationDao: documentationDao,
	}
}

func (d DocumentationServiceImpl) GetDocumentation(
	ctx context.Context, documentationID *primitive.ObjectID,
) (*common.GetDocumentationResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (d DocumentationServiceImpl) GetDocumentationList(
	ctx context.Context, page *int, updateBefore, updateAfter *time.Time,
) (*common.GetDocumentationListResponse, error) {
	// TODO implement me
	panic("implement me")
}
