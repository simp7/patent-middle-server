package storage

type Cache interface {
	Find(applicationNumber string) (Data, bool)
	Register(Data) error
}
