/*
@Time : 2021/7/9 10:25
@Author : sunxy
@File : solution_test.go
@description:
*/
package Find_Majority_Element_LCCI

import "testing"

func Test_majorityElement(t *testing.T) {
	nums := []int{2, 2, 1}
	t.Log(majorityElement(nums))
}
