/*
Package _931_minFallingPathSum
@Time : 2023/7/13 9:21
@Author : sunxy
@File : minFallingPathSum
@description:
*/
package _931_minFallingPathSum

// 正方形
func minFallingPathSum(matrix [][]int) int {
	var res, l = (1 << 31) - 1, len(matrix)
	for i := 1; i < l; i++ {
		for j := 0; j < l; j++ {
			mins := matrix[i-1][j]
			if j-1 >= 0 && matrix[i-1][j-1] < mins {
				mins = matrix[i-1][j-1]
			}
			if j+1 < l && matrix[i-1][j+1] < mins {
				mins = matrix[i-1][j+1]
			}
			matrix[i][j] += mins
		}
	}
	for i := 0; i < l; i++ {
		if res > matrix[l-1][i] {
			res = matrix[l-1][i]
		}
	}
	return res
}
