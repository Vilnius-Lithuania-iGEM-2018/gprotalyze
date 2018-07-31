package plugins

import "C"

// #cgo pkg-config: python-2.7
// #cgo CFLAGS: -DPNG_DEBUG=1 -Og -g
// #include "go-python.h"
import "C"

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
)

// PythonPlugin is an inherited struct from the generic Plugin
type PythonPlugin struct {
	pythonModule *C.PyObject    // pythonModule is the gprotalyze package that is defined in CPython
	pythonPlugin *C.PyObject    // pythonPlugin is the plugin that gprotalyze application loads and runs
	context      PluginContext  // context contains all the misc data about the plugin
	log          *logrus.Logger // log is the instance that this concrete plugin uses to report errors
}

var initializedModule *C.PyObject = nil

// LoadPythonPlugin loads a python plugin from python path, according to the filename
func LoadPythonPlugin(filename string) (*PythonPlugin, error) {
	C.PyErr_Clear()
	plugin := C.PyImport_ImportModule(C.CString(filename))
	pyErr := C.PyErr_Occurred()
	if pyErr != nil {
		C.PyErr_Print()
		return nil, errors.New("cannot import plugin")
	}

	return &PythonPlugin{
		log:          logrus.New(),
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

	C.PySys_SetPath(C.CString(cwd + "/python-plugins"))

	args := &str{value: true}
	go func() {
		runtime.SetFinalizer(args, func(finalize *str) {
			C.Py_Finalize()
		})
	}()

	initializedModule = C.InitGprotalyzeModule()

	if C.PyErr_Occurred() != nil {
		C.PyErr_Print()
		panic("Failed to load module")
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
