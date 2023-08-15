/*
@Time : 2021/7/19 14:59
@Author : sunxy
@File : try_atomic_test.go
@description:
*/
package learnAtomic

import (
	"sync"
	"testing"
)

func Test_add(t *testing.T) {
	var num int64 = 1
	var w sync.WaitGroup
	w.Add(2)
	go func() {
		for i := 0; i < 1e6; i++ {
			add(&num)
		}
		w.Done()
	}()

	go func() {
		for i := 0; i < 1e6; i++ {
			add(&num)
		}
		w.Done()
	}()
	w.Wait()
	t.Log(num)
}

func Test_add2(t *testing.T) {
	var num int64 = 1
	var w sync.WaitGroup
	w.Add(2)
	go func() {
		for i := 0; i < 1e6; i++ {
			add2(&num)
		}
		w.Done()
	}()

	go func() {
		for i := 0; i < 1e6; i++ {
			add2(&num)
		}
		w.Done()
	}()
	w.Wait()
	t.Log(num)
}

func Test_add3(t *testing.T) {
	var td *TD = nil
	t.Logf("%v", td == nil)
	td2 := getTD()
	t.Logf("%v", td2 == nil)
}
