/*
@Time : 2021/7/11 10:58
@Author : sunxy
@File : solution_test.go
@description:
*/
package _808_Array_Concatenation

import (
	"testing"
)

func Test_getConcatenation(t *testing.T) {
	nums := []int{1, 2, 1}
	t.Logf("%+v", getConcatenation(nums))
}
