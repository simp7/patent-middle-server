package storage

type Rest interface {
	GetNumbers(formula string) chan string
	GetClaims(applicationNumber string) ClaimTuple
}
