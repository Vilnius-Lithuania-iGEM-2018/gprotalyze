#include "go-python.h"
#include <stdio.h>

extern void cgoPythonSaysHi();

PyObject* cPythonSaysHi(PyObject *argc, PyObject *args) {
    printf("%s\n", "CPython says hi!!");
    cgoPythonSaysHi();
    return NULL;
}

PyObject* InitGprotalyzeModule(){
    PyMethodDef *methods = (PyMethodDef*)malloc(2 * sizeof(PyMethodDef));
    methods[0].ml_name = "run_hello";
    methods[0].ml_meth = cPythonSaysHi;
    methods[0].ml_flags = METH_VARARGS;
    methods[0].ml_doc = "Execute a greeting.";
    methods[1].ml_name = NULL;
    methods[1].ml_meth = NULL;
    methods[1].ml_flags = 0;
    methods[1].ml_doc = NULL;
	//PyMethodDef methods[] = {
	//    {"run_hello", cPythonSaysHi, METH_VARARGS, "Execute a greeting."},
	//    {NULL, NULL, 0, NULL}
	//};
	return Py_InitModule("gprotalyze", methods);
}

int runHello(PyObject *pythonPlugin) {
    printf("%s\n", "Run hello started!");
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