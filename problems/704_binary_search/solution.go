/*
@Time : 2021/7/3 14:00
@Author : sunxy
@File : solution
@description:
*/
package _04_binary_search

// 2 4 5
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	// 注意 数组取中间位置的方法是什么 left + (right - left)/2
	for left <= right {
		i := left + (right-left)/2
		if nums[i] == target {
			return i
		} else if nums[i] < target {
			left = i + 1
		} else {
			right = i - 1
		}
	}
	return -1
}
