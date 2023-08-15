/*
Package testCh
@Time : 2022/8/2 10:10
@Author : sunxy
@File : client2
@description:
*/
package testCh

import "time"

func cht2() {
	var ch = make(chan string, 0)
	go func() {
		time.Sleep(time.Second * 10)
		ch <- ""
	}()
	select {
	case <-ch:
		println("1111")
	}
}
