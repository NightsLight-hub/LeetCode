/*
Package _980
@Time : 2023/8/4 10:09
@Author : sunxy
@File : uniquePathsIII_test.go
@description:
*/
package _980

import "testing"

func Test_uniquePathsIII(t *testing.T) {
	grid := [][]int{{1, 0}, {0, 2}}
	t.Log(uniquePathsIII(grid))
	grid = [][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 2, -1}}
	t.Log(uniquePathsIII(grid))
}
