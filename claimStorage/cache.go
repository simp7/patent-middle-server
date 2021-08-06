package claimStorage

type Cache interface {
	Find(applicationNumber string) (ClaimTuple, bool)
	Register(tuple ClaimTuple) error
}
