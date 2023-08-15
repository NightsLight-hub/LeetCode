/*
@Time : 2021/7/20 16:16
@Author : sunxy
@File : solution
@description:
*/
package _877_Minimize_Maximum_Pair_Sum_in_Array

import "sort"

func minPairSum(nums []int) int {
	sort.Ints(nums)
	l := len(nums)
	m := 0
	for i := 0; i < l/2; i++ {
		m = max(m, nums[i]+nums[l-1-i])
	}
	return m
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
