/*
Package _352_EqualRowAndColumnPairs
@Time : 2023/6/6 20:25
@Author : sunxy
@File : Solution_test.go
@description:
*/
package _352_EqualRowAndColumnPairs

import "testing"

func Test_equalPairs(t *testing.T) {
	grid := make([][]int, 0)
	grid = append(grid, []int{3, 2, 1})
	grid = append(grid, []int{1, 7, 6})
	grid = append(grid, []int{2, 7, 7})
	t.Logf("res: %v", equalPairs(grid))
	grid = make([][]int, 0)
	grid = append(grid, []int{3, 1, 2, 2})
	grid = append(grid, []int{1, 4, 4, 5})
	grid = append(grid, []int{2, 4, 2, 2})
	grid = append(grid, []int{2, 4, 2, 2})
	t.Logf("res: %v", equalPairs(grid))
}
