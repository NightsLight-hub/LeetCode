/*
Package _2236_checkTree
@Time : 2023/8/20 8:09
@Author : sunxy
@File : checkTree
@description:
*/
package _2236_checkTree

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

func checkTree(root *TreeNode) bool {
	if root.Left != nil && root.Right != nil {
		return root.Left.Val+root.Right.Val == root.Val
	}
	return false

}
