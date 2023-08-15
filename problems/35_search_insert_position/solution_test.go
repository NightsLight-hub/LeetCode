/*
@Time : 2021/7/3 14:37
@Author : sunxy
@File : solution_test.go
@description:
*/
package _5_search_insert_position

import (
	"testing"
)

func Test_searchInsert(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 7
	t.Log(searchInsert(nums, target))
}
