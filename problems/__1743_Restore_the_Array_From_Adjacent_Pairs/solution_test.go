/*
@Time : 2021/7/25 8:28
@Author : sunxy
@File : solution_test.go
@description:
*/
package __1743_Restore_the_Array_From_Adjacent_Pairs

import (
	"testing"
)

func Test_restoreArray(t *testing.T) {
	input := [][]int{{4, -2}, {1, 4}, {-3, 1}}
	t.Logf("%+v", restoreArray(input))
}
