/*
Package _43_reorderList
@Time : 2023/7/31 9:03
@Author : sunxy
@File : reorderList_test.go
@description:
*/
package _43_reorderList

import "testing"

func Test_reorderList(t *testing.T) {
	h := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	reorderList(h)
	println("ok")
}
