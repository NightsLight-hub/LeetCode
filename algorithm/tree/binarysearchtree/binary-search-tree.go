/*
Package binarysearchtree
@Time : 2022/6/3 10:21
@Author : sunxy
@File : binary-search-tree
@description:
*/
package binarysearchtree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// search key
func Get(node *TreeNode, key int) *TreeNode {
	if node == nil {
		return nil
	}
	if node.Val < key {
		return Get(node.Right, key)
	}
	if node.Val > key {
		return Get(node.Left, key)
	}
	return node
}

// put node
func Put(node *TreeNode, val int) *TreeNode {
	if node == nil {
		node = &TreeNode{
			Val:   val,
			Left:  nil,
			Right: nil,
		}
		return node
	}
	if node.Val < val {
		node.Right = Put(node.Right, val)
	} else if node.Val > val {
		node.Left = Put(node.Left, val)
	}
	return node
}

func PreTravel(root *TreeNode) {
	if root == nil {
		return
	}
	PreTravel(root.Left)
	fmt.Printf("%d,", root.Val)
	PreTravel(root.Right)
}

func Min(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	for root.Left != nil {
		root = root.Left
	}
	return root
}

// floor 找到小于key的最大值
func Floor(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	// root 小于key，root 可能是答案
	if root.Val < key {
		// 如果root 不是答案，肯定在右子树上
		node := Floor(root.Right, key)
		// 右子树没有答案
		if node == nil {
			// 当前这个节点就是合适的答案
			return root
		}
		// 右子树上的答案
		return node
	}
	// root 大于可以， 只能在左子树上
	if root.Val > key {
		return Floor(root.Left, key)
	}
	return root
}

// ceil
func Ceil(root *TreeNode, val int) {

}

func DeleteMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left == nil {
		root = root.Right
	} else {
		root.Left = DeleteMin(root.Left)
	}
	return root
}

// delete
func Delete(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return nil
	}
	if node.Val < val {
		node.Right = Delete(node.Right, val)
	} else if node.Val > val {
		node.Left = Delete(node.Left, val)
	} else {
		if node.Left == nil {
			node = node.Right
		} else if node.Right == nil {
			node = node.Left
		} else {
			t := Min(node.Right)
			node.Right = DeleteMin(node.Right)
			node.Val = t.Val
		}
	}
	return node
}
