package storage

import (
	"github.com/simp7/patent-middle-server/model"
	"sync"
)

var instance *storage
var once sync.Once

type storage struct {
	cache  Cache
	source Rest
}

func New(source Rest, cacheDB Cache) *storage {

	once.Do(func() {
		instance = &storage{
			cacheDB,
			source,
		}
	})

	return instance

}

func (s *storage) GetClaims(input string) *model.CSVGroup {

	result := model.NewCSV(input)

	numberChan := make(chan string)
	tupleChan := make(chan ClaimTuple)

	go s.source.GetNumbers(input, numberChan)
	go s.searchClaims(numberChan, tupleChan)
	s.processCSV(result, tupleChan)

	return result

}

func (s *storage) searchClaims(inCh <-chan string, outCh chan<- ClaimTuple) {

	for number := range inCh {

		tuple, ok := s.cache.Find(number)
		if !ok {
			tuple = s.source.GetClaims(number)
			s.cache.Register(tuple)
		}

		outCh <- tuple

	}

	close(outCh)

}

func (s *storage) processCSV(group *model.CSVGroup, inCh <-chan ClaimTuple) {
	for tuple := range inCh {
		row := tuple.CSVRow()
		group.Append(row)
	}
}
