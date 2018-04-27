package store

type DatabaseItem struct {
	DataType string
	Id string
	Data interface{}
}

type Store interface {
	Store(item DatabaseItem) error
	BulkStore(items []DatabaseItem) error
	Get(id string) (DatabaseItem, error)
	Update(item DatabaseItem) error
	MassUpdate(items []DatabaseItem) error
	Query(query string) ([]DatabaseItem, error)
}