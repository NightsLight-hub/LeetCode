/*
@Time : 2021/7/3 14:52
@Author : sunxy
@File : solution
@description:
*/
package _77_Squares_of_a_Sorted_Array

import "math"

func sortedSquares(nums []int) []int {
	negArrays, PosArrays := splitTowArrays(nums)
	negArrays = squares(negArrays, false)
	PosArrays = squares(PosArrays, true)
	return merge(negArrays, PosArrays)
}

func merge(arrays1, arrays2 []int) []int {
	var res []int
	a1Left, a2Left := 0, 0
	for a1Left != len(arrays1) || a2Left != len(arrays2) {
		var a1, a2 int
		if a1Left < len(arrays1) {
			a1 = arrays1[a1Left]
		} else {
			a1 = math.MaxInt32
		}
		if a2Left < len(arrays2) {
			a2 = arrays2[a2Left]
		} else {
			a2 = math.MaxInt32
		}

		if a1 >= a2 {
			res = append(res, a2)
			a2Left++
		} else {
			res = append(res, a1)
			a1Left++
		}
	}
	return res
}

func splitTowArrays(nums []int) ([]int, []int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] >= 0 {
			return nums[0:i], nums[i:]
		}
	}
	return nums, []int{}
}

func squares(nums []int, flag bool) []int {
	if flag {
		for i, n := range nums {
			nums[i] = n * n
		}
		return nums
	} else {
		var res []int = make([]int, len(nums))
		copy(res, nums)
		l := len(nums) - 1
		for i, n := range nums {
			res[l-i] = n * n
		}
		return res
	}

}

func min(a, b int) int {
	if a > b {
		return a
	}
	return b
}
