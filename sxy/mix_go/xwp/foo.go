/*
@Time : 2021/6/10 9:38
@Author : sunxy
@File : foo
@description:
*/
package xwp

import (
	"fmt"
	"sync/atomic"
)
import "github.com/mix-go/xwp"

type Foo struct {
}

var count int64

func (t *Foo) Do(data interface{}) {
	i := data.(int)

	atomic.AddInt64(&count, 1)
	fmt.Println(i)
}

func try() {
	jobQueue := make(chan interface{}, 200)
	p := &xwp.WorkerPool{
		JobQueue:       jobQueue,
		MaxWorkers:     1000,
		InitWorkers:    100,
		MaxIdleWorkers: 100,
		RunI:           &Foo{},
	}

	go func() {
		// 投放任务
		for i := 0; i < 10000; i++ {
			jobQueue <- i
		}

		// 投放完停止调度
		p.Stop()
	}()

	p.Run() // 阻塞等待
	fmt.Println(count)
}

func tryPanic() {
	fmt.Println("start ...")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	panic("error")

}
