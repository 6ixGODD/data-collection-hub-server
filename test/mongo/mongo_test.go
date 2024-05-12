package mongo__test

import (
	"context"
	"testing"

	"data-collection-hub-server/internal/pkg/config"
	"data-collection-hub-server/pkg/mongo"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestMongo(t *testing.T) {
	ctx := context.TODO()

	cfg, err := config.New()
	assert.NoError(t, err)
	t.Logf("QmgoConfig: %+v", cfg.MongoConfig.GetQmgoConfig())
	t.Logf("PingTimeoutS: %d", cfg.MongoConfig.PingTimeoutS)
	t.Logf("Database: %s", cfg.MongoConfig.Database)
	mongoInstance, err := mongo.New(
		ctx, cfg.MongoConfig.GetQmgoConfig(), cfg.MongoConfig.PingTimeoutS, cfg.MongoConfig.Database,
	)
	assert.NoError(t, err)
	assert.NotNil(t, mongoInstance)

	client, err := mongoInstance.GetClient()
	assert.NoError(t, err)
	assert.NotNil(t, client)

	db, err := mongoInstance.GetDatabase()
	assert.NoError(t, err)
	assert.NotNil(t, db)

	collection := db.Collection("test")
	assert.NotNil(t, collection)
	doc := bson.M{
		"test": "test",
	}
	one, err := collection.InsertOne(ctx, doc)
	assert.NoError(t, err)
	assert.NotNil(t, one)
	t.Logf("inserted id: %v", one.InsertedID)

	filter := bson.M{
		"test": "test",
	}
	var result bson.M
	err = collection.Find(ctx, filter).One(&result)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	t.Logf("result: %+v", result)

	err = collection.RemoveId(ctx, one.InsertedID)
	assert.NoError(t, err)
}
