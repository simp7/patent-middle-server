package claimDB

import (
	"github.com/simp7/patent-middle-server/model"
	"strings"
)

type claim struct {
	title  string
	claims []string
}

func (c claim) Tokenize() model.CSVUnit {
	return model.CSVUnit{Key: c.title, Value: strings.Join(c.claims, "\n")}
}
