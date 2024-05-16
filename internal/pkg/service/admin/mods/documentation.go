package mods

import (
	"context"

	dao "data-collection-hub-server/internal/pkg/dao/mods"
	"data-collection-hub-server/internal/pkg/service"
	"data-collection-hub-server/pkg/errors"
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
	_, err := d.documentationDao.InsertDocumentation(ctx, *title, *content)
	if err != nil {
		return errors.MongoError(errors.WriteError(err))
	}
	return nil
}

func (d DocumentationServiceImpl) UpdateDocumentation(
	ctx context.Context, documentationID *primitive.ObjectID, title, content *string,
) error {
	err := d.documentationDao.UpdateDocumentation(ctx, *documentationID, title, content)
	if err != nil {
		return errors.MongoError(errors.WriteError(err))
	}
	return nil
}

func (d DocumentationServiceImpl) DeleteDocumentation(ctx context.Context, documentationID *primitive.ObjectID) error {
	err := d.documentationDao.DeleteDocumentation(ctx, *documentationID)
	if err != nil {
		return errors.MongoError(errors.WriteError(err))
	}
	return nil
}
