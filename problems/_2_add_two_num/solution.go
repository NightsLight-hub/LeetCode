/*
@Time : 2021/7/12 14:56
@Author : sunxy
@File : solution
@description:
*/
package _2_add_two_num

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
	weight := 1
	num1 := l1.Val
	n1 := l1.Next
	for n1 != nil {
		weight *= 10
		num1 += n1.Val * weight
		n1 = n1.Next
	}
	weight = 1
	num2 := l2.Val
	n2 := l2.Next
	for n2 != nil {
		weight *= 10
		num2 += n2.Val * weight
		n2 = n2.Next
	}
	ans := num1 + num2

	ansL := new(ListNode)
	ansL.Val = ans % 10
	root := ansL
	ans = ans / 10
	for ans != 0 {
		ansL.Next = new(ListNode)
		ansL.Next.Val = ans % 10
		ansL = ansL.Next
		ans = ans / 10
	}
	return root
}
