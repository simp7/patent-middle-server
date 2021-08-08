package main

import (
	"github.com/simp7/patent-middle-server/model"
)

type Storage interface {
	GetClaims(input string) *model.CSVGroup
}
