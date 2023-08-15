/*
Package HJ103
@Time : 2023/7/25 11:20
@Author : sunxy
@File : main
@description:
*/
package main

import (
	"fmt"
)

/*
*
dp[i] 表示以nums[i]结尾的中最长递增子序列
d[i] = max(d[i], dp[j]+1 (nums[i)>nums[j])) 0<=j<i
*/
func main() {
	n := 0
	fmt.Scanln(&n)
	if n == 1 {
		fmt.Println(1)
	}
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	res := 0
	for _, v := range dp {
		res = max(res, v)
	}
	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
