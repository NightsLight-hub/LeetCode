/*
Package try_cgo_callstatic
@Time : 2024/4/1 15:57
@Author : sunxy
@File : main
@description:
*/
package main

// #cgo CXXFLAGS: -std=c++17
// #cgo CFLAGS: -I/mnt/e/CODE/LeetCode/cpp/gstreamer-cmake/include
// #cgo LDFLAGS: -L/mnt/e/CODE/LeetCode/cpp/gstreamer-cmake/build -ldemolib
//
// #include "demolib.h"
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	C.hello()
	var ptr unsafe.Pointer
	ptr = C.getInstance()
	fmt.Println(C.producer(ptr))
	fmt.Println(C.pop(ptr))
}
