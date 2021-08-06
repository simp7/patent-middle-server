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

func (k *kipris) GetClaims(input string) model.CSVGroup {

	data := make([]model.CSVUnit, 0)

	for number := range k.source.GetNumbers(input) {

		tuple, ok := k.cache.Find(number)

		if !ok {
			tuple = k.source.GetClaims(number)
			err := k.cache.Register(tuple)
			if err != nil {
				continue
			}
		}

		data = append(data, tuple.Process())

	}

	return model.NewCSV(time.Now().String()+"@"+input, data)

}