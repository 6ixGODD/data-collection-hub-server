package modules

import (
	"context"

	"data-collection-hub-server/internal/pkg/dal"
	"data-collection-hub-server/internal/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var documentationCollectionName = "documentation"

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
type DocumentationDaoImpl struct{ *dal.Dao }

// NewDocumentationDao creates a new instance of DocumentationDaoImpl with the qmgo.Collection instance
func NewDocumentationDao(dao *dal.Dao) DocumentationDao {
	var _ DocumentationDao = new(DocumentationDaoImpl) // Ensure that the interface is implemented
	return &DocumentationDaoImpl{dao}
}

func (d *DocumentationDaoImpl) GetDocumentationById(documentationId string, ctx context.Context) (*models.DocumentationModel, error) {
	var documentation models.DocumentationModel
	collection := d.Dao.MongoClient.MongoDatabase.Collection(documentationCollectionName)
	err := collection.Find(ctx, bson.M{"_id": documentationId}).One(&documentation)
	if err != nil {
		d.Dao.Logger.Logger.Error(
			"DocumentationDaoImpl.GetDocumentationById",
			zap.Field{Key: "documentationId", Type: zapcore.StringType, String: documentationId},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		d.Dao.Logger.Logger.Info(
			"DocumentationDaoImpl.GetDocumentationById",
			zap.Field{Key: "documentationId", Type: zapcore.StringType, String: documentationId},
		)
		return &documentation, nil
	}
}

func (d *DocumentationDaoImpl) GetDocumentationList(offset, limit int64, ctx context.Context) ([]models.DocumentationModel, error) {
	var documentationList []models.DocumentationModel
	collection := d.Dao.MongoClient.MongoDatabase.Collection(documentationCollectionName)
	err := collection.Find(ctx, bson.M{}).Skip(offset).Limit(limit).All(&documentationList)
	if err != nil {
		d.Dao.Logger.Logger.Error(
			"DocumentationDaoImpl.GetDocumentationList",
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		d.Dao.Logger.Logger.Info(
			"DocumentationDaoImpl.GetDocumentationList",
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return documentationList, nil
	}
}

func (d *DocumentationDaoImpl) GetDocumentationListByCreatedTime(startTime, endTime string, offset, limit int64, ctx context.Context) ([]models.DocumentationModel, error) {
	var documentationList []models.DocumentationModel
	collection := d.Dao.MongoClient.MongoDatabase.Collection(documentationCollectionName)
	err := collection.Find(ctx, bson.M{"created_at": bson.M{"$gte": startTime, "$lte": endTime}}).Skip(offset).Limit(limit).All(&documentationList)
	if err != nil {
		d.Dao.Logger.Logger.Error(
			"DocumentationDaoImpl.GetDocumentationListByCreatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		d.Dao.Logger.Logger.Info(
			"DocumentationDaoImpl.GetDocumentationListByCreatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return documentationList, nil
	}
}

func (d *DocumentationDaoImpl) GetDocumentationListByUpdatedTime(startTime, endTime string, offset, limit int64, ctx context.Context) ([]models.DocumentationModel, error) {
	var documentationList []models.DocumentationModel
	collection := d.Dao.MongoClient.MongoDatabase.Collection(documentationCollectionName)
	err := collection.Find(ctx, bson.M{"updated_at": bson.M{"$gte": startTime, "$lte": endTime}}).Skip(offset).Limit(limit).All(&documentationList)
	if err != nil {
		d.Dao.Logger.Logger.Error(
			"DocumentationDaoImpl.GetDocumentationListByUpdatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
		return nil, err
	} else {
		d.Dao.Logger.Logger.Info(
			"DocumentationDaoImpl.GetDocumentationListByUpdatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime},
			zap.Field{Key: "offset", Type: zapcore.Int64Type, Integer: offset},
			zap.Field{Key: "limit", Type: zapcore.Int64Type, Integer: limit},
		)
		return documentationList, nil
	}
}

func (d *DocumentationDaoImpl) InsertDocumentation(documentation *models.DocumentationModel, ctx context.Context) error {
	collection := d.Dao.MongoClient.MongoDatabase.Collection(documentationCollectionName)
	_, err := collection.InsertOne(ctx, documentation)
	if err != nil {
		d.Dao.Logger.Logger.Error(
			"DocumentationDaoImpl.InsertDocumentation",
			zap.Field{Key: "documentation", Type: zapcore.ObjectMarshalerType, Interface: documentation},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
	} else {
		d.Dao.Logger.Logger.Info(
			"DocumentationDaoImpl.InsertDocumentation",
			zap.Field{Key: "documentation", Type: zapcore.ObjectMarshalerType, Interface: documentation},
		)
	}
	return err
}

func (d *DocumentationDaoImpl) UpdateDocumentation(documentation *models.DocumentationModel, ctx context.Context) error {
	collection := d.Dao.MongoClient.MongoDatabase.Collection(documentationCollectionName)
	err := collection.UpdateOne(ctx, bson.M{"_id": documentation.DocumentID}, bson.M{"$set": documentation})
	if err != nil {
		d.Dao.Logger.Logger.Error(
			"DocumentationDaoImpl.UpdateDocumentation",
			zap.Field{Key: "documentation", Type: zapcore.ObjectMarshalerType, Interface: documentation},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
	} else {
		d.Dao.Logger.Logger.Info(
			"DocumentationDaoImpl.UpdateDocumentation",
			zap.Field{Key: "documentation", Type: zapcore.ObjectMarshalerType, Interface: documentation},
		)
	}
	return err
}

func (d *DocumentationDaoImpl) DeleteDocumentation(documentation *models.DocumentationModel, ctx context.Context) error {
	collection := d.Dao.MongoClient.MongoDatabase.Collection(documentationCollectionName)
	err := collection.RemoveId(ctx, documentation.DocumentID)
	if err != nil {
		d.Dao.Logger.Logger.Error(
			"DocumentationDaoImpl.DeleteDocumentation",
			zap.Field{Key: "documentation", Type: zapcore.ObjectMarshalerType, Interface: documentation},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
	} else {
		d.Dao.Logger.Logger.Info(
			"DocumentationDaoImpl.DeleteDocumentation",
			zap.Field{Key: "documentation", Type: zapcore.ObjectMarshalerType, Interface: documentation},
		)
	}
	return err
}

func (d *DocumentationDaoImpl) DeleteDocumentationListByCreatedTime(startTime, endTime string, ctx context.Context) error {
	collection := d.Dao.MongoClient.MongoDatabase.Collection(documentationCollectionName)
	result, err := collection.RemoveAll(ctx, bson.M{"created_at": bson.M{"$gte": startTime, "$lte": endTime}})
	if err != nil {
		d.Dao.Logger.Logger.Error(
			"DocumentationDaoImpl.DeleteDocumentationListByCreatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
	} else {
		d.Dao.Logger.Logger.Info(
			"DocumentationDaoImpl.DeleteDocumentationListByCreatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime},
			zap.Field{Key: "result", Type: zapcore.ObjectMarshalerType, Interface: result},
		)
	}
	return err
}

func (d *DocumentationDaoImpl) DeleteDocumentationListByUpdatedTime(startTime, endTime string, ctx context.Context) error {
	collection := d.Dao.MongoClient.MongoDatabase.Collection(documentationCollectionName)
	result, err := collection.RemoveAll(ctx, bson.M{"updated_at": bson.M{"$gte": startTime, "$lte": endTime}})
	if err != nil {
		d.Dao.Logger.Logger.Error(
			"DocumentationDaoImpl.DeleteDocumentationListByUpdatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime},
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
	} else {
		d.Dao.Logger.Logger.Info(
			"DocumentationDaoImpl.DeleteDocumentationListByUpdatedTime",
			zap.Field{Key: "startTime", Type: zapcore.StringType, String: startTime},
			zap.Field{Key: "endTime", Type: zapcore.StringType, String: endTime},
			zap.Field{Key: "result", Type: zapcore.ObjectMarshalerType, Interface: result},
		)
	}
	return err
}
