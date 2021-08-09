package storage

import (
	"github.com/simp7/patent-middle-server/model"
	"sync"
	"time"
)

var instance *kipris
var once sync.Once

type kipris struct {
	cache  Cache
	source Rest
}

func New(source Rest, cacheDB Cache) *kipris {

	once.Do(func() {
		instance = &kipris{
			cacheDB,
			source,
		}
	})

	return instance

}

func (k *kipris) GetClaims(input string) *model.CSVGroup {

	result := model.NewCSV(time.Now().String() + "@" + input)

	for numbers := range k.source.GetNumbers(input) {
		a := 1
		for number := range numbers {

			data, err := k.getClaimsEach(number)
			if err != nil {
				result.Append(data)
			}
			a++

		}

	}

	return result

}

func (k *kipris) getClaimsEach(number string) (result model.CSVUnit, err error) {

	tuple, ok := k.cache.Find(number)
	if !ok {
		tuple = k.source.GetClaims(number)
		err = k.cache.Register(tuple)
	}

	result = tuple.Process()

	return

}
