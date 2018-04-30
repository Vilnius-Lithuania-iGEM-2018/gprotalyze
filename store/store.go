// Package store defines the interface of storing data.
// The method and place depends on implementation
package store

// Document is a single database document
// No schema past DataType and Id
type Document struct {
	DataType string
	Id       string
	Data     interface{}
}

// Store is the client interface for whatever storage mechanism implemented
type Store interface {

	// Store stores a single document into database
	Store(item Document) error

	// BulkStore stores multiple documents at once
	BulkStore(items []Document) error

	// Get returns an item according to it's id
	Get(id string) (Document, error)

	// Update finds the item by id and changes it's value
	Update(item Document) error

	// MassUpdate gets a lot of items and changes their value
	MassUpdate(items []Document) error

	// Query performs a search on existing documents in the database
	Query(query string) ([]Document, error)
}
