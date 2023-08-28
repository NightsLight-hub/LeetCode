/*
Package _7_insert
@Time : 2023/8/28 17:38
@Author : sunxy
@File : insert_test.go
@description:
*/
package _7_insert

import (
	"testing"
)

func Test_insert(t *testing.T) {
	ret := insert([][]int{{1, 3}, {6, 9}}, []int{2, 5})
	// []int{{1,5},{6,9}}
	t.Logf("%#v", ret)
	ret = insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8})
	// []int{{1,2},{3,10},{12,16}}
	t.Logf("%#v", ret)

}
