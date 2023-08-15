/*
Package diagonalSum
@Time : 2023/8/11 9:37
@Author : sunxy
@File : diagonalSum_test.go
@description:
*/
package diagonalSum

import "testing"

func Test_diagonalSum(t *testing.T) {
	mat := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	if diagonalSum(mat) != 25 {
		t.Fatal("wrong")
	}
	mat = [][]int{{1, 2, 3, 4}, {4, 5, 6, 8}, {7, 8, 9, 10}, {11, 12, 13, 14}}
	t.Log(diagonalSum(mat))
}
