/*
@Time : 2021/7/11 10:55
@Author : sunxy
@File : solution
@description:
*/
package _808_Array_Concatenation

func getConcatenation(nums []int) []int {
	l := len(nums)
	target := make([]int, len(nums)*2)
	copy(target, nums)
	copy(target[l:], nums)
	return target
}
