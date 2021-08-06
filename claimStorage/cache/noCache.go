package cache

import "github.com/simp7/patent-middle-server/claimStorage"

type empty struct{}

func (e *empty) Find(_ string) (claimStorage.ClaimTuple, bool) {
	return claimStorage.ClaimTuple{}, false
}

func (e *empty) Register(_ claimStorage.ClaimTuple) error {
	return nil
}

func Nocache() claimStorage.Cache {
	return &empty{}
}
