/*
Package _24_swapPairs
@Time : 2023/8/6 8:12
@Author : sunxy
@File : swapPairs
@description:
*/
package _24_swapPairs

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

// 1->2->3->4->5->6
func swapPairs(head *ListNode) *ListNode {
	var swapf = func(a *ListNode) bool {
		if a.Next == nil || a.Next.Next == nil {
			return false
		}
		b, c, d := a.Next, a.Next.Next, a.Next.Next.Next
		a.Next = c
		c.Next = b
		b.Next = d
		return true
	}
	if head == nil || head.Next == nil {
		return head
	}
	res := head.Next
	temp := head.Next.Next
	head.Next.Next = head
	head.Next = temp
	cur := res.Next
	for swapf(cur) {
		cur = cur.Next.Next
	}
	return res
}
