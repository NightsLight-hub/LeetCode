/*
Package _136_singleNumber
@Time : 2023/10/14 9:12
@Author : sunxy
@File : singleNumber
@description:
*/
package _136_singleNumber

func singleNumber(nums []int) int {
	a := nums[0]
	for i := 1; i < len(nums); i++ {
		a ^= nums[i]
	}
	return a
}
