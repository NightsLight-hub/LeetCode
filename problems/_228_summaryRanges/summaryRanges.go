/*
Package _228_summaryRanges
@Time : 2023/8/26 9:14
@Author : sunxy
@File : summaryRanges
@description: https://leetcode.cn/problems/summary-ranges/
*/
package _228_summaryRanges

import "fmt"

// 双指针
func summaryRanges(nums []int) []string {
	i, j := 0, 1
	var res []string
	for ; j <= len(nums); j++ {
		if (j < len(nums) && nums[j] != nums[j-1]+1) || j == len(nums) {
			if j-i == 1 {
				res = append(res, fmt.Sprintf("%d", nums[i]))
			} else {
				res = append(res, fmt.Sprintf("%d->%d", nums[i], nums[j-1]))
			}
			i = j
		}
	}
	// if j-i == 1 {
	// 	res = append(res, fmt.Sprintf("%d", i))
	// } else {
	// 	res = append(res, fmt.Sprintf("%d->%d", i, j-1))
	// }
	return res
}
