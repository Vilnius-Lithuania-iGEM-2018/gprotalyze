#include "go-python.h"

extern void cgoPythonSaysHi();

PyMethodDef* initMethDef(size_t num) {
    return (PyMethodDef*) malloc(num * sizeof(PyMethodDef));
}

void cPythonSaysHi() {
    printf("%s", "C code says hi!");
    cgoPythonSaysHi();
}

void setMethDef(PyMethodDef *definitions, size_t item, char *name, PyObject* (*pFunc)(PyObject*, PyObject*), int flags, char *doc) {
	definitions[item].ml_name = name;
	definitions[item].ml_meth = pFunc;
	definitions[item].ml_flags = flags;
	definitions[item].ml_doc = doc;
}

PyObject* InitGprotalyzeModule(){
	PyMethodDef *methods = initMethDef(1);
	setMethDef(methods, 0, "run_hello", (PyCFunction)cPythonSaysHi, METH_VARARGS, "Execute a greeting.");
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