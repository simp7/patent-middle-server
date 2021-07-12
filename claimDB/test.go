package claimDB

import (
	"encoding/csv"
	"os"
)

type testDB struct {
}

func Test() *testDB {
	return &testDB{}
}

func (t *testDB) GetClaims(input string) (*csv.Reader, error) {
	file, err := os.Open(input + ".csv")
	return csv.NewReader(file), err
}
