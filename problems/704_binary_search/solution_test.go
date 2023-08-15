/*
@Time : 2021/7/3 14:05
@Author : sunxy
@File : solution_test.go
@description:
*/
package _04_binary_search

import "testing"

func Test_search(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	t.Log(search(nums, 2))
}
