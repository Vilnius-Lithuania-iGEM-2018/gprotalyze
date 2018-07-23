package main

import (
	"context"
	"github.com/Vilnius-Lithuania-iGEM-2018/gprotalyze/plugins"
	"github.com/Vilnius-Lithuania-iGEM-2018/gprotalyze/store"
)

func main() {
	plugin, err := plugins.LoadPythonPlugin("plugin")
	if err != nil {
		panic(err)
	}
	plugin.Run()

	elastic, err := store.NewElasticStore(context.Background())
	if err != nil {
		panic(err)
	}

	elastic.Store(store.Document{
		Id:       "1",
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
