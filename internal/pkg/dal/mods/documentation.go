package mods

import (
	"context"
	"time"

	"data-collection-hub-server/internal/pkg/dal"
	"data-collection-hub-server/internal/pkg/models"
	"github.com/goccy/go-json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

const documentationCollectionName = "documentation"

// DocumentationDao defines the crud methods that the infrastructure layer should implement
type DocumentationDao interface {
	GetDocumentationById(ctx context.Context, documentationId primitive.ObjectID) (*models.DocumentationModel, error)
	GetDocumentationList(
		ctx context.Context,
		offset, limit int64, desc bool, createStartTime, createEndTime, updateStartTime, updateEndTime *time.Time,
	) ([]models.DocumentationModel, *int64, error)
	InsertDocumentation(ctx context.Context, title, content string) (primitive.ObjectID, error)
	UpdateDocumentation(ctx context.Context, documentationId primitive.ObjectID, title, content *string) error
	DeleteDocumentation(ctx context.Context, documentationId primitive.ObjectID) error
	DeleteDocumentationList(
		ctx context.Context, createStartTime, createEndTime, updateStartTime, updateEndTime *time.Time,
		title, content *string,
	) (*int64, error)
}

// DocumentationDaoImpl implements the DocumentationDao interface and contains a qmgo.Collection instance
type DocumentationDaoImpl struct{ *dal.Dao }

// NewDocumentationDao creates a new instance of DocumentationDaoImpl with the qmgo.Collection instance
func NewDocumentationDao(dao *dal.Dao) DocumentationDao {
	var _ DocumentationDao = (*DocumentationDaoImpl)(nil) // Ensure that the interface is implemented
	return &DocumentationDaoImpl{dao}
}

func (d *DocumentationDaoImpl) GetDocumentationById(
	ctx context.Context, documentationId primitive.ObjectID,
) (*models.DocumentationModel, error) {
	var documentation models.DocumentationModel
	collection := d.Dao.Mongo.MongoDatabase.Collection(documentationCollectionName)
	err := collection.Find(ctx, bson.M{"_id": documentationId}).One(&documentation)
	if err != nil {
		d.Dao.Zap.Logger.Error(
			"DocumentationDaoImpl.GetDocumentationById",
			zap.String("documentationId", documentationId.Hex()),
			zap.Error(err),
		)
		return nil, err
	} else {
		d.Dao.Zap.Logger.Info(
			"DocumentationDaoImpl.GetDocumentationById",
			zap.String("documentationId", documentationId.Hex()),
		)
		return &documentation, nil
	}
}

func (d *DocumentationDaoImpl) GetDocumentationList(
	ctx context.Context,
	offset, limit int64, desc bool, createStartTime, createEndTime, updateStartTime, updateEndTime *time.Time,
) ([]models.DocumentationModel, *int64, error) {
	var documentationList []models.DocumentationModel
	var err error
	collection := d.Dao.Mongo.MongoDatabase.Collection(documentationCollectionName)
	doc := bson.M{}
	if createStartTime != nil && createEndTime != nil {
		doc["created_at"] = bson.M{"$gte": createStartTime, "$lte": createEndTime}
	}
	if updateStartTime != nil && updateEndTime != nil {
		doc["updated_at"] = bson.M{"$gte": updateStartTime, "$lte": updateEndTime}
	}
	docJSON, _ := json.Marshal(doc)
	if desc {
		err = collection.Find(ctx, doc).Sort("-created_at").Skip(offset).Limit(limit).All(&documentationList)
	} else {
		err = collection.Find(ctx, doc).Skip(offset).Limit(limit).All(&documentationList)
	}
	if err != nil {
		d.Dao.Zap.Logger.Error(
			"DocumentationDaoImpl.GetDocumentationList",
			zap.Int64("offset", offset), zap.Int64("limit", limit),
			zap.Bool("desc", desc), zap.ByteString(documentationCollectionName, docJSON), zap.Error(err),
		)
		return nil, nil, err
	}
	count, err := collection.Find(ctx, doc).Count()
	if err != nil {
		d.Dao.Zap.Logger.Error(
			"DocumentationDaoImpl.GetDocumentationList",
			zap.Int64("offset", offset), zap.Int64("limit", limit),
			zap.ByteString(documentationCollectionName, docJSON), zap.Error(err),
		)
		return nil, nil, err
	}

	d.Dao.Zap.Logger.Info(
		"DocumentationDaoImpl.GetDocumentationList",
		zap.Int64("offset", offset), zap.Int64("limit", limit),
		zap.ByteString(documentationCollectionName, docJSON), zap.Int64("count", count),
	)
	return documentationList, &count, nil
}

func (d *DocumentationDaoImpl) InsertDocumentation(
	ctx context.Context, title, content string,
) (primitive.ObjectID, error) {
	collection := d.Dao.Mongo.MongoDatabase.Collection(documentationCollectionName)
	doc := bson.M{
		"title": title, "content": content, "created_at": time.Now(), "updated_at": time.Now(),
	}
	docJSON, _ := json.Marshal(doc)
	result, err := collection.InsertOne(ctx, doc)
	if err != nil {
		d.Dao.Zap.Logger.Error(
			"DocumentationDaoImpl.InsertDocumentation",
			zap.ByteString(documentationCollectionName, docJSON),
			zap.Error(err),
		)
	} else {
		d.Dao.Zap.Logger.Info(
			"DocumentationDaoImpl.InsertDocumentation",
			zap.ByteString(documentationCollectionName, docJSON),
			zap.String("documentation_id", result.InsertedID.(primitive.ObjectID).Hex()),
		)
	}
	return result.InsertedID.(primitive.ObjectID), err
}

func (d *DocumentationDaoImpl) UpdateDocumentation(
	ctx context.Context, documentationId primitive.ObjectID, title, content *string,
) error {
	collection := d.Dao.Mongo.MongoDatabase.Collection(documentationCollectionName)
	doc := bson.M{"updated_at": time.Now()}
	if title != nil {
		doc["title"] = *title
	}
	if content != nil {
		doc["content"] = *content
	}
	docJSON, _ := json.Marshal(doc)
	err := collection.UpdateId(ctx, documentationId, bson.M{"$set": doc})
	if err != nil {
		d.Dao.Zap.Logger.Error(
			"DocumentationDaoImpl.UpdateDocumentation",
			zap.String("documentationId", documentationId.Hex()),
			zap.ByteString(documentationCollectionName, docJSON),
			zap.Error(err),
		)
	} else {
		d.Dao.Zap.Logger.Info(
			"DocumentationDaoImpl.UpdateDocumentation",
			zap.String("documentationId", documentationId.Hex()),
			zap.ByteString(documentationCollectionName, docJSON),
		)
	}
	return err
}

func (d *DocumentationDaoImpl) DeleteDocumentation(
	ctx context.Context, documentationId primitive.ObjectID,
) error {
	collection := d.Dao.Mongo.MongoDatabase.Collection(documentationCollectionName)
	err := collection.RemoveId(ctx, documentationId)
	if err != nil {
		d.Dao.Zap.Logger.Error(
			"DocumentationDaoImpl.DeleteDocumentation",
			zap.String("documentationId", documentationId.Hex()),
			zap.Error(err),
		)
	} else {
		d.Dao.Zap.Logger.Info(
			"DocumentationDaoImpl.DeleteDocumentation",
			zap.String("documentationId", documentationId.Hex()),
		)
	}
	return err
}

func (d *DocumentationDaoImpl) DeleteDocumentationList(
	ctx context.Context, createStartTime, createEndTime, updateStartTime, updateEndTime *time.Time,
	title, content *string,
) (*int64, error) {
	collection := d.Dao.Mongo.MongoDatabase.Collection(documentationCollectionName)
	doc := bson.M{}
	if title != nil {
		doc["title"] = *title
	}
	if content != nil {
		doc["content"] = *content
	}
	if createStartTime != nil && createEndTime != nil {
		doc["created_at"] = bson.M{"$gte": createStartTime, "$lte": createEndTime}
	}
	if updateStartTime != nil && updateEndTime != nil {
		doc["updated_at"] = bson.M{"$gte": updateStartTime, "$lte": updateEndTime}
	}
	docJSON, _ := json.Marshal(doc)
	result, err := collection.RemoveAll(ctx, doc)
	if err != nil {
		d.Dao.Zap.Logger.Error(
			"DocumentationDaoImpl.DeleteDocumentationList",
			zap.ByteString(documentationCollectionName, docJSON),
			zap.Error(err),
		)
		return nil, err
	}
	d.Dao.Zap.Logger.Info(
		"DocumentationDaoImpl.DeleteDocumentationList",
		zap.ByteString(documentationCollectionName, docJSON),
		zap.Int64("deleted_count", result.DeletedCount),
	)
	return &result.DeletedCount, nil
}
