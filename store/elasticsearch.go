package store

import (
	"github.com/olivere/elastic"
	"github.com/sirupsen/logrus"

	"context"
	"errors"
	"fmt"
)

// NewElasticStore creates and returns elastic client
func NewElasticStore(context context.Context) (Store, error) {
	logInstance := logrus.New()
	storeClient, err := elastic.NewSimpleClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		logInstance.Error(err)
		return nil, err
	}
	store := ElasticStore{
		client:  storeClient,
		context: context,
		logger:  logInstance,
	}

	return store, err
}

// ElasticStore is the elastic client implementation
type ElasticStore struct {
	client  *elastic.Client
	context context.Context
	logger  *logrus.Logger
}

// Store stores a single document in a single network operation
func (store ElasticStore) Store(item Document) error {
	response, err := store.client.Index().
		Index("gprotalyze").
		Type(item.DataType).
		Id(item.Id).
		BodyJson(item.Data).
		Do(store.context)
	store.logger.WithField("result", response.Result).Debug()
	return err
}

// BulkStore stores multiple documents in one network operation
func (store ElasticStore) BulkStore(items []Document) error {
	bulk := store.client.Bulk()
	for _, item := range items {
		bulk.Add(elastic.NewBulkIndexRequest().
			Index("gprotalyze").
			Id(item.Id).
			Type(item.DataType).
			Doc(item.Data))
	}
	response, err := bulk.Do(store.context)
	if err != nil {
		return err
	}
	if len(response.Succeeded()) != bulk.NumberOfActions() {
		return fmt.Errorf(
			"Expected %d succeeded, got %d",
			bulk.NumberOfActions(),
			len(response.Succeeded()),
		)
	}
	store.logger.WithFields(logrus.Fields{
		"launched":  bulk.NumberOfActions(),
		"succeeded": response.Succeeded(),
		"failed":    response.Failed(),
		"time":      response.Took,
	}).Debug("elasticsearch BulkStore")
	return nil
}

// Get gets a document according to it's id
func (store ElasticStore) Get(id string) (Document, error) {
	response, err := store.client.Get().Index("gprotalyze").Do(store.context)

	if err != nil {
		return Document{}, err
	}

	if !response.Found {
		store.logger.WithField("id", id).Debug("Not found")
		return Document{}, errors.New("document not found")
	}

	return Document{
		DataType: response.Type,
		Id:       response.Id,
		Data:     response.Fields,
	}, nil
}

// Update just does an update operation
func (store ElasticStore) Update(item Document) error {
	return errors.New("unimplemented")
}

// MassUpdate updates many documents in a single operation
func (store ElasticStore) MassUpdate(items []Document) error {
	return errors.New("unimplemented")
}

// Query performs a search function and returns all of the found documents
func (store ElasticStore) Query(query string) ([]Document, error) {
	return nil, errors.New("unimplemented")
}
