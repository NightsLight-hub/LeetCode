/*
Package _1022_Sum_of_Root_To_Leaf_Binary_Numbers
@Time : 2022/5/30 20:41
@Author : sunxy
@File : solution
@description:
*/
package _1022_Sum_of_Root_To_Leaf_Binary_Numbers

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

func sumRootToLeaf(root *TreeNode) int {
	sum := 0
	travel(root, root.Val, &sum)
	return sum
}

func travel(node *TreeNode, base int, pathSums *int) {
	if node.Left == nil && node.Right == nil {
		*pathSums += base
		return
	}
	if node.Left != nil {
		n := node.Left
		travel(n, (base<<1)+n.Val, pathSums)
	}
	if node.Right != nil {
		n := node.Right
		travel(n, (base<<1)+n.Val, pathSums)
	}
}
