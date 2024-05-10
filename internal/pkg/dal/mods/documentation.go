package mods

import (
	"context"
	"time"

	"data-collection-hub-server/internal/pkg/dal"
	"data-collection-hub-server/internal/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const documentationCollectionName = "documentation"

// DocumentationDao defines the crud methods that the infrastructure layer should implement
type DocumentationDao interface {
	GetDocumentationById(documentationId primitive.ObjectID, ctx context.Context) (*models.DocumentationModel, error)
	GetDocumentationList(offset, limit int64, ctx context.Context) ([]models.DocumentationModel, error)
	GetDocumentationListByCreatedTime(
		startTime, endTime time.Time, offset, limit int64, ctx context.Context,
	) ([]models.DocumentationModel, error)
	GetDocumentationListByUpdatedTime(
		startTime, endTime time.Time, offset, limit int64, ctx context.Context,
	) ([]models.DocumentationModel, error)
	InsertDocumentation(title, content string, ctx context.Context) (primitive.ObjectID, error)
	UpdateDocumentation(documentationId primitive.ObjectID, title, content string, ctx context.Context) error
	DeleteDocumentation(documentationId primitive.ObjectID, ctx context.Context) error
	DeleteDocumentationListByCreatedTime(startTime, endTime time.Time, ctx context.Context) error
	DeleteDocumentationListByUpdatedTime(startTime, endTime time.Time, ctx context.Context) error
}

// DocumentationDaoImpl implements the DocumentationDao interface and contains a qmgo.Collection instance
type DocumentationDaoImpl struct{ *dal.Dao }

// NewDocumentationDao creates a new instance of DocumentationDaoImpl with the qmgo.Collection instance
func NewDocumentationDao(dao *dal.Dao) DocumentationDao {
	var _ DocumentationDao = (*DocumentationDaoImpl)(nil) // Ensure that the interface is implemented
	return &DocumentationDaoImpl{dao}
}

func (d *DocumentationDaoImpl) GetDocumentationById(
	documentationId primitive.ObjectID, ctx context.Context,
) (*models.DocumentationModel, error) {
	var documentation models.DocumentationModel
	collection := d.Dao.Mongo.MongoDatabase.Collection(documentationCollectionName)
	err := collection.Find(ctx, bson.M{"_id": documentationId}).One(&documentation)
	if err != nil {
		d.Dao.Zap.Logger.Error(
			"DocumentationDaoImpl.GetDocumentationById",
			zap.Field{Key: "documentationId", Type: zapcore.StringType, String: documentationId.Hex()},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		d.Dao.Zap.Logger.Info(
			"DocumentationDaoImpl.GetDocumentationById",
			zap.Field{Key: "documentationId", Type: zapcore.StringType, String: documentationId.Hex()},
		)
		return &documentation, nil
	}
}

func (d *DocumentationDaoImpl) GetDocumentationList(
	offset, limit int64, ctx context.Context,
) ([]models.DocumentationModel, error) {
	var documentationList []models.DocumentationModel
	collection := d.Dao.Mongo.MongoDatabase.Collection(documentationCollectionName)
	err := collection.Find(ctx, bson.M{}).Skip(offset).Limit(limit).All(&documentationList)
	if err != nil {
		d.Dao.Zap.Logger.Error(
			"DocumentationDaoImpl.GetDocumentationList",
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		d.Dao.Zap.Logger.Info(
			"DocumentationDaoImpl.GetDocumentationList",
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return documentationList, nil
	}
}

func (d *DocumentationDaoImpl) GetDocumentationListByCreatedTime(
	startTime, endTime time.Time, offset, limit int64, ctx context.Context,
) ([]models.DocumentationModel, error) {
	var documentationList []models.DocumentationModel
	collection := d.Dao.Mongo.MongoDatabase.Collection(documentationCollectionName)
	err := collection.Find(
		ctx, bson.M{"created_at": bson.M{"$gte": startTime, "$lte": endTime}},
	).Skip(offset).Limit(limit).All(&documentationList)
	if err != nil {
		d.Dao.Zap.Logger.Error(
			"DocumentationDaoImpl.GetDocumentationListByCreatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime.Format(time.RFC3339)},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime.Format(time.RFC3339)},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		d.Dao.Zap.Logger.Info(
			"DocumentationDaoImpl.GetDocumentationListByCreatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime.Format(time.RFC3339)},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime.Format(time.RFC3339)},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return documentationList, nil
	}
}

func (d *DocumentationDaoImpl) GetDocumentationListByUpdatedTime(
	startTime, endTime time.Time, offset, limit int64, ctx context.Context,
) ([]models.DocumentationModel, error) {
	var documentationList []models.DocumentationModel
	collection := d.Dao.Mongo.MongoDatabase.Collection(documentationCollectionName)
	err := collection.Find(
		ctx, bson.M{"updated_at": bson.M{"$gte": startTime, "$lte": endTime}},
	).Skip(offset).Limit(limit).All(&documentationList)
	if err != nil {
		d.Dao.Zap.Logger.Error(
			"DocumentationDaoImpl.GetDocumentationListByUpdatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime.Format(time.RFC3339)},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime.Format(time.RFC3339)},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		d.Dao.Zap.Logger.Info(
			"DocumentationDaoImpl.GetDocumentationListByUpdatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime.Format(time.RFC3339)},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime.Format(time.RFC3339)},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return documentationList, nil
	}
}

func (d *DocumentationDaoImpl) InsertDocumentation(
	title, content string, ctx context.Context,
) (primitive.ObjectID, error) {
	documentation := models.DocumentationModel{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	collection := d.Dao.Mongo.MongoDatabase.Collection(documentationCollectionName)
	result, err := collection.InsertOne(ctx, documentation)
	if err != nil {
		d.Dao.Zap.Logger.Error(
			"DocumentationDaoImpl.InsertDocumentation",
			zap.Field{Key: "documentation", Type: zapcore.ObjectMarshalerType, Interface: documentation},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
	} else {
		d.Dao.Zap.Logger.Info(
			"DocumentationDaoImpl.InsertDocumentation",
			zap.Field{
				Key: "documentId", Type: zapcore.StringType, String: result.InsertedID.(primitive.ObjectID).Hex(),
			},
			zap.Field{Key: "documentation", Type: zapcore.ObjectMarshalerType, Interface: documentation},
		)
	}
	return result.InsertedID.(primitive.ObjectID), err
}

func (d *DocumentationDaoImpl) UpdateDocumentation(
	documentationId primitive.ObjectID, title, content string, ctx context.Context,
) error {
	collection := d.Dao.Mongo.MongoDatabase.Collection(documentationCollectionName)
	err := collection.UpdateOne(
		ctx, bson.M{"_id": documentationId},
		bson.M{"$set": bson.M{"title": title, "content": content, "updated_at": time.Now()}},
	)
	if err != nil {
		d.Dao.Zap.Logger.Error(
			"DocumentationDaoImpl.UpdateDocumentation",
			zap.Field{Key: "documentationId", Type: zapcore.StringType, String: documentationId.Hex()},
			zap.Field{Key: "title", Type: zapcore.StringType, String: title},
			zap.Field{Key: "content", Type: zapcore.StringType, String: content},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
	} else {
		d.Dao.Zap.Logger.Info(
			"DocumentationDaoImpl.UpdateDocumentation",
			zap.Field{Key: "documentationId", Type: zapcore.StringType, String: documentationId.Hex()},
			zap.Field{Key: "title", Type: zapcore.StringType, String: title},
			zap.Field{Key: "content", Type: zapcore.StringType, String: content},
		)
	}
	return err
}

func (d *DocumentationDaoImpl) DeleteDocumentation(documentationId primitive.ObjectID, ctx context.Context) error {
	collection := d.Dao.Mongo.MongoDatabase.Collection(documentationCollectionName)
	err := collection.RemoveId(ctx, documentationId)
	if err != nil {
		d.Dao.Zap.Logger.Error(
			"DocumentationDaoImpl.DeleteDocumentation",
			zap.Field{Key: "documentationId", Type: zapcore.StringType, String: documentationId.Hex()},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
	} else {
		d.Dao.Zap.Logger.Info(
			"DocumentationDaoImpl.DeleteDocumentation",
			zap.Field{Key: "documentationId", Type: zapcore.StringType, String: documentationId.Hex()},
		)
	}
	return err
}

func (d *DocumentationDaoImpl) DeleteDocumentationListByCreatedTime(
	startTime, endTime time.Time, ctx context.Context,
) error {
	collection := d.Dao.Mongo.MongoDatabase.Collection(documentationCollectionName)
	result, err := collection.RemoveAll(ctx, bson.M{"created_at": bson.M{"$gte": startTime, "$lte": endTime}})
	if err != nil {
		d.Dao.Zap.Logger.Error(
			"DocumentationDaoImpl.DeleteDocumentationListByCreatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime.Format(time.RFC3339)},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime.Format(time.RFC3339)},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
	} else {
		d.Dao.Zap.Logger.Info(
			"DocumentationDaoImpl.DeleteDocumentationListByCreatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime.Format(time.RFC3339)},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime.Format(time.RFC3339)},
			zap.Field{Key: "result", Type: zapcore.ObjectMarshalerType, Interface: result},
		)
	}
	return err
}

func (d *DocumentationDaoImpl) DeleteDocumentationListByUpdatedTime(
	startTime, endTime time.Time, ctx context.Context,
) error {
	collection := d.Dao.Mongo.MongoDatabase.Collection(documentationCollectionName)
	result, err := collection.RemoveAll(ctx, bson.M{"updated_at": bson.M{"$gte": startTime, "$lte": endTime}})
	if err != nil {
		d.Dao.Zap.Logger.Error(
			"DocumentationDaoImpl.DeleteDocumentationListByUpdatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime.Format(time.RFC3339)},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime.Format(time.RFC3339)},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
	} else {
		d.Dao.Zap.Logger.Info(
			"DocumentationDaoImpl.DeleteDocumentationListByUpdatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime.Format(time.RFC3339)},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime.Format(time.RFC3339)},
			zap.Field{Key: "result", Type: zapcore.ObjectMarshalerType, Interface: result},
		)
	}
	return err
}
