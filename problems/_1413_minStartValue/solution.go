/*
Package _1413_minStartValue
@Time : 2022/8/9 9:08
@Author : sunxy
@File : solution
@description:
*/
package _1413_minStartValue

import "math"

func minStartValue(nums []int) int {
	var min, sum = math.MaxInt32, 0

	for _, a := range nums {
		sum += a
		if sum < min {
			min = sum
		}
	}
	if min <= 0 {
		return 1 - min
	}
	return 1
}
