/*
Package _817_linkedListCom
@Time : 2022/10/12 10:27
@Author : sunxy
@File : solution
@description:
*/
package _817_linkedListCom

type ListNode struct {
	Val  int
	Next *ListNode
}

const wrongNum = 10e5

func numComponents(head *ListNode, nums []int) int {
	ret := 0
	var node *ListNode
	var numsH = make([]int, len(nums))
	copy(numsH, nums)
	for len(numsH) != 0 {
		var prev int
		var lstNode *ListNode
		node, prev = getNode(head, numsH[0])
		if node == nil {
			panic("not possible")
		}
		numsH = numsH[1:]
		ret++
		for {
			lstNode = node
			node = node.Next
			if node == nil {
				break
			}
			index := -1
			for i, n := range numsH {
				if n == node.Val || n == prev {
					index = i
					break
				}
			}
			if index >= 0 {
				numsH = append(numsH[:index], numsH[index+1:]...)
			} else {
				break
			}
			prev = lstNode.Val
		}
	}
	return ret
}

func getNode(head *ListNode, v int) (*ListNode, int) {
	a := head
	var prev int = wrongNum
	for a != nil {
		if a.Val == v {
			return a, prev
		}
		prev = a.Val
		a = a.Next
	}
	return nil, prev
}
