/*
Package _186_maximumSum
@Time : 2023/6/27 9:44
@Author : sunxy
@File : maximumSum
@description: https://leetcode.cn/problems/maximum-subarray-sum-with-one-deletion/
*/
package _186_maximumSum

/*
*
转移方程
dp[i][0] = max(dp[i-1][0], 0) + arr[i]
dp[i][1] = max( dp[i-1][1] + arr[i], dp[i-1][0] )
*/
func maximumSum(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}

	dpi0 := arr[0]
	dpi1 := 0
	res := dpi0
	for i := 1; i < len(arr); i++ {
		dpi0, dpi1 = max(dpi0, 0)+arr[i], max(dpi1+arr[i], dpi0)
		res = max(max(dpi1, dpi0), res)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
