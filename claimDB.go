package main

import "encoding/csv"

type ClaimDB interface {
	GetClaims(input string) (*csv.Reader, error)
}
