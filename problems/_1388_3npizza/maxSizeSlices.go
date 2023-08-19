/*
Package _1388_3npizza
@Time : 2023/8/18 9:56
@Author : sunxy
@File : maxSizeSlices
@description: https://leetcode.cn/problems/pizza-with-3n-slices/description/
*/
package _1388_3npizza

import "math"

/*
*
dp[i][j] = max(s[i]+dp[i-2][j-1], dp[i-1][j])
j <= (i+1+1)/2     i 从0开始，0 就表示从下表为 0 到下标为0的slice中选j块
*/
func maxSizeSlices(slices []int) int {
	var dp = make([][]int, len(slices))
	res := math.MinInt32
	for i := range dp {
		dp[i] = make([]int, len(slices))
	}
	dp[0][0] = 0
	dp[0][1] = slices[0]
	dp[1][0] = 0
	dp[1][1] = max(slices[0], slices[1])
	dp[1][2] = math.MinInt32
	for i := 2; i < len(slices)-1; i++ {
		for j := 1; j <= (i+2)/2; j++ {
			dp[i][j] = max(slices[i]+dp[i-2][j-1], dp[i-1][j])
		}
	}
	res = dp[len(slices)-2][len(slices)/3]
	for i := range dp {
		dp[i] = make([]int, len(slices))
	}
	dp[1][0] = 0
	dp[1][1] = slices[1]
	dp[2][1] = max(slices[1], slices[2])
	dp[2][2] = math.MinInt32
	for i := 3; i < len(slices); i++ {
		for j := 1; j <= (i+1)/2; j++ {
			dp[i][j] = max(slices[i]+dp[i-2][j-1], dp[i-1][j])
		}
	}
	res = max(res, dp[len(slices)-1][len(slices)/3])
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
