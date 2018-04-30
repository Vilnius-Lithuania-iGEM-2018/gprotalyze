// Package plugins defines an interface and protocol to
// for data supply and storage options and possibly other plugins
package plugins

type PluginContext struct {
	Name     string
	FilePath string
	Version  string
}

type Plugin interface {
	Run() error
	GetContext() *PluginContext
}
