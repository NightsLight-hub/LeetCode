/*
@Time : 2021/7/4 7:57
@Author : sunxy
@File : solution_test.go
@description:
*/
package _45_Set_Mismatch

import (
	"testing"
)

func Test_findErrorNums(t *testing.T) {
	nums := []int{3, 2, 2}
	t.Log(findErrorNums(nums))
}
