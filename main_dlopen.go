package main

/*
#cgo LDFLAGS: -ldl
#include <dlfcn.h>
typedef void (*voidfun)();
void caller(voidfun f) { f(); }
*/
import "C"

func main() {
	h := C.dlopen(C.CString("./libdynamic.so"), C.RTLD_LAZY)
	f := C.dlsym(h, C.CString("function"))
	C.caller(C.voidfun(f))
}
