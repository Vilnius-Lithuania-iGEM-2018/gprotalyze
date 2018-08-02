#include "go-python.h"

// Exported Go functions
extern void cgoPythonSaysHi();

// Python module defining functions
PyObject* cPythonSaysHi(PyObject*, PyObject*);
PyObject* storeDocument(PyObject*, PyObject*);

PyObject* InitGprotalyzeModule(){
    PyMethodDef *methods = (PyMethodDef*)malloc(3 * sizeof(PyMethodDef));
    methods[0].ml_name = "run_hello";
    methods[0].ml_meth = cPythonSaysHi;
    methods[0].ml_flags = METH_VARARGS;
    methods[0].ml_doc = "Execute a greeting.";
    methods[1].ml_name = "storeDocument";
    methods[1].ml_meth = storeDocument;
    methods[1].ml_flags = METH_VARARGS;
    methods[1].ml_doc = "Store some document to elastic";
    methods[2].ml_name = NULL;
    methods[2].ml_meth = NULL;
    methods[2].ml_flags = 0;
    methods[2].ml_doc = NULL;
	return Py_InitModule("gprotalyze", methods);
}

int runHello(PyObject *pythonPlugin) {
	PyErr_Clear();
	PyObject *attribute = PyObject_GetAttrString(pythonPlugin, "hello");
	if (PyErr_Occurred() == NULL || PyCallable_Check(attribute)) {
		PyObject_CallObject(attribute, PyTuple_New(0));
	} else {
		PyErr_SetString(PyExc_RuntimeError, "The attribute is not callable");
		return 1;
	}
	return 0;
}

PyObject* cPythonSaysHi(PyObject *argc, PyObject *args) {
    cgoPythonSaysHi();
    return Py_None;
}