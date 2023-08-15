// Package _189_Rotate_Array /*
// https://leetcode-cn.com/problems/rotate-array/
package _189_Rotate_Array

func rotate(nums []int, k int) {
	l := len(nums)
	nb := make([]int, l)
	copy(nb, nums)
	for i, n := range nb {
		index := (i + k) % l
		nums[index] = n
	}
}
