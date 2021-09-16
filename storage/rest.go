package storage

type Rest interface {
	GetNumbers(formula string, outCh chan<- string)
	GetClaims(string) Data
}
