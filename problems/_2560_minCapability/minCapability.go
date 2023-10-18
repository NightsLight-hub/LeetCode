/*
Package _2560_minCapability
@Time : 2023/9/19 9:45
@Author : sunxy
@File : minCapability
@description: https://leetcode.cn/problems/house-robber-iv/description/
*/
package _2560_minCapability

/**
1. 如果想打劫最大数目的房屋，要能早抢就早抢。因为实际上，只能抢 n/2+n%2个，一旦能抢不抢，一定少
2. 二分法。既然是在n个中，寻找不相邻k个数字的最小值， 最小值一定在nums中。
可以用f(n) 表示 最小值为n时候， 能抢的最多房屋数。 显然具备单调性，f(n)越大，说明n越大，可选的房屋越多。反之亦然
则如果f(n)>k,说明 还可以选择更小的值n2，因此n为右值
*/

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minCapability(nums []int, k int) int {
	lower, upper := nums[0], nums[0]
	for _, x := range nums {
		lower = min(lower, x)
		upper = max(upper, x)
	}

	for lower < upper {
		middle := (lower + upper) / 2
		count, visited := 0, false
		for _, x := range nums {
			if x <= middle && !visited {
				count, visited = count+1, true
			} else {
				visited = false
			}
		}
		if count >= k {
			upper = middle - 1
		} else {
			lower = middle + 1
		}
	}
	/**
	有没有可能，二分出来的答案，不在 nums 中呢？
	这是不可能的，答案一定在 nums 中。证明如下：
	设答案为 ans，也就是当最大金额为 ans 时，可以偷至少 k 间房子。
	如果 ans 不在 nums 中，那么当最大金额为 ans-1 时，也可以偷至少 k 间房子。
	这与二分算法相矛盾：根据 视频 中讲的红蓝染色法，循环结束时，ans 和 ans-1
	的颜色必然是不同的，即 ans 可以满足题目要求，而 ans-1 不满足题目要求。所
	以，二分出来的答案，一定在 nums 中。
	*/
	return lower
}
