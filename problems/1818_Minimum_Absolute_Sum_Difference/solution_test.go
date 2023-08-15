/*
@Time : 2021/7/14 16:57
@Author : sunxy
@File : solution_test.go
@description:
*/
package _818_Minimum_Absolute_Sum_Difference

import "testing"

func Test_minAbsoluteSumDiff(t *testing.T) {
	nums1 := []int{2, 1, 7}
	nums2 := []int{3, 5, 6}
	t.Log(minAbsoluteSumDiff(nums1, nums2))
}
