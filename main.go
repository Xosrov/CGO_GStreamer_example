package main

/*
#cgo pkg-config: gstreamer-1.0
#cgo CFLAGS: -I${SRCDIR}
#include "gstexample.h"
*/
import "C"
import "fmt"

func main() {
	result := C.start_pipeline()
	fmt.Println(result)
}
