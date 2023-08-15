/*
@Time : 2021/7/11 10:43
@Author : sunxy
@File : solution_test.go
@description:
*/
package _274_hIndex

import "testing"

func Test_hIndex(t *testing.T) {
	nums := []int{0, 1, 3, 5, 6}
	t.Log(hIndex(nums))
}
