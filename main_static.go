package main

/*
#cgo LDFLAGS: -L. -lstatic
#include "header.h"
*/
import "C"

func main() {
	C.function()
}
