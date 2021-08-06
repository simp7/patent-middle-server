package storage

import (
	"github.com/simp7/patent-middle-server/model"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
)

type ClaimTuple struct {
	ApplicationNumber string
	Name              string
	Claims            []string
}

func (c ClaimTuple) Process() model.CSVUnit {
	claims := strings.Join(c.Claims, "\n")
	return model.CSVUnit{Key: c.Name, Value: claims}
}

func (c ClaimTuple) BSON() bson.D {
	return bson.D{{"_id", c.ApplicationNumber}, {"name", c.Name}, {"claim", bson.A{c.Claims}}}
}
