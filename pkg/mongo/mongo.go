package mongo

import (
	"context"

	"github.com/qiniu/qmgo"
)

type Mongo struct {
	MongoClient   *qmgo.Client
	MongoDatabase *qmgo.Database
	mongoConfig   *qmgo.Config
	pingTimeout   int64
	databaseName  string
	ctx           context.Context
}

var mongoInstance *Mongo // Singleton

func New(ctx context.Context, config *qmgo.Config, PingTimeout int64, databaseName string) (m *Mongo, err error) {
	if mongoInstance != nil {
		return mongoInstance, nil
	}
	m = &Mongo{
		mongoConfig:  config,
		pingTimeout:  PingTimeout,
		databaseName: databaseName,
		ctx:          ctx,
	}
	if err := m.Init(); err != nil {
		return nil, err
	}
	mongoInstance = m
	return m, nil
}

func (m *Mongo) Init() error {
	client, err := qmgo.NewClient(m.ctx, m.mongoConfig)
	if err != nil {
		return err
	}
	if err = client.Ping(m.pingTimeout); err != nil {
		return err
	}
	m.MongoClient = client
	m.MongoDatabase = client.Database(m.databaseName)
	return nil
}

func (m *Mongo) GetClient() (client *qmgo.Client, err error) {
	if m.MongoClient == nil || m.MongoClient.Ping(m.pingTimeout) != nil {
		if err = m.Init(); err != nil {
			return nil, err
		}
	}
	return m.MongoClient, nil
}

func (m *Mongo) GetDatabase() (database *qmgo.Database, err error) {
	if m.MongoDatabase == nil || m.MongoClient.Ping(m.pingTimeout) != nil {
		if err = m.Init(); err != nil {
			return nil, err
		}
	}
	return m.MongoDatabase, nil

}

func (m *Mongo) Close(ctx context.Context) error {
	return m.MongoClient.Close(ctx)
}
