package claimDB

import (
	"github.com/simp7/patent-middle-server/model"
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
