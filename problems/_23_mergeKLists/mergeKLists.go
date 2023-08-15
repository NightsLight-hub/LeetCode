/*
Package _23_mergeKLists
@Time : 2023/8/12 7:51
@Author : sunxy
@File : mergeKLists
@description:
*/
package _23_mergeKLists

import "math"

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

func mergeKLists(lists []*ListNode) *ListNode {
	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	cursor := head
	for {
		minIndex, minValue := -1, math.MaxInt32
		for i := range lists {
			if lists[i] != nil && lists[i].Val < minValue {
				minIndex = i
				minValue = lists[minIndex].Val
			}
		}
		if minIndex == -1 {
			break
		}
		cursor.Next = &ListNode{
			Val:  minValue,
			Next: nil,
		}
		cursor = cursor.Next
		lists[minIndex] = lists[minIndex].Next
	}
	return head.Next
}
