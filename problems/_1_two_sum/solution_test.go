/*
@Time : 2021/7/6 17:50
@Author : sunxy
@File : solution_test.go
@description:
*/
package _1_two_sum

import (
	"testing"
)

func Test_twoSum(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6
	t.Logf("%+v", twoSum(nums, target))
}
