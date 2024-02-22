/*
Package _300_LengthOfLis
@Time : 2023/12/4 10:09
@Author : sunxy
@File : lenghthOfLis
@description: https://leetcode.cn/problems/longest-increasing-subsequence/
*/
package _300_LengthOfLis

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	f := make([]int, len(nums))
	for i := range f {
		f[i] = -1
	}
	f[0] = 1
	answer := -1
	for i := range f {
		lengthOfLIS1(nums, i, f, &answer)
	}
	return answer
}

// f(i)是以nums[i]为结尾的最长增长子序列的长度
// f(i) = max(f(j)) + 1     0<=j<i && nums[j] < nums[i]

func lengthOfLIS1(nums []int, i int, f []int, answer *int) int {
	if f[i] != -1 {
		*answer = max(*answer, f[i])
		return f[i]
	}
	res := 1
	// f(j) = max(f(i)) + 1     0<=i<=n && nums[i] < nums[j]
	for j := 0; j < i; j++ {
		if nums[j] < nums[i] {
			res = max(res, lengthOfLIS1(nums, j, f, answer)+1)
		}
	}
	f[i] = res
	*answer = max(*answer, res)
	return f[i]
}

func lengthOfLIS2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}
	res := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j <= i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
				res = max(res, dp[i])
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
