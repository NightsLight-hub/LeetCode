/*
Package _1448_goodNodes
@Time : 2023/8/25 16:59
@Author : sunxy
@File : goodNodes
@description: https://leetcode.cn/problems/count-good-nodes-in-binary-tree/
*/
package _1448_goodNodes

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

func goodNodes(root *TreeNode) int {
	var walk func(*TreeNode, int) int
	walk = func(node *TreeNode, max int) int {
		if node == nil {
			return 0
		}
		var res int
		if node.Val >= max {
			res = 1
			max = node.Val
		}
		return res + walk(node.Left, max) + walk(node.Right, max)
	}
	return walk(root, root.Val)
}
