/*
Package _2530_maxKelements
@Time : 2023/10/18 9:37
@Author : sunxy
@File : maxkelements
@description:  https://leetcode.cn/problems/maximal-score-after-applying-k-operations/

给你一个下标从 0 开始的整数数组 nums 和一个整数 k 。你的 起始分数 为 0 。

在一步 操作 中：

选出一个满足 0 <= i < nums.length 的下标 i ，
将你的 分数 增加 nums[i] ，并且
将 nums[i] 替换为 ceil(nums[i] / 3) 。
返回在 恰好 执行 k 次操作后，你可能获得的最大分数。

向上取整函数 ceil(val) 的结果是大于或等于 val 的最小整数。
*/
package _2530_maxKelements

import (
	"github.com/emirpasic/gods/trees/binaryheap"
	"math"
)

// IntComparator provides a basic comparison on int
func MaxHeapIntComparator(a, b interface{}) int {
	aAsserted := a.(int)
	bAsserted := b.(int)
	switch {
	case aAsserted > bAsserted:
		return -1
	case aAsserted < bAsserted:
		return 1
	default:
		return 0
	}
}

func maxKelements(nums []int, k int) int64 {
	// max-heap
	h := binaryheap.NewWith(MaxHeapIntComparator)
	for _, v := range nums {
		h.Push(v)
	}
	var res int64 = 0
	for i := 0; i < k; i++ {
		c, _ := h.Pop()
		res += int64(c.(int))
		rc := int(math.Ceil(float64(c.(int)) / 3))
		h.Push(rc)
	}
	return res
}
