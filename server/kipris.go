package server

import "encoding/csv"

type Kipris interface {
	GetClaims(input string) *csv.Reader
}
