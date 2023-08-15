/*
@Time : 2021/7/22 15:31
@Author : sunxy
@File : solution_test.go
@description:
*/
package _138_Copy_List_with_Random_Pointer

import (
	"fmt"
	"testing"
)

func Test_copyRandomList(t *testing.T) {
	n0 := new(Node)
	n0.Val = 1

	n1 := new(Node)
	n1.Val = 2
	n1.Random = n0

	n0.Next = n1
	n0.Random = n1
	n1.Random = n1
	output(n0)

	output(copyRandomList(n0))
}

func output(n *Node) {
	for n != nil {
		fmt.Printf("%d  ", n.Val)
		n = n.Next
	}
	fmt.Println()
}
