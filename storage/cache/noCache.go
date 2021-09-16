package cache

import "github.com/simp7/patent-middle-server/storage"

type empty struct{}

func (e *empty) Find(_ string) (storage.Data, bool) {
	return storage.Data{}, false
}

func (e *empty) Register(_ storage.Data) error {
	return nil
}

func Nocache() storage.Cache {
	return &empty{}
}
