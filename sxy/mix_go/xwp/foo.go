/*
@Time : 2021/6/10 9:38
@Author : sunxy
@File : foo
@description:
*/
package xwp

type Foo struct {
}

var count int64

func fuck() {
	var t func()
	t = func() {
		i := 0
		for {
			i++
		}
	}
	go t()
	go t()
	go t()
	go t()
	go t()
	go t()
	select {}
}
