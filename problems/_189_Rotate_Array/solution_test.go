/*
@Time : 2021/7/5 17:47
@Author : sunxy
@File : solution_test.go
@description:
*/
package _189_Rotate_Array

import "testing"

func Test_rotate(t *testing.T) {
	nums := []int{-1, -100, 3, 99}
	k := 2
	rotate(nums, k)
	t.Logf("%+v", nums)
}
