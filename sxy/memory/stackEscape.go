/*
Package memory
@Time : 2023/6/8 17:10
@Author : sunxy
@File : stackEscape
@description:
*/
package main

import (
	"fmt"
	"unsafe"
)

type student struct {
	name string // 16byte
}

func stackSpace() {
	// 动态大小，发生逃逸
	const length = 10
	space1 := make([]int, length)
	for i := 0; i < len(space1); i++ {
		space1[i] = i
	}
	len1 := unsafe.Sizeof(space1)
	fmt.Println("space1 length:", len1)

	// 未发生逃逸
	space2 := make([]int, 100, 100)
	for i := 0; i < len(space2); i++ {
		space2[i] = i
	}
	len2 := unsafe.Sizeof(space2)
	fmt.Println("space2 length:", len2)

	// 当int32整型数组容量大于8192*2时（64KB，单个int32大小为4字节），发生逃逸
	space3 := make([]int32, 8192*2+1, 8192*2+1)
	for i := 0; i < len(space3); i++ {
		space3[i] = int32(i)
	}
	len3 := unsafe.Sizeof(space3)
	fmt.Println("space3 length:", len3)

	// 当int64整型数组容量大于8192时（64KB，单个int64大小为8字节），发生逃逸
	space4 := make([]int64, 8193, 8193)
	for i := 0; i < len(space4); i++ {
		space4[i] = int64(i)
	}
	len4 := unsafe.Sizeof(space4)
	fmt.Println("space4 length:", len4)

	// 当字符串指针数组容量大于4096时（64KB，单个字符串大小为16字节），发生逃逸
	space5 := make([]string, 4097, 4097)
	for i := 0; i < len(space5); i++ {
		s := "a"
		space5[i] = s
	}
	len5 := unsafe.Sizeof(space5)
	fmt.Println("space5 length:", len5)

	// 当字符串指针数组容量大于8192时（64KB，单个指针大小为8字节），发生逃逸，变量s也发生逃逸
	space6 := make([]*string, 8193, 8193)
	for i := 0; i < len(space6); i++ {
		s := "abc"
		space6[i] = &s
	}
	len6 := unsafe.Sizeof(space6)
	fmt.Println("space6 length:", len6)

	// 当自定义对象数组容量大于4096时（64KB，单个对象大小为16字节），发生逃逸
	space7 := make([]student, 4097, 4097)
	for i := 0; i < len(space7); i++ {
		space7[i].name = "abc"
	}
	len7 := unsafe.Sizeof(space7)
	fmt.Println("space7 length:", len7)
}
func main() {
	stackSpace()
}
