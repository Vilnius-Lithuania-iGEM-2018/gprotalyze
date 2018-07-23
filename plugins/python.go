package plugins

// #cgo pkg-config: python-2.7
// #include "go-python.h"
import "C"

import (
	"github.com/sirupsen/logrus"
	"runtime"
	"errors"
	"os"
	"fmt"
)

// PythonPlugin is an inherited struct from the generic Plugin
type PythonPlugin struct {

	// pythonModule is the gprotalyze package that is defined in CPython
	pythonModule *C.PyObject

	// pythonPlugin is the plugin that gprotalyze application loads and runs
	pythonPlugin *C.PyObject

	// context contains all the misc data about the plugin
	context      PluginContext

	// log is the instance that this concrete plugin uses to report errors
	log          *logrus.Logger
}

var initializedModule *C.PyObject = nil
var loggerInstance = logrus.New()

// LoadPythonPlugin loads a python plugin from python path, according to the filename
func LoadPythonPlugin(filename string) (*PythonPlugin, error) {
	//module := python.PyImport_Import(python.PyString_FromString(filename))
	//loggerInstance.WithField("module", module).Debug()
	//if python.PyErr_Occurred() != nil {
	//	python.PyErr_Print()
	//	loggerInstance.WithFields(logrus.Fields{
	//		"importFile": filename,
	//	}).Debug("loaded module")
	//	return nil, errors.New("cannot load python module")
	//}

	plugin := C.PyImport_ImportModule(C.CString(filename))

	return &PythonPlugin{
		log:          loggerInstance,
		pythonModule: initializedModule,
		pythonPlugin: plugin,
		context: PluginContext{
			Name:     filename,
			FilePath: filename,
			Version:  "1",
		},
	}, nil
}

func init() {
	type str struct {
		value bool
	}

	C.Py_InitializeEx(0)
	pyErr := C.PyErr_Occurred()
	if pyErr != nil {
		panic(pyErr)
	}

	cwd, err := os.Getwd()
	if err != nil {
		panic(errors.New("unable to get current working directory"))
	}

	C.PySys_SetPath(C.CString(cwd + "python-plugins"))

	args := &str{value: true}
	go func() {
		runtime.SetFinalizer(args, func(finalize *str) {
			C.Py_Finalize()
		})
	}()

	initializedModule = C.InitGprotalyzeModule()
	loggerInstance.WithField("module", initializedModule).Debug()

	if initializedModule == nil {
		err := errors.New("python module could not be loaded")
		panic(err)
	}
}

// Run performs a run on the plugin
func (plugin PythonPlugin) Run() error {
	errorNum := C.runHello(plugin.pythonPlugin)
	var err error = nil
	switch errorNum {
	case 1:
		err = errors.New("cannot call plugin")
	default:
		if errorNum != 0 {
			err = errors.New("unknown error")
		}
	}
	return err
}

//export cgoPythonSaysHi
func cgoPythonSaysHi() {
	fmt.Printf("Hello from GO!")
}

