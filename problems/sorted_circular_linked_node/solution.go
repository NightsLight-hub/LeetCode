/*
Package sorted_circular_linked_node
@Time : 2022/6/18 9:26
@Author : sunxy
@File : solution
@description:https://leetcode.cn/problems/4ueAj6/
*/
package sorted_circular_linked_node

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 * }
 */

type Node struct {
	Val  int
	Next *Node
}

func insert(aNode *Node, x int) *Node {
	if aNode == nil {
		head := &Node{
			Val:  x,
			Next: nil,
		}
		head.Next = head
		return head
	}
	origin := aNode
	val := aNode.Val
	var head, tail *Node
	findFlag := false
	for aNode.Next != origin && aNode.Next != nil {
		if aNode.Next.Val <= x {
			aNode = aNode.Next
			findFlag = true
			break
		}
		if aNode.Next.Val < val {
			tail = aNode
			head = aNode.Next
		}
		aNode = aNode.Next
	}
	if findFlag {
		aNode.Next = &Node{
			Val:  x,
			Next: aNode.Next,
		}
	} else if tail != nil {
		tail.Next = &Node{
			Val:  x,
			Next: head,
		}
	} else {
		aNode.Next = &Node{
			Val:  x,
			Next: origin,
		}
	}
	return origin
}
