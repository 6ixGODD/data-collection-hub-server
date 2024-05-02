package mongo

import (
	"context"

	"github.com/qiniu/qmgo"
)

type Database struct {
	MongoClient   *qmgo.Client
	MongoDatabase *qmgo.Database
	MongoConfig   *qmgo.Config
	DatabaseName  string
}

func New(ctx context.Context, config *qmgo.Config, databaseName string) (c *Database, err error) {
	c = &Database{
		MongoConfig:  config,
		DatabaseName: databaseName,
	}
	if err := c.Init(ctx); err != nil {
		return nil, err
	}
	return c, nil
}

func (d *Database) Init(ctx context.Context) error {
	client, err := qmgo.NewClient(ctx, d.MongoConfig)
	if err != nil {
		return err
	}
	d.MongoClient = client
	d.MongoDatabase = client.Database(d.DatabaseName)
	return nil
}

func (d *Database) GetClient(ctx context.Context) (client *qmgo.Client, err error) {
	if d.MongoClient == nil {
		if err = d.Init(ctx); err != nil {
			return nil, err
		}
	}
	return d.MongoClient, nil
}

func (d *Database) GetDatabase(ctx context.Context) (database *qmgo.Database, err error) {
	if d.MongoDatabase == nil {
		if err = d.Init(ctx); err != nil {
			return nil, err
		}
	}
	return d.MongoDatabase, nil

}

func (d *Database) Close(ctx context.Context) error {
	return d.MongoClient.Close(ctx)
}
