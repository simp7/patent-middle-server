package claimDB

import (
	"github.com/simp7/patent-middle-server/model"
	"os"
)

type testDB struct {
}

func Test() *testDB {
	return &testDB{}
}

func (t *testDB) GetClaims(input string) ([]model.CSVUnit, error) {
	_, err := os.Open(input + ".csv")
	return nil, err
}
