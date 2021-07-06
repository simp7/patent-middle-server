package server

import (
	"strconv"
)

type unit struct {
	Word       string
	Similarity float64
}

func Unit(word string, similarity string) (unit, error) {
	simValue, err := strconv.ParseFloat(similarity, 64)
	return unit{word, simValue}, err
}
