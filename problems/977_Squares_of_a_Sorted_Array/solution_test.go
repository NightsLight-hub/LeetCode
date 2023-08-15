/*
@Time : 2021/7/3 15:08
@Author : sunxy
@File : solution_test.go
@description:
*/
package _77_Squares_of_a_Sorted_Array

import (
	"testing"
)

func Test_sortedSquares(t *testing.T) {
	nums := []int{-7, -3, 2, 3, 11}
	t.Log(sortedSquares(nums))
}
