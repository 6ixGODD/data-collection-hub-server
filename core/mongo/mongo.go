package mongo

import (
	"context"

	"data-collection-hub-server/core/config/modules"
	"github.com/qiniu/qmgo"
)

var (
	mongoClient   *qmgo.Client
	mongoDatabase *qmgo.Database
	mongoConfig   *modules.MongoConfig
)

func InitMongo(ctx context.Context, config *modules.MongoConfig) error {
	var err error
	mongoConfig = config
	mongoClient, err = qmgo.NewClient(ctx, &qmgo.Config{
		Uri:              mongoConfig.Uri,
		ConnectTimeoutMS: &mongoConfig.ConnectTimeoutMS,
		SocketTimeoutMS:  &mongoConfig.SocketTimeoutMS,
		MaxPoolSize:      &mongoConfig.MaxPoolSize,
		MinPoolSize:      &mongoConfig.MinPoolSize,
	})
	if err != nil {
		return err
	} else {
		mongoDatabase = mongoClient.Database(config.Database)
		return nil
	}
}

func GetMongoDatabase() (d *qmgo.Database, e error) {
	if mongoDatabase == nil {
		if err := InitMongo(context.Background(), mongoConfig); err != nil {
			return nil, err
		}
	}
	return mongoDatabase, nil
}

func CloseMongo(ctx context.Context) error {
	return mongoClient.Close(ctx)
}
