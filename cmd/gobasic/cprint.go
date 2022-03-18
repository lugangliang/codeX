package main

/*
#include <stdio.h>
*/

import (
	"C"
	"unsafe"
)

func main() {
	cstr := C.CString("Hello", "world")
	C.puts(cstr)
	C.free(unsafe.Pointer(cstr))

	return
}
