/*
Package sync
@Time : 2023/8/4 9:00
@Author : sunxy
@File : pool_test.go
@description:
*/
package sync

import "testing"

func Test_pack(t *testing.T) {
	a := pack(100, 200)
	t.Log(a)
	head, tail := unpack(a)
	t.Log(head, tail)
}
