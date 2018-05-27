package plugins

import (
	"errors"
	"github.com/sbinet/go-python"
	"github.com/sirupsen/logrus"
	"runtime"
)

func init() {
	type str struct {
		value bool
	}

	err := python.Initialize()
	if err != nil {
		panic(err)
	}
	python.PySys_SetPath("/Users/lukas.praninskas/Documents/Projects/golang/src/github.com/Vilnius-Lithuania-iGEM-2018/gprotalyze/examples")

	args := &str{value: true}
	go func() {
		runtime.SetFinalizer(args, func(finalize *str) {
			python.Finalize()
		})
	}()
}

// LoadPythonPlugin loads a python plugin from python path, according to the filename
func LoadPythonPlugin(filename string) (*PythonPlugin, error) {
	loggerInstance := logrus.New()
	module := python.PyImport_Import(python.PyString_FromString(filename))
	loggerInstance.WithField("module", module).Debug()
	if python.PyErr_Occurred() != nil {
		python.PyErr_Print()
		loggerInstance.WithFields(logrus.Fields{
			"importFile": filename,
		}).Debug("loaded module")
		return nil, errors.New("cannot load python module")
	} else {
		return &PythonPlugin{
			log:          loggerInstance,
			pythonModule: module,
			context: PluginContext{
				Name:     filename,
				FilePath: filename,
				Version:  "1",
			},
		}, nil
	}
}

// PythonPlugin is an inherited struct from the generic Plugin
type PythonPlugin struct {
	pythonModule *python.PyObject
	context      PluginContext
	log          *logrus.Logger
}

// Run performs a run on the plugin
func (plugin PythonPlugin) Run() error {
	python.PyErr_Clear()
	pFunc := plugin.pythonModule.GetAttrString("hello")
	if python.PyErr_Occurred() == nil || pFunc.Check_Callable() {
		pFunc.CallObject(python.PyTuple_New(0))
	} else {
		return errors.New("the module loaded is not callable")
	}
	return nil
}

// GetContext returns the context of the plugin
func (plugin PythonPlugin) GetContext() *PluginContext {
	return &plugin.context
}
