/*
Package testCh
@Time : 2022/8/2 10:11
@Author : sunxy
@File : client2_test.go
@description:
*/
package testCh

import (
	"fmt"
	"testing"
	"time"
)

func Test_test2(t *testing.T) {
	cht2()
}

func Test_aa(t *testing.T) {
	ch := make(chan string, 0)
	ch <- "string"
	t.Log("asdf")
}

func Test_Tick(t *testing.T) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case <-ticker.C:
			fmt.Println("Current time: ", time.Now())
			time.Sleep(2 * time.Second)
		}
	}
}
