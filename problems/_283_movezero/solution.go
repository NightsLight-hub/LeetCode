/*
@Time : 2021/7/5 17:52
@Author : sunxy
@File : solution
@description:
*/
package _283_movezero

func moveZeroes(nums []int) {
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[index] = nums[i]
			index++
		}
	}
	for ; index < len(nums); index++ {
		nums[index] = 0
	}
}
