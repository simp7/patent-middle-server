package storage

import (
	"github.com/simp7/patent-middle-server/model"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
)

type Data struct {
	ApplicationNumber string   `bson:"_id"`
	Name              string   `bson:"name"`
	Claims            []string `bson:"claims"`
}

func NewData(number string, name string, claims []string) Data {
	return Data{number, name, claims}
}

func (d Data) CSVRow() model.CSVUnit {
	claims := strings.Join(d.Claims, "\n")
	return model.CSVUnit{Key: d.Name, Value: claims}
}

func (d Data) BSON() bson.D {
	return bson.D{{"_id", d.ApplicationNumber}, {"name", d.Name}, {"claims", d.Claims}}
}
