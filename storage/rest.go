package storage

type Rest interface {
	GetNumbers(formula string) chan chan string
	GetClaims(applicationNumber string) ClaimTuple
}
