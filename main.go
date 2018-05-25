package main

import (
	"context"
	"fmt"
	"github.com/Vilnius-Lithuania-iGEM-2018/gprotalyze/plugins"
	"github.com/Vilnius-Lithuania-iGEM-2018/gprotalyze/store"
)

func main() {
	// /Users/lukas.praninskas/Documents/Projects/golang/src/github.com/Vilnius-Lithuania-iGEM-2018/gprotalyze
	plugin, err := plugins.LoadPythonPlugin("plugin")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Module: %s\n", plugin)
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
