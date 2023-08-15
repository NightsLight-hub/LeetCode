/*
@Time : 2021/7/22 14:44
@Author : sunxy
@File : solution
@description:
*/
package _138_Copy_List_with_Random_Pointer

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

var nodeIndexMap = make(map[*Node]int)
var newNodeAddArray []*Node

func copyRandomList(head *Node) *Node {
	intiNodeIndexMap(head)
	newList := copyList(head)
	copyRandom(head, newList)
	return newList
}

func copyRandom(old, new *Node) {
	for old != nil {
		if old.Random == nil {
			new.Random = nil
		} else {
			new.Random = newNodeAddArray[nodeIndexMap[old.Random]]
		}
		new = new.Next
		old = old.Next
	}
}

func copyList(head *Node) *Node {
	newHead := new(Node)
	newHead.Val = head.Val
	cursor := newHead
	newNodeAddArray = append(newNodeAddArray, cursor)
	for head.Next != nil {
		cursor.Next = new(Node)
		head = head.Next
		cursor = cursor.Next
		cursor.Val = head.Val
		newNodeAddArray = append(newNodeAddArray, cursor)
	}
	return newHead
}

func intiNodeIndexMap(head *Node) {
	n := head
	i := 0
	for n != nil {
		nodeIndexMap[n] = i
		n = n.Next
		i++
	}
}
