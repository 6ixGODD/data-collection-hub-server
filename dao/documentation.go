package dao

import (
	"context"
	"data-collection-hub-server/models"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
)

type DocumentationDao interface {
	GetDocumentationById(documentationId string, mongoClient *qmgo.QmgoClient, ctx context.Context) (*models.DocumentationModel, error)
	GetDocumentationList(mongoClient *qmgo.QmgoClient, offset, limit int64, ctx context.Context) ([]models.DocumentationModel, error)
	GetDocumentationListByCreatedTime(mongoClient *qmgo.QmgoClient, startTime, endTime string, offset, limit int64, ctx context.Context) ([]models.DocumentationModel, error)
	GetDocumentationListByUpdatedTime(mongoClient *qmgo.QmgoClient, startTime, endTime string, offset, limit int64, ctx context.Context) ([]models.DocumentationModel, error)
	InsertDocumentation(documentation *models.DocumentationModel, mongoClient *qmgo.QmgoClient, ctx context.Context) error
	UpdateDocumentation(documentation *models.DocumentationModel, mongoClient *qmgo.QmgoClient, ctx context.Context) error
	DeleteDocumentation(documentation *models.DocumentationModel, mongoClient *qmgo.QmgoClient, ctx context.Context) error
	DeleteDocumentationListByCreatedTime(startTime, endTime string, mongoClient *qmgo.QmgoClient, ctx context.Context) error
	DeleteDocumentationListByUpdatedTime(startTime, endTime string, mongoClient *qmgo.QmgoClient, ctx context.Context) error
}

type DocumentationDaoImpl struct{}

func (documentationDao *DocumentationDaoImpl) GetDocumentationById(documentationId string, mongoClient *qmgo.QmgoClient, ctx context.Context) (*models.DocumentationModel, error) {
	var documentation models.DocumentationModel
	err := mongoClient.Find(ctx, bson.M{"_id": documentationId}).One(&documentation)
	if err != nil {
		return nil, err
	} else {
		return &documentation, nil
	}
}

func (documentationDao *DocumentationDaoImpl) GetDocumentationList(mongoClient *qmgo.QmgoClient, offset, limit int64, ctx context.Context) ([]models.DocumentationModel, error) {
	var documentationList []models.DocumentationModel
	err := mongoClient.Find(ctx, bson.M{}).Skip(offset).Limit(limit).All(&documentationList)
	if err != nil {
		return nil, err
	} else {
		return documentationList, nil
	}
}

func (documentationDao *DocumentationDaoImpl) GetDocumentationListByCreatedTime(mongoClient *qmgo.QmgoClient, startTime, endTime string, offset, limit int64, ctx context.Context) ([]models.DocumentationModel, error) {
	var documentationList []models.DocumentationModel
	err := mongoClient.Find(ctx, bson.M{"created_at": bson.M{"$gte": startTime, "$lte": endTime}}).Skip(offset).Limit(limit).All(&documentationList)
	if err != nil {
		return nil, err
	} else {
		return documentationList, nil
	}
}

func (documentationDao *DocumentationDaoImpl) GetDocumentationListByUpdatedTime(mongoClient *qmgo.QmgoClient, startTime, endTime string, offset, limit int64, ctx context.Context) ([]models.DocumentationModel, error) {
	var documentationList []models.DocumentationModel
	err := mongoClient.Find(ctx, bson.M{"updated_at": bson.M{"$gte": startTime, "$lte": endTime}}).Skip(offset).Limit(limit).All(&documentationList)
	if err != nil {
		return nil, err
	} else {
		return documentationList, nil
	}
}

func (documentationDao *DocumentationDaoImpl) InsertDocumentation(documentation *models.DocumentationModel, mongoClient *qmgo.QmgoClient, ctx context.Context) error {
	_, err := mongoClient.InsertOne(ctx, documentation)
	return err
}

func (documentationDao *DocumentationDaoImpl) UpdateDocumentation(documentation *models.DocumentationModel, mongoClient *qmgo.QmgoClient, ctx context.Context) error {
	return nil
}

func (documentationDao *DocumentationDaoImpl) DeleteDocumentation(documentation *models.DocumentationModel, mongoClient *qmgo.QmgoClient, ctx context.Context) error {
	// TODO implement me
	panic("implement me")
}

func (documentationDao *DocumentationDaoImpl) DeleteDocumentationListByCreatedTime(startTime, endTime string, mongoClient *qmgo.QmgoClient, ctx context.Context) error {
	return nil
}

func (documentationDao *DocumentationDaoImpl) DeleteDocumentationListByUpdatedTime(startTime, endTime string, mongoClient *qmgo.QmgoClient, ctx context.Context) error {
	return nil
}

func NewDocumentationDaoImpl() *DocumentationDaoImpl {
	var _ DocumentationDao = new(DocumentationDaoImpl) // Ensure that the interface is implemented
	return &DocumentationDaoImpl{}
}
