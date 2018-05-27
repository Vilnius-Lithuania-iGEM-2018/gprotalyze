// Package plugins defines an interface and protocol to
// for data supply and storage options and possibly other plugins
package plugins

// PluginContext describes an analysis plugin
type PluginContext struct {
	Name     string
	FilePath string
	Version  string
}

// Plugin contains language specific details
// and whatnot, handles all aspects of it.
type Plugin interface {
	Run() error
	GetContext() *PluginContext
}
