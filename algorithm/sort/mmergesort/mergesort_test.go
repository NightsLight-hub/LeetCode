/*
Package mmergesort
@Time : 2024/7/11 10:06
@Author : sunxy
@File : mergesort_test.go
@description:
*/
package mmergesort

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	fmt.Println(MergeSort([]int{3, 2, 1, 5, 4, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}))
}

func Test_a(t *testing.T) {
	a := []int{1}
	a = a[1:]
	fmt.Printf("%v, len %d", a, len(a))
	a = a[1:]
	fmt.Println(a)
}
