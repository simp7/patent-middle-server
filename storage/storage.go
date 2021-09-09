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

	for numbers := range s.source.GetNumbers(input) {
		for number := range numbers {
			if data, err := s.getClaimsEach(number); err == nil {
				result.Append(data)
			}
		}
	}

	return result

}

func (s *storage) getClaimsEach(number string) (result model.CSVUnit, err error) {

	tuple, ok := s.cache.Find(number)
	if !ok {
		tuple = s.source.GetClaims(number)
		err = s.cache.Register(tuple)
	}

	result = tuple.CSVRow()

	return

}
