package main

import (
	"github.com/simp7/patent-middle-server/model"
)

type ClaimDB interface {
	GetClaims(input string) ([]model.CSVUnit, error)
}
