/*
Package _32_K_diff_Pairs_in_an_Array
@Time : 2022/6/17 17:40
@Author : sunxy
@File : solution
@description:
*/
package _32_K_diff_Pairs_in_an_Array

import "sort"

func findPairs(nums []int, k int) int {
	m := make(map[int]int)
	sort.Ints(nums)
	for i, n := range nums {
		if _, ok := m[n]; ok {
			continue
		}
		if i != len(nums)-1 {
			m[n] = searchN(nums[i+1:], n+k)
		}
	}
	c := 0
	for _, v := range m {
		c += v
	}
	return c
}

// Returns the number of occurrences of the target
func searchN(nums []int, target int) int {
	l, r := 0, len(nums)-1
	flag := -1
	for flag < 0 && l >= 0 && r >= 0 && l <= r {
		mid := (l + r + 1) / 2
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			flag = mid
		}
	}
	if flag < 0 {
		return 0
	}
	return 1
}
