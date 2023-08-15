/*
Package _24_swapPairs
@Time : 2023/8/6 8:49
@Author : sunxy
@File : swapPairs_test.go
@description:
*/
package _24_swapPairs

import (
	"testing"
)

func generateListNode(a []int) *ListNode {
	l := &ListNode{}
	h := l
	for _, v := range a {
		l.Next = &ListNode{Val: v}
		l = l.Next
	}
	return h.Next
}

func Test_swapPairs3(t *testing.T) {
	l := generateListNode([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	r := swapPairs(l)
	for r != nil {
		t.Log(r.Val)
		r = r.Next
	}
}

func Test_swapPairs2(t *testing.T) {
	l := generateListNode([]int{0})
	r := swapPairs(l)
	for r != nil {
		t.Log(r.Val)
		r = r.Next
	}
}

func Test_swapPairs(t *testing.T) {
	l := generateListNode([]int{1, 2, 3, 4})
	r := swapPairs(l)
	for r != nil {
		t.Log(r.Val)
		r = r.Next
	}
}
