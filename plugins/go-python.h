#ifndef C_GO_PYTHON_H_
#define C_GO_PYTHON_H_

#include <Python.h>

void cPythonSaysHi();

PyMethodDef* initMethDef(size_t);
void setMethDef(PyMethodDef*, size_t, char*, PyCFunction pFunc, int, char*);
PyObject* InitGprotalyzeModule();
int runHello(PyObject*);

#endif // C_GO_PYTHON_H_