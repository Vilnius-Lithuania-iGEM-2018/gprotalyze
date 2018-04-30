package main

import (
	"github.com/Vilnius-Lithuania-iGEM-2018/gprotalyze/store"
	"context"
)

func main() {
	elastic, err := store.NewElasticStore(context.Background())
	if err != nil {
		panic(err)
	}

	elastic.Store(store.Document{
		Id: "1",
		DataType: "analysis",
		Data: struct {
			Data []string `json:"data"`
			Tags []string `json:"tags"`
		}{
			Data: []string{"data1", "data2", "data3"},
			Tags: []string{"tag1", "tag2", "tag3"},
		},
	})
}
