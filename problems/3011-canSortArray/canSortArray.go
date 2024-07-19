/*
Package _011_canSortArray
@Time : 2024/7/13 16:27
@Author : sunxy
@File : canSortArray
@description:
https://leetcode.cn/problems/find-if-array-can-be-sorted/?envType=daily-question&envId=2024-07-13
给你一个下标从 0 开始且全是 正 整数的数组 nums 。

一次 操作 中，如果两个 相邻 元素在二进制下数位为 1 的数目 相同 ，那么你可以将这两个元素交换。你可以执行这个操作 任意次 （也可以 0 次）。

如果你可以使数组变有序，请你返回 true ，否则返回 false 。

思路 分成bit位相同的组， 下一组的最小值要比上组最大值大，能排序，否则 不能
*/
package _011_canSortArray

import "math/bits"

func canSortArray(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	currentMax, currentMin := nums[0], nums[0]
	lastMax := -1
	lastBit := calculateBit(nums[0])
	for i := 1; i < len(nums); i++ {
		nBit := calculateBit(nums[i])
		if nBit == lastBit {
			// 比特相同，算一组的，更新下最大值
			if currentMax < nums[i] {
				currentMax = nums[i]
			}
			if currentMin > nums[i] {
				currentMin = nums[i]
			}
			continue
		}
		// 比特不同，本组就结束了，如果本组的最小值比上一组的最大值小，说明无法排序。
		if currentMin < lastMax {
			return false
		}
		// 前面的能排序，继续下一组的计算
		lastBit = nBit
		lastMax = currentMax
		currentMax, currentMin = nums[i], nums[i]
	}
	// 到达长度，最后一组的最小值比上一组的最大值小，说明无法排序。
	if currentMin < lastMax {
		return false
	}
	return true
}

func calculateBit(num int) int {
	// count := 0
	// for num > 0 {
	// 	if num&1 == 1 {
	// 		count++
	// 	}
	// 	num >>= 1
	// }
	// return count
	return bits.OnesCount(uint(num))
}
