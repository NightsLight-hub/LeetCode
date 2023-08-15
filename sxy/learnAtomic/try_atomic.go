/*
@Time : 2021/7/19 14:57
@Author : sunxy
@File : try_atomic
@description:
*/
package learnAtomic

import "sync/atomic"

func add(num *int64) {
	atomic.AddInt64(num, 1)
}
func add2(num *int64) {
	*num = *num + 1
}

type TDInterface interface {
	Name() string
}

type TD struct {
}

func (_self TD) Name() string {
	return "1231"
}

func getTD() TDInterface {
	var td *TD = nil
	return td
}
