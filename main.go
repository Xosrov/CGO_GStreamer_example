package main

/*
#cgo pkg-config: --define-variable=prefix=C:/gstreamer/1.0/mingw_x86_64 gstreamer-1.0
#cgo CFLAGS: -I${SRCDIR}
#include "gstexample.h"
*/
import "C"
import "fmt"

func main() {
	result := C.start_pipeline()
	fmt.Println(result)
}
