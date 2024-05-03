package mongo

import (
	"context"

	"github.com/qiniu/qmgo"
)

type Database struct {
	MongoClient   *qmgo.Client
	MongoDatabase *qmgo.Database
	MongoConfig   *qmgo.Config
	PingTimeout   int64
	DatabaseName  string
	Ctx           context.Context
}

func New(ctx context.Context, config *qmgo.Config, PingTimeout int64, databaseName string) (c *Database, err error) {
	c = &Database{
		MongoConfig:  config,
		PingTimeout:  PingTimeout,
		DatabaseName: databaseName,
		Ctx:          ctx,
	}
	if err := c.Init(); err != nil {
		return nil, err
	}
	return c, nil
}

func (d *Database) Init() error {
	client, err := qmgo.NewClient(d.Ctx, d.MongoConfig)
	if err != nil {
		return err
	}
	if err = client.Ping(d.PingTimeout); err != nil {
		return err
	}
	d.MongoClient = client
	d.MongoDatabase = client.Database(d.DatabaseName)
	return nil
}

func (d *Database) GetClient() (client *qmgo.Client, err error) {
	if d.MongoClient == nil || d.MongoClient.Ping(d.PingTimeout) != nil {
		if err = d.Init(); err != nil {
			return nil, err
		}
	}
	return d.MongoClient, nil
}

func (d *Database) GetDatabase() (database *qmgo.Database, err error) {
	if d.MongoDatabase == nil || d.MongoClient.Ping(d.PingTimeout) != nil {
		if err = d.Init(); err != nil {
			return nil, err
		}
	}
	return d.MongoDatabase, nil

}

func (d *Database) Close(ctx context.Context) error {
	return d.MongoClient.Close(ctx)
}
