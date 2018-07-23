#include "go-python.h"
#include <stdio.h>

extern void cgoPythonSaysHi();

PyObject* cPythonSaysHi(PyObject *argn, PyObject *args) {
    printf("%s!", "CPython says hi!!");
    cgoPythonSaysHi();
    return NULL;
}

PyObject* InitGprotalyzeModule(){
	PyMethodDef methods[] = {
	    {"run_hello", cPythonSaysHi, METH_VARARGS, "Execute a greeting."},
	    {NULL, NULL, 0, NULL}
	};
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