/*
Package __addtownumbers
@Time : 2023/6/2 18:05
@Author : sunxy
@File : main
@description:
*/
package __addtownumbers

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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	res := head
	node1, node2 := l1, l2
	j := 0
	for {
		v1, v2 := 0, 0
		if node1 != nil {
			v1 = node1.Val
			node1 = node1.Next
		}
		if node2 != nil {
			v2 = node2.Val
			node2 = node2.Next
		}
		res.Val = (v1 + v2 + j) % 10
		res.Next = &ListNode{
			Val:  0,
			Next: nil,
		}
		j = (v1 + v2 + j) / 10
		if node1 == nil && node2 == nil {
			if j == 0 {
				res.Next = nil
			} else {
				res.Next = &ListNode{
					Val:  j,
					Next: nil,
				}
			}
			return head
		}
		res = res.Next
	}
}
