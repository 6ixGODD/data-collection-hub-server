package mongo

import (
	"context"

	"github.com/qiniu/qmgo"
)

var client *qmgo.QmgoClient

func InitMongoClient(options *qmgo.Config) *qmgo.QmgoClient {
}

func GetMongoClient(coll string) *qmgo.QmgoClient {
	var err error
	client, err = qmgo.Open(context.Background(), &qmgo.Config{Uri: "mongodb://localhost:27017", Database: "test", Coll: coll})
	if err != nil {
		return nil
	} else {
		return client
	}
}

func ReleaseMongo(ctx context.Context) (err error) {
	err = client.Close(ctx)
	if err != nil {
		return err
	} else {
		return nil
	}
}
