package storage

import (
	"github.com/simp7/patent-middle-server/model"
	"os"
)

type testDB struct {
}

func Test() *testDB {
	return &testDB{}
}

func (t *testDB) GetClaims(input string) *model.CSVGroup {
	os.Open(input + ".csv")
	return &model.CSVGroup{}
}
