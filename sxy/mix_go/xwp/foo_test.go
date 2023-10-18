/*
@Time : 2021/6/10 9:39
@Author : sunxy
@File : foo_test.go
@description:
*/
package xwp

import "testing"

func Test_fuck(t *testing.T) {
	fuck()
}

func Test_ff2(t *testing.T) {
	a := map[string]string{}
	b, ok := a["asdf"]
	print(b)
	print(ok)
}
