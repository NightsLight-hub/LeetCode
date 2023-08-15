/*
@Time : 2021/7/3 14:31
@Author : sunxy
@File : solution
@description:
https://leetcode-cn.com/problems/search-insert-position/
*/
package _5_search_insert_position

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		i := left + (right-left)/2
		if nums[i] == target {
			return i
		} else if nums[i] > target {
			right = i - 1
		} else {
			left = i + 1
		}
	}
	return right + 1
}
