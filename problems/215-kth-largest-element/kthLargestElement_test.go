/*
@Time : 2021/7/1 11:30
@Author : sunxy
@File : kthLargestElement_test.go
@description:
*/
package _15_kth_largest_element

import "testing"

func Test_quickSelect(t *testing.T) {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := 4
	t.Log(findKthLargest(nums, k))
}
