/*
Package _513_findBottomLeftValue
@Time : 2022/6/22 16:44
@Author : sunxy
@File : solution
@description:
*/
package _513_findBottomLeftValue

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

// 前序遍历 （深度优先）
// func findBottomLeftValue(root *TreeNode) int {
//     // valArray := make([]int, 0)
//     // depthArray := make([]int, 0)
//     var preTravel func(*TreeNode, int)
//     max := -1
//     ret := 0
//     preTravel = func(node *TreeNode, depth int) {
//         if node.Left != nil {
//             preTravel(node.Left, depth+1)
//         }
//         // depthArray = append(depthArray, depth)
//         // valArray = append(valArray, node.Val)
//         if max < depth {
//             max = depth
//             ret = node.Val
//         }
//         if node.Right != nil {
//             preTravel(node.Right, depth+1)
//         }
//     }
//     preTravel(root, 0)
//     return ret
// }

// 广度优先
// 前序遍历 （深度优先）
func findBottomLeftValue(root *TreeNode) int {
	currentDepthNode := []*TreeNode{root}
	var node *TreeNode
	for len(currentDepthNode) != 0 {
		node = currentDepthNode[0]
		currentDepthNode = currentDepthNode[1:]
		if node.Right != nil {
			currentDepthNode = append(currentDepthNode, node.Right)
		}
		if node.Left != nil {
			currentDepthNode = append(currentDepthNode, node.Left)
		}
	}
	return node.Val
}
