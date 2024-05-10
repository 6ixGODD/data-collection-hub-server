package mods

import (
	"context"

	dao "data-collection-hub-server/internal/pkg/dal/mods"
	"data-collection-hub-server/internal/pkg/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DocumentationService interface {
	InsertDocumentation(ctx context.Context, title, content *string) error
	UpdateDocumentation(ctx context.Context, documentationID *primitive.ObjectID, title, content *string) error
	DeleteDocumentation(ctx context.Context, documentationID *primitive.ObjectID) error
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

func (d DocumentationServiceImpl) InsertDocumentation(ctx context.Context, title, content *string) error {
	// TODO implement me
	panic("implement me")
}

func (d DocumentationServiceImpl) UpdateDocumentation(
	ctx context.Context, documentationID *primitive.ObjectID, title, content *string,
) error {
	// TODO implement me
	panic("implement me")
}

func (d DocumentationServiceImpl) DeleteDocumentation(ctx context.Context, documentationID *primitive.ObjectID) error {
	// TODO implement me
	panic("implement me")
}
