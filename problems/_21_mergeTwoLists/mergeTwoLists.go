/*
Package _21_mergeTwoLists
@Time : 2023/8/5 19:38
@Author : sunxy
@File : mergeTwoLists
@description:
*/
package _21_mergeTwoLists

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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	c1, c2 := list1, list2
	res := &ListNode{}
	resc := res
	for c1 != nil && c2 != nil {
		if c1.Val <= c2.Val {
			resc.Val = c1.Val
			c1 = c1.Next
		} else {
			resc.Val = c2.Val
			c2 = c2.Next
		}
		resc.Next = &ListNode{}
		resc = resc.Next
	}
	if c1 != nil {
		resc.Val = c1.Val
		resc.Next = c1.Next
	} else if c2 != nil {
		resc.Val = c2.Val
		resc.Next = c2.Next
	}
	return res
}
