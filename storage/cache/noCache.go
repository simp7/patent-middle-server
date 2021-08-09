package cache

import "github.com/simp7/patent-middle-server/storage"

type empty struct{}

func (e *empty) Find(_ string) (storage.ClaimTuple, bool) {
	return storage.ClaimTuple{}, false
}

func (e *empty) Register(_ storage.ClaimTuple) error {
	return nil
}

func Nocache() storage.Cache {
	return &empty{}
}
