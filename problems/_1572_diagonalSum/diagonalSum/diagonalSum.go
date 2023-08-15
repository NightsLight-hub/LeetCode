/*
Package diagonalSum
@Time : 2023/8/11 9:34
@Author : sunxy
@File : diagonalSum
@description:
*/
package diagonalSum

func diagonalSum(mat [][]int) int {
	n := len(mat)
	sum := 0
	for i := 0; i < n; i++ {
		sum += mat[i][i]
		sum += mat[n-1-i][i]
	}
	sum -= mat[n/2][n/2] * (n & 1)
	return sum
}
