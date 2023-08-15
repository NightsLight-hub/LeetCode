/*
Package _1289_minFallingPathSum
@Time : 2023/8/10 9:52
@Author : sunxy
@File : minFallingPathSum_test.go
@description:
*/
package _1289_minFallingPathSum

import "testing"

func Test_minFallingPathSum(t *testing.T) {
	grid := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	println(minFallingPathSum(grid))
	grid = [][]int{{2, 1, 3}, {50, 51, 52}, {1, 101, 102}}
	println(minFallingPathSum(grid))
}
