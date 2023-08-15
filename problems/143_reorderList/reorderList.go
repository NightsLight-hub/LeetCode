/*
Package _43_reorderList
@Time : 2023/7/31 8:27
@Author : sunxy
@File : reorderList
@description:
*/
package _43_reorderList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	var arr = make([]*ListNode, 0)
	h := head
	for h != nil {
		arr = append(arr, h)
		h = h.Next
	}
	l := len(arr)
	if l <= 2 {
		return
	}
	i, j := 0, len(arr)-1
	for i < j {
		arr[i].Next = arr[j]
		i++
		if i == j {
			break
		}
		arr[j].Next = arr[i]
		j--
	}
	arr[i].Next = nil
}
