package cache

import (
	"context"
	"github.com/simp7/patent-middle-server/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type mongoDB struct {
	collection *mongo.Collection
}

func Mongo(url string) (storage.Cache, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}

	db := client.Database("Patent")
	collection := db.Collection("claim")

	return &mongoDB{collection}, err

}

func (m *mongoDB) Find(applicationNumber string) (tuple storage.ClaimTuple, ok bool) {

	ok = false
	dbResult := m.collection.FindOne(context.TODO(), bson.D{{"applicationNumber", applicationNumber}})

	if dbResult.Err() == nil {
		err := dbResult.Decode(&tuple)
		if err == nil {
			ok = true
		}
	}

	return

}

func (m *mongoDB) Register(tuple storage.ClaimTuple) error {
	_, err := m.collection.InsertOne(context.TODO(), tuple.BSON())
	return err
}
