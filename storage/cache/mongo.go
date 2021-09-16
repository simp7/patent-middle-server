package cache

import (
	"context"
	"errors"
	"github.com/google/logger"
	"github.com/simp7/patent-middle-server/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
	"time"
)

var once sync.Once
var instance *mongoDB
var connectionError = errors.New("could not connect to MongoDB")

type mongoDB struct {
	collection *mongo.Collection
	*logger.Logger
}

func Mongo(config Config, lg *logger.Logger) (storage.Cache, error) {

	var err error
	once.Do(func() {

		var client *mongo.Client
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if client, err = mongo.Connect(ctx, options.Client().ApplyURI(config.URL)); err != nil {
			lg.Error(err)
			instance = nil
			return
		}

		db := client.Database(config.DBName)
		collection := db.Collection(config.CollectionName)

		instance = &mongoDB{collection: collection, Logger: lg}
		lg.Info("mongoDB has been loaded")

	})

	if instance == nil {
		return nil, connectionError
	}

	return instance, err

}

func (m *mongoDB) Find(applicationNumber string) (data storage.Data, ok bool) {

	ok = false
	dbResult := m.collection.FindOne(context.TODO(), bson.D{{"_id", applicationNumber}})

	if err := dbResult.Err(); err == nil {
		if err = dbResult.Decode(&data); err == nil {
			ok = true
			m.Info("cache has " + applicationNumber)
		}
	}

	return

}

func (m *mongoDB) Register(data storage.Data) (err error) {
	if _, err = m.collection.InsertOne(context.TODO(), data.BSON()); err != nil {
		m.Error(err)
		return
	}
	m.Info("register " + data.ApplicationNumber + " successfully")
	return
}
