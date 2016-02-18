package main

/*
#cgo LDFLAGS: -L. -ldynamic
#include "header.h"
*/
import "C"

func main() {
	C.function()
}
