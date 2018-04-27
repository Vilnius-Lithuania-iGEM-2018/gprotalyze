package store

import (
	"github.com/sirupsen/logrus"
	"github.com/olivere/elastic"

	"context"
	"fmt"
	"errors"
)

func NewElasticStore(context context.Context) (*ElasticStore, error) {
	storeClient, err := elastic.NewClient()
	if err != nil {
		return nil, err
	}
	store := ElasticStore{
		client:  storeClient,
		context: context,
		logger: logrus.New(),
	}
	return &store, err
}

type ElasticStore struct {
	client *elastic.Client
	context context.Context
	logger *logrus.Logger
}

func (store ElasticStore) Store(item DatabaseItem) error {
	response, err := store.client.Index().
		Index("gprotalyze").
		Type(item.DataType).
		Id(item.Id).
		BodyJson(item.Data).
		Do(store.context)
	store.logger.WithField("result", response.Result).Debug()
	return err
}

func (store ElasticStore) BulkStore(items []DatabaseItem) error {
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
		return errors.New(
			fmt.Sprintf(
				"Expected %d succeeded, got %d",
				bulk.NumberOfActions(),
				len(response.Succeeded())))
	}
	store.logger.WithFields(logrus.Fields{
		"launched": bulk.NumberOfActions(),
		"succeeded": response.Succeeded(),
		"failed": response.Failed(),
		"time": response.Took,
	}).Debug("elasticsearch BulkStore")
	return nil
}

func (store ElasticStore) Get(id string) (DatabaseItem, error) {
	response, err := store.client.Get().Index("gprotalyze").Do(store.context)

	if err != nil {
		return DatabaseItem{}, err
	}

	if !response.Found {
		store.logger.WithField("id", id).Debug("Not found")
		return DatabaseItem{}, errors.New("document not found")
	}

	return DatabaseItem{
		DataType: response.Type,
		Id:       response.Id,
		Data:     response.Fields,
	}, nil
}

func (store ElasticStore) Update(item DatabaseItem) error {
	return errors.New("unimplemented")
}

func (store ElasticStore) MassUpdate(items []DatabaseItem) error {
	return errors.New("unimplemented")
}

func (store ElasticStore) Query(query string) ([]DatabaseItem, error) {
	return nil, errors.New("unimplemented")
}