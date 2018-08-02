package main

import (
	"context"
	"github.com/Vilnius-Lithuania-iGEM-2018/gprotalyze/plugins"
	"github.com/Vilnius-Lithuania-iGEM-2018/gprotalyze/store"
)

func main() {
	elastic, err := store.NewElasticStore(context.Background())
	if err != nil {
		panic(err)
	}

	plugin, err := plugins.LoadPythonPlugin("plugin", elastic)
	if err != nil {
		panic(err)
	}
	plugin.Run()
}
