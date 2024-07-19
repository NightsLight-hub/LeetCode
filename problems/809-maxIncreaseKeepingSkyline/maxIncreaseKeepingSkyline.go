/*
Package _09_maxIncreaseKeepingSkyline
@Time : 2024/7/14 10:39
@Author : sunxy
@File : maxIncreaseKeepingSkyline
@description: https://leetcode.cn/problems/max-increase-to-keep-city-skyline/description/
*/
package _809_maxIncreaseKeepingSkyline

func maxIncreaseKeepingSkyline(grid [][]int) int {
	// 计算水平方向 的最高值， 垂直方向的最高值
	var maxHorizontalH = make([]int, len(grid))
	var maxVerticalH = make([]int, len(grid[0]))
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			// 计算水平方向 的最高值
			if j == 0 {
				maxHorizontalH[i] = grid[i][j]
			} else {
				if grid[i][j] > maxHorizontalH[i] {
					maxHorizontalH[i] = grid[i][j]
				}
			}
			if i == 0 {
				maxVerticalH[j] = grid[i][j]
			} else {
				if grid[i][j] > maxVerticalH[j] {
					maxVerticalH[j] = grid[i][j]
				}
			}
		}
	}
	// 遍历每个点，根据坐标， 计算出该点可以增加的高度，不超过水平与垂直方向最高点的min值
	var sum = 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			sum += min(maxHorizontalH[i], maxVerticalH[j]) - grid[i][j]
		}
	}
	// 计算增加的高度和
	return sum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
