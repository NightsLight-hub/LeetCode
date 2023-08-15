/*
Package testslice
@Time : 2023/6/9 9:15
@Author : sunxy
@File : try1.go
@description:
*/
package testslice

import (
	"fmt"
)

var a = make([]int, 1000000)

func t1() {
	l := 100
	b := make([]int, l)
	for i := 0; i < l; i++ {
		b[i] = i
	}
	c := b[:4]

	fmt.Printf("%#v\n", c)
	fmt.Printf("%d, %d\n", len(c), cap(c))
}

type ss struct {
}

func (s *ss) t() {
	fmt.Printf("%p", s)
}

func t2() {
	sinstance := new(ss)
	func(f func()) {
		f()
	}(sinstance.t)
}
