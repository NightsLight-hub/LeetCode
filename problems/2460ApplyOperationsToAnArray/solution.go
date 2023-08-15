/*
Package _460ApplyOperationsToAnArray
@Time : 2023/6/5 10:54
@Author : sunxy
@File : solution
@description:
*/
package _460ApplyOperationsToAnArray

func applyOperations(nums []int) []int {
	op(nums, 0)
	res := make([]int, len(nums))
	i := 0
	for _, n := range nums {
		if n != 0 {
			res[i] = n
			i++
		}
	}
	return res
}

func op(nums []int, index int) {
	if index == len(nums)-1 {
		return
	}
	if nums[index] == nums[index+1] {
		nums[index] = 2 * nums[index]
		nums[index+1] = 0
	}
	index++
	op(nums, index)
}
