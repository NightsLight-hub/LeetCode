/*
Package quicksort
@Time : 2023/7/28 15:05
@Author : sunxy
@File : quickSort_test.go
@description:
*/
package quicksort

import "testing"

func Test_quickSort(t *testing.T) {
	nums := []int{1, 3, 2, 5, 4, 100, 89, 176, 19, 299, 300}
	// nums := []int{3, 4, 2, 6, 5}
	quickSort(nums)
	t.Log(nums)
}
