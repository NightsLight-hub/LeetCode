/*
Package binarysearchtree
@Time : 2022/6/3 10:29
@Author : sunxy
@File : binary-search-tree_test.go
@description:
*/
package binarysearchtree

import (
	"testing"
)

func buildTree() *TreeNode {
	root := &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}
	Put(root, 2)
	Put(root, 4)
	Put(root, 1)
	Put(root, -5)
	Put(root, -3)
	Put(root, -6)
	Put(root, -10)
	Put(root, 9)
	Put(root, 7)
	Put(root, 5)
	return root
}

func TestPutAndGet(t *testing.T) {
	root := buildTree()
	PreTravel(root)
	println()

	n := Get(root, 9)
	print(n.Val)
}

func TestDeleteMin(t *testing.T) {
	tree := buildTree()
	PreTravel(tree)
	println()
	tree = DeleteMin(tree)
	PreTravel(tree)
	println()

	tree = DeleteMin(tree)
	PreTravel(tree)
	println()

	tree = DeleteMin(tree)
	PreTravel(tree)
	println()

	tree = DeleteMin(tree)
	PreTravel(tree)
	println()
	tree = DeleteMin(tree)
	PreTravel(tree)
	println()
	tree = DeleteMin(tree)
	PreTravel(tree)
	println()
}

func TestDelete(t *testing.T) {
	tree := buildTree()
	PreTravel(tree)
	println()
	tree = Delete(tree, -6)
	PreTravel(tree)
	println()

	tree = Delete(tree, 1)
	PreTravel(tree)
	println()

	tree = Delete(tree, 5)
	PreTravel(tree)
	println()

	tree = Put(tree, 3)
	PreTravel(tree)
	println()

	tree = Delete(tree, 3)
	PreTravel(tree)
	println()

	tree = Delete(tree, 0)
	PreTravel(tree)
	println()
}
