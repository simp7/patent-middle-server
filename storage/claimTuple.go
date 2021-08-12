package storage

import (
	"github.com/simp7/patent-middle-server/model"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
)

type ClaimTuple struct {
	ApplicationNumber string   `bson:"_id"`
	Name              string   `bson:"name"`
	Claims            []string `bson:"claims"`
}

func (c ClaimTuple) CSVRow() model.CSVUnit {
	claims := strings.Join(c.Claims, "\n")
	return model.CSVUnit{Key: c.Name, Value: claims}
}

func (c ClaimTuple) BSON() bson.D {
	return bson.D{{"_id", c.ApplicationNumber}, {"name", c.Name}, {"claims", c.Claims}}
}
