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
	dataChan := make(chan Data)

	go s.source.GetNumbers(input, numberChan)
	go s.searchClaims(numberChan, dataChan)
	s.processCSV(result, dataChan)

	return result

}

func (s *storage) searchClaims(inCh <-chan string, outCh chan<- Data) {

	for number := range inCh {

		data, ok := s.cache.Find(number)
		if !ok {
			data = s.source.GetClaims(number)
			s.cache.Register(data)
		}

		outCh <- data

	}

	close(outCh)

}

func (s *storage) processCSV(group *model.CSVGroup, inCh <-chan Data) {
	for data := range inCh {
		row := data.CSVRow()
		group.Append(row)
	}
}
