/*
Package _16_threeSumClosest
@Time : 2023/7/10 11:03
@Author : sunxy
@File : threeSumClosest
@description: https://leetcode.cn/problems/3sum-closest/
*/
package _16_threeSumClosest

import "sort"

func threeSumClosest(nums []int, target int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	res := 0
	minx := 1 << 30
	for a := 0; a < len(nums)-2; a++ {
		b, c := a+1, len(nums)-1
		for {
			sum := nums[a] + nums[b] + nums[c]
			if sum == target {
				return target
			}
			dist := distance(sum, target)
			if dist < minx {
				minx, res = dist, sum
			}
			if b == c-1 {
				break
			}
			if sum < target {
				b++
			} else {
				c--
			}
		}
	}
	return res
}

func distance(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
