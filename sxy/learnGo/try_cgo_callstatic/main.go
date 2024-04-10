/*
Package try_cgo_callstatic
@Time : 2024/4/1 15:57
@Author : sunxy
@File : main
@description:
*/
package main

import "C"

// #cgo CXXFLAGS: -std=c++17
// #cgo CPPFLAGS: -I/mnt/e/CODE/LeetCode/cpp/gstreamer-cmake/include
// #cgo LDFLAGS: -L/mnt/e/CODE/LeetCode/cpp/gstreamer-cmake/build -ldemolib
//
// #include "demolib.h"
// #include <stdlib.h>
// #include <stdint.h>
import "C"
import (
	"fmt"
	"time"
	"unsafe"
)

func main() {
	C.hello()
	var ptr unsafe.Pointer
	ptr = C.getQueueInstance()
	fmt.Printf("queue size : %v\n", C.queueSize(ptr))
	C.producer(ptr)
	fmt.Printf("queue size : %v\n", C.queueSize(ptr))
	var res = uint32(C.pop(ptr))
	fmt.Printf("pop one, %d \n", res)
	fmt.Printf("queue size : %v\n", C.queueSize(ptr))
	ss := "hello"
	for i := 0; i < 20; i++ {
		ss += ss
	}
	for i := 0; i < 100000; i++ {
		var str = C.CBytes([]byte(ss))
		C.outputCharArray(str, C.int(len([]byte(ss))))
		C.free(unsafe.Pointer(str))
		time.Sleep(time.Millisecond * 10)
	}
}
