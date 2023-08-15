/*
@Time : 2021/5/27 16:12
@Author : sunxy
@File : search_2d_matrix_test.go
@description:
*/
package Search_a_2D_Matrix

import "testing"

func Test_searchMatrix(t *testing.T) {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	target := 13
	t.Logf("%v", searchMatrix(matrix, target))
}
