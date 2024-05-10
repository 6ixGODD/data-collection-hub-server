package mongo

import (
	"context"

	"github.com/qiniu/qmgo"
)

type Mongo struct {
	MongoClient   *qmgo.Client
	MongoDatabase *qmgo.Database
	mongoConfig   *Config
	ctx           context.Context
}

type Config struct {
	qmgoConfig   *qmgo.Config
	pingTimeout  int64
	databaseName string
}

var mongoInstance *Mongo // Singleton

func New(ctx context.Context, config *qmgo.Config, PingTimeout int64, databaseName string) (m *Mongo, err error) {
	if mongoInstance != nil {
		return mongoInstance, nil
	}
	m = &Mongo{
		mongoConfig: &Config{
			qmgoConfig:   config,
			pingTimeout:  PingTimeout,
			databaseName: databaseName,
		},
		ctx: ctx,
	}
	if err := m.Init(); err != nil {
		return nil, err
	}
	mongoInstance = m
	return m, nil
}

func (m *Mongo) Init() error {
	client, err := qmgo.NewClient(m.ctx, m.mongoConfig.qmgoConfig)
	if err != nil {
		return err
	}
	if err = client.Ping(m.mongoConfig.pingTimeout); err != nil {
		return err
	}
	m.MongoClient = client
	m.MongoDatabase = client.Database(m.mongoConfig.databaseName)
	return nil
}

func (m *Mongo) GetClient() (client *qmgo.Client, err error) {
	if m.MongoClient == nil || m.MongoClient.Ping(m.mongoConfig.pingTimeout) != nil {
		if err = m.Init(); err != nil {
			return nil, err
		}
	}
	return m.MongoClient, nil
}

func (m *Mongo) GetDatabase() (database *qmgo.Database, err error) {
	if m.MongoDatabase == nil || m.MongoClient.Ping(m.mongoConfig.pingTimeout) != nil {
		if err = m.Init(); err != nil {
			return nil, err
		}
	}
	return m.MongoDatabase, nil

}

func (m *Mongo) Close(ctx context.Context) error {
	return m.MongoClient.Close(ctx)
}
