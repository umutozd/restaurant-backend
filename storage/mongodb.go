package storage

type Storage interface{}

type storage struct {
	dbName string
}

// TODO: implement this
func NewStorage(dbURL string) Storage {
	return &storage{}
}
