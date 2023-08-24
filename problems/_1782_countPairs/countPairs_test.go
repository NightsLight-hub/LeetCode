/*
Package _1782_countPairs
@Time : 2023/8/23 10:54
@Author : sunxy
@File : countPairs_test.go
@description:
*/
package _1782_countPairs

import (
	"testing"
)

func Test_countPairs(t *testing.T) {
	// t.Log(countPairs(4, [][]int{{1, 2}, {2, 4}, {1, 3}, {2, 3}, {2, 1}}, []int{2, 3}))
	t.Log(countPairs(5, [][]int{{1, 5}, {1, 5}, {3, 4}, {2, 5}, {1, 3}, {5, 1}, {2, 3}, {2, 5}}, []int{1, 2, 3, 4, 5}))
}
