package cache

import (
	"context"
	"errors"
	"github.com/simp7/patent-middle-server/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
	"time"
)

var once sync.Once
var instance *mongoDB
var connectionError error = errors.New("could not connect to MongoDB")

type mongoDB struct {
	collection *mongo.Collection
}

func Mongo(config Config) (storage.Cache, error) {

	var err error
	once.Do(func() {

		var client *mongo.Client
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if client, err = mongo.Connect(ctx, options.Client().ApplyURI(config.URL)); err != nil {
			instance = nil
			return
		}

		db := client.Database(config.DBName)
		collection := db.Collection(config.CollectionName)

		instance = &mongoDB{collection: collection}

	})

	if instance == nil {
		return nil, connectionError
	}

	return instance, err

}

func (m *mongoDB) Find(applicationNumber string) (tuple storage.ClaimTuple, ok bool) {

	ok = false
	dbResult := m.collection.FindOne(context.TODO(), bson.D{{"_id", applicationNumber}})

	if dbResult.Err() == nil {
		if err := dbResult.Decode(&tuple); err == nil {
			ok = true
		}
	}

	return

}

func (m *mongoDB) Register(tuple storage.ClaimTuple) error {
	_, err := m.collection.InsertOne(context.TODO(), tuple.BSON())
	return err
}
