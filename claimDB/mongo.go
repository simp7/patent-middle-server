package claimDB

import (
	"context"
	"github.com/simp7/patent-middle-server/model"
	"github.com/simp7/patent-middle-server/model/formula"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
	"time"
)

var once sync.Once
var instance *mongodb

type mongodb struct {
	claim *mongo.Collection
}

func (m *mongodb) GetClaims(input string) ([]model.CSVUnit, error) {
	f := formula.Interpret(input)
	f.Excluded.String()
	return nil, nil
}

func Mongo() (*mongodb, error) {

	var err error

	once.Do(func() {

		instance = new(mongodb)
		var client *mongo.Client

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		client, err = mongo.Connect(ctx, options.Client().ApplyURI("localhost"))
		if err != nil {
			return
		}

		db := client.Database("Patent")
		instance.claim = db.Collection("claim")

	})

	return instance, err

}
