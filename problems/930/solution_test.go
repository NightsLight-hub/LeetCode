/*
@Time : 2021/7/8 15:39
@Author : sunxy
@File : solution_test.go
@description:
*/
package _30

import "testing"

func Test_numSubarraysWithSum(t *testing.T) {
	nums := []int{1, 0, 1, 0, 1}
	goal := 2
	t.Log(numSubarraysWithSum(nums, goal))
}
