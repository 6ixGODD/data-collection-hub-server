package mongo_test

import (
	"testing"

	"data-collection-hub-server/test/wire"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestMongo(t *testing.T) {
	var (
		injector = wire.GetInjector()
		ctx      = injector.Ctx
		m        = injector.Mongo
		err      error
	)

	coll := m.MongoClient.Database(m.DatabaseName).Collection("test")
	assert.NotNil(t, coll)
	doc := bson.M{
		"test": "test",
	}
	one, err := coll.InsertOne(ctx, doc)
	assert.NoError(t, err)
	assert.NotNil(t, one)
	t.Logf("inserted id: %v", one.InsertedID)

	filter := bson.M{
		"test": "test",
	}
	var result bson.M
	err = coll.Find(ctx, filter).One(&result)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	t.Logf("result: %+v", result)

	err = coll.RemoveId(ctx, one.InsertedID)
	assert.NoError(t, err)
}
