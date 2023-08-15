/*
Package binary_tree
@Time : 2022/2/14 9:48
@Author : sunxy
@File : binarytree
@description:
*/
package main

import "fmt"

type Value struct {
	Val int
}

type TreeNode struct {
	Value
	Left  *TreeNode
	Right *TreeNode
}

type stack struct {
	val []*TreeNode
}

func newStack() *stack {
	s := new(stack)
	s.val = make([]*TreeNode, 0)
	return s
}

func (s *stack) put(node *TreeNode) {
	s.val = append(s.val, node)
}

func (s *stack) pop() *TreeNode {
	if len(s.val) == 0 {
		return nil
	}
	n := s.val[len(s.val)-1]
	s.val = s.val[:len(s.val)-1]
	return n
}

func (s *stack) isEmpty() bool {
	return len(s.val) == 0
}

func (s *stack) head() *TreeNode {
	return s.val[len(s.val)-1]
}

func (s *stack) len() int {
	return len(s.val)
}

// 前序遍历
func preOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Val)
	preOrderTraversal(root.Left)
	preOrderTraversal(root.Right)
}

// 前序遍历非递归
func preOrderTraversalNoRecursive(root *TreeNode) {
	s := newStack()
	if root == nil {
		return
	}
	node := root
	for node != nil || !s.isEmpty() {
		for node != nil {
			fmt.Printf("%d ", node.Val)
			s.put(node)
			node = node.Left
		}
		node = s.pop().Right
	}
}

// 中序遍历
func inOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	inOrderTraversal(root.Left)
	fmt.Printf("%d ", root.Val)
	inOrderTraversal(root.Right)
}

// 中序遍历非递归
func inOrderTraversalNoRecursive(root *TreeNode) {
	s := newStack()
	if root == nil {
		return
	}
	node := root
	for node != nil || !s.isEmpty() {
		for node != nil {
			s.put(node)
			node = node.Left
		}
		node = s.pop()
		fmt.Printf("%d ", node.Val)
		node = node.Right
	}
}

// 后序遍历
func postOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	postOrderTraversal(root.Left)
	postOrderTraversal(root.Right)
	fmt.Printf("%d ", root.Val)
}

// 后序遍历非递归
// 左 右 根      我们先 根 右 左  然后倒序
func postOrderTraversalNoRecursive(root *TreeNode) {
	s := newStack()
	res := newStack()
	if root == nil {
		return
	}
	node := root
	for node != nil || !s.isEmpty() {
		for node != nil {
			res.put(node)
			s.put(node)
			node = node.Right
		}
		node = s.pop().Left
	}
	for !res.isEmpty() {
		fmt.Printf("%d ", res.pop().Val)
	}
}

// 左 右 中
// https://zhuanlan.zhihu.com/p/80578741
func postOrderTraversalNoRecursive2(root *TreeNode) {
	s := newStack()
	if root == nil {
		return
	}
	var lastVisited *TreeNode
	node := root
	for node != nil || !s.isEmpty() {
		for node != nil {
			s.put(node)
			node = node.Left
		}
		// 代码至此，node.left 为空
		node = s.pop()
		if node.Right == nil || node.Right == lastVisited {
			fmt.Printf("%d ", node.Val)
			lastVisited = node
			node = nil
		} else {
			s.put(node)
			node = node.Right
		}
	}
}

func newTree() *TreeNode {
	root := &TreeNode{
		Value: Value{Val: 0},
		Left:  nil,
		Right: nil,
	}
	root.Left = &TreeNode{
		Value: Value{Val: 1},
		Left: &TreeNode{
			Value: Value{Val: 2},
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Value: Value{Val: 3},
			Left:  nil,
			Right: nil,
		},
	}
	root.Right = &TreeNode{
		Value: Value{Val: 8},
		Left:  nil,
		Right: nil,
	}
	return root
}

func main() {
	t := newTree()
	// preOrderTraversal(t)
	preOrderTraversalNoRecursive(t)
	fmt.Println()
	inOrderTraversalNoRecursive(t)
	fmt.Println()
	postOrderTraversal(t)
	fmt.Println()
	postOrderTraversalNoRecursive(t)
	fmt.Println()
	postOrderTraversalNoRecursive2(t)
}
