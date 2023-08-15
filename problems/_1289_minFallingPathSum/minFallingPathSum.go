/*
Package _1289_minFallingPathSum
@Time : 2023/8/10 9:41
@Author : sunxy
@File : minFallingPathSum
@description:
*/
package _1289_minFallingPathSum

/*
*
f[x][y] = g[0][y]     if  x==0

	min(f[x-1][y']) + g[x][y]      if x>0 && y' !=y
*/
func minFallingPathSum(grid [][]int) int {
	n := len(grid)
	if n == 1 {
		return grid[0][0]
	}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	} // O(n)
	// f[x][y] = g[0][y]     if  x==0
	for i := 0; i < n; i++ {
		dp[0][i] = grid[0][i]
	} // O(n)
	// min(f[x-1][y']) + g[x][y]      if x>0 && y' !=y
	for x := 1; x < n; x++ {
		for y := 0; y < n; y++ {
			dp[x][y] = minx(dp[x-1], y) + grid[x][y]
		}
	} // O(n^3)
	return minx(dp[n-1], -1)
}

func minx(nums []int, except int) int {
	min := 1 << 31
	for i := range nums {
		if i != except && min > nums[i] {
			min = nums[i]
		}
	}
	return min
}
