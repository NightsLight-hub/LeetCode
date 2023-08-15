/*
Package stack
@Time : 2022/5/28 10:18
@Author : sunxy
@File : stack_test.go
@description:
*/
package stack

import (
	"strconv"
	"testing"
)

func se(e Element) string {
	return strconv.Itoa(e.(int))
}

func TestStack_1(t *testing.T) {
	s := NewStack()

	s.Push(1).Push(2).Push(3)
	if ret := s.Show(se); ret != "1 2 3" {
		t.Errorf("wrong")
	}
	s.Pop()
	if ret := s.Show(se); ret != "1 2" {
		t.Errorf("wrong")
	}
	s.Push(4).Push(5)
	if ret := s.Show(se); ret != "1 2 4 5" {
		t.Errorf("wrong")
	}

}
