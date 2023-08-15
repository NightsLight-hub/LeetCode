/*
@Time : 2021/7/5 18:03
@Author : sunxy
@File : solution_test.go
@description:
*/
package _283_movezero

import "testing"

func Test_moveZeroes(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	t.Logf("%+v", nums)
}
