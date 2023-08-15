/*
@Time : 2021/7/26 16:16
@Author : sunxy
@File : client
@description:
*/
package testCh

import (
	"fmt"
	"go/types"
	"sync"
	"time"
)

/**
golang中，一个协程panic，会影响所有协程
协程的panic只能被自己recover
*/

var ch = make(chan string, 2)
var ch2 = types.NewChan(types.SendRecv, types.Typ[types.String])

func server() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("capture panic %v\n", p)
		}
	}()
	time.Sleep(time.Second * 2)
	panic("test panic")
}
func client(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second * 4)
	fmt.Println("heihei")

}

func test() {
	var wg = sync.WaitGroup{}
	wg.Add(1)
	go server()
	go client(&wg)
	wg.Wait()
}
