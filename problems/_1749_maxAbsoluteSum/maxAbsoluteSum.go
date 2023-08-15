/*
Package _1749_maxAbsoluteSum
@Time : 2023/8/8 11:18
@Author : sunxy
@File : maxAbsoluteSum
@description:
*/
package _1749_maxAbsoluteSum

func maxAbsoluteSum(nums []int) int {
	positiveMax, negativeMin := 0, 0
	positiveSum, negativeSum := 0, 0
	for _, num := range nums {
		positiveSum += num
		positiveMax = max(positiveMax, positiveSum)
		positiveSum = max(0, positiveSum)
		negativeSum += num
		negativeMin = min(negativeMin, negativeSum)
		negativeSum = min(0, negativeSum)
	}
	return max(positiveMax, -negativeMin)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
