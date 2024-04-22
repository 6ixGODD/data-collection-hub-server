package modules

import (
	"context"

	"data-collection-hub-server/models"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
)

// DocumentationDao defines the crud methods that the infrastructure layer should implement
type DocumentationDao interface {
	GetDocumentationById(documentationId string, ctx context.Context) (*models.DocumentationModel, error)
	GetDocumentationList(offset, limit int64, ctx context.Context) ([]models.DocumentationModel, error)
	GetDocumentationListByCreatedTime(startTime, endTime string, offset, limit int64, ctx context.Context) ([]models.DocumentationModel, error)
	GetDocumentationListByUpdatedTime(startTime, endTime string, offset, limit int64, ctx context.Context) ([]models.DocumentationModel, error)
	InsertDocumentation(documentation *models.DocumentationModel, ctx context.Context) error
	UpdateDocumentation(documentation *models.DocumentationModel, ctx context.Context) error
	DeleteDocumentation(documentation *models.DocumentationModel, ctx context.Context) error
	DeleteDocumentationListByCreatedTime(startTime, endTime string, ctx context.Context) error
	DeleteDocumentationListByUpdatedTime(startTime, endTime string, ctx context.Context) error
}

// DocumentationDaoImpl implements the DocumentationDao interface and contains a qmgo.Collection instance
type DocumentationDaoImpl struct {
	documentationClient *qmgo.Collection
}

// NewDocumentationDao creates a new instance of DocumentationDaoImpl with the qmgo.Collection instance
func NewDocumentationDao(mongoDatabase *qmgo.Database) DocumentationDao {
	var _ DocumentationDao = new(DocumentationDaoImpl) // Ensure that the interface is implemented
	return &DocumentationDaoImpl{documentationClient: mongoDatabase.Collection("documentation")}
}

func (d *DocumentationDaoImpl) GetDocumentationById(documentationId string, ctx context.Context) (*models.DocumentationModel, error) {
	var documentation models.DocumentationModel
	err := d.documentationClient.Find(ctx, bson.M{"_id": documentationId}).One(&documentation)
	if err != nil {
		return nil, err
	} else {
		return &documentation, nil
	}
}

func (d *DocumentationDaoImpl) GetDocumentationList(offset, limit int64, ctx context.Context) ([]models.DocumentationModel, error) {
	var documentationList []models.DocumentationModel
	err := d.documentationClient.Find(ctx, bson.M{}).Skip(offset).Limit(limit).All(&documentationList)
	if err != nil {
		return nil, err
	} else {
		return documentationList, nil
	}
}

func (d *DocumentationDaoImpl) GetDocumentationListByCreatedTime(startTime, endTime string, offset, limit int64, ctx context.Context) ([]models.DocumentationModel, error) {
	var documentationList []models.DocumentationModel
	err := d.documentationClient.Find(ctx, bson.M{"created_at": bson.M{"$gte": startTime, "$lte": endTime}}).Skip(offset).Limit(limit).All(&documentationList)
	if err != nil {
		return nil, err
	} else {
		return documentationList, nil
	}
}

func (d *DocumentationDaoImpl) GetDocumentationListByUpdatedTime(startTime, endTime string, offset, limit int64, ctx context.Context) ([]models.DocumentationModel, error) {
	var documentationList []models.DocumentationModel
	err := d.documentationClient.Find(ctx, bson.M{"updated_at": bson.M{"$gte": startTime, "$lte": endTime}}).Skip(offset).Limit(limit).All(&documentationList)
	if err != nil {
		return nil, err
	} else {
		return documentationList, nil
	}
}

func (d *DocumentationDaoImpl) InsertDocumentation(documentation *models.DocumentationModel, ctx context.Context) error {
	_, err := d.documentationClient.InsertOne(ctx, documentation)
	return err
}

func (d *DocumentationDaoImpl) UpdateDocumentation(documentation *models.DocumentationModel, ctx context.Context) error {
	err := d.documentationClient.UpdateOne(ctx, bson.M{"_id": documentation.DocumentID}, bson.M{"$set": documentation})
	return err
}

func (d *DocumentationDaoImpl) DeleteDocumentation(documentation *models.DocumentationModel, ctx context.Context) error {
	err := d.documentationClient.RemoveId(ctx, documentation.DocumentID)
	return err
}

func (d *DocumentationDaoImpl) DeleteDocumentationListByCreatedTime(startTime, endTime string, ctx context.Context) error {
	_, err := d.documentationClient.RemoveAll(ctx, bson.M{"created_at": bson.M{"$gte": startTime, "$lte": endTime}})
	return err
}

func (d *DocumentationDaoImpl) DeleteDocumentationListByUpdatedTime(startTime, endTime string, ctx context.Context) error {
	_, err := d.documentationClient.RemoveAll(ctx, bson.M{"updated_at": bson.M{"$gte": startTime, "$lte": endTime}})
	return err
}
