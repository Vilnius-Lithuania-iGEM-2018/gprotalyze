package plugins
import "C"

// #cgo pkg-config: python-2.7
// #include <Python.h>
//
// PyMethodDef* initMethDef(size_t num) {
// 		return (PyMethodDef*) malloc(num * sizeof(PyMethodDef));
// }
//
// void setMethDef(PyMethodDef *definitions, size_t item, char *name, PyCFunction pFunc, int flags, char *doc) {
//		definitions[item].ml_name = name;
//		definitions[item].ml_meth = pFunc;
//		definitions[item].ml_flags = flags;
//		definitions[item].ml_doc = doc;
// }
import "C"

import (
	"github.com/sirupsen/logrus"
	"runtime"
	"errors"
	"unsafe"
)

var initializedModule *C.PyObject = nil

// PythonPlugin is an inherited struct from the generic Plugin
type PythonPlugin struct {
	pythonModule *C.PyObject
	context      PluginContext
	log          *logrus.Logger
}

// LoadPythonPlugin loads a python plugin from python path, according to the filename
func LoadPythonPlugin(filename string) (*PythonPlugin, error) {
	loggerInstance := logrus.New()
	//module := python.PyImport_Import(python.PyString_FromString(filename))
	//loggerInstance.WithField("module", module).Debug()
	//if python.PyErr_Occurred() != nil {
	//	python.PyErr_Print()
	//	loggerInstance.WithFields(logrus.Fields{
	//		"importFile": filename,
	//	}).Debug("loaded module")
	//	return nil, errors.New("cannot load python module")
	//}
	return &PythonPlugin{
		log:          loggerInstance,
		pythonModule: nil, //module
		context: PluginContext{
			Name:     filename,
			FilePath: filename,
			Version:  "1",
		},
	}, nil
}

func initMethDef(num uint) *C.PyMethodDef {
	return C.initMethDef(num)
}

func setMethDef(definitions *C.PyMethodDef, item uint, name string, pFunc unsafe.Pointer, flags int, doc string) {
	C.setMethDef(definitions, item, C.CString(name), pFunc, flags, C.CString(doc))
}

func init() {
	type str struct {
		value bool
	}

	C.Py_InitializeEx(0)
	err := C.PyErr_Occurred()
	if err != nil {
		panic(err)
	}
	C.PySys_SetPath(C.CString("/Users/lukas.praninskas/Documents/Projects/golang/src/github.com/Vilnius-Lithuania-iGEM-2018/gprotalyze/examples"))

	args := &str{value: true}
	go func() {
		runtime.SetFinalizer(args, func(finalize *str) {
			C.Py_Finalize()
		})
	}()

	methodDef := initMethDef(1)
	setMethDef(methodDef, 0, "run_hello", nil, C.METH_VARARGS, "Execute a greeting.")
	initializedModule = C.Py_InitModule(C.CString("gprotalyze"), methodDef)
}

func int2bool(i C.int) bool {
	switch i {
	case -1:
		return false
	case 0:
		return false
	case 1:
		return true
	default:
		return true
	}
	return false
}

// Run performs a run on the plugin
func (plugin PythonPlugin) Run() error {
	C.PyErr_Clear()
	pFunc := C.PyObject_GetAttrString(plugin.pythonModule, C.CString("hello"))
	if C.PyErr_Occurred() == nil || int2bool(C.PyCallable_Check(pFunc)) {
		C.PyObject_CallObject(pFunc, C.PyTuple_New(C.Py_ssize_t(0)))
	} else {
		return errors.New("the module loaded is not callable")
	}
	return nil
}

