package dao

import (
	"context"

	"data-collection-hub-server/pkg/mongo"
	logging "data-collection-hub-server/pkg/zap"
	"go.uber.org/zap"
)

type Dao struct {
	Mongo  *mongo.Mongo
	Logger *zap.Logger
}

func New(ctx context.Context, mongo *mongo.Mongo, zap *logging.Zap) (*Dao, error) {
	c := zap.SetTagInContext(ctx, logging.MongoTag)
	logger, err := zap.GetLogger(c)
	if err != nil {
		return nil, err
	}
	return &Dao{
		Mongo:  mongo,
		Logger: logger,
	}, nil
}
