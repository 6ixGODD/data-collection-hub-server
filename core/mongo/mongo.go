package mongo

import (
	"github.com/qiniu/qmgo"
)

var (
	client *qmgo.QmgoClient
)

func InitMongo() {

}

func GetMongoClient() *qmgo.QmgoClient {
	return client
}
