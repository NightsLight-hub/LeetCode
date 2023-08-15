/*
Package _50_Delete_Node_in_a_BST
@Time : 2022/6/2 11:42
@Author : sunxy
@File : solution
@description:
*/
package _450_Delete_Node_in_a_BST

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	return Delete(root, key)
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
