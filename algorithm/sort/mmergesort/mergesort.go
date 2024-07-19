/*
Package mmergesort
@Time : 2024/7/11 9:58
@Author : sunxy
@File : mergesort
@description:
*/
package mmergesort

func MergeSort(arr []int) []int {
	// 归并排序
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	return merge(left, right)
}
func merge(left, right []int) []int {
	var result []int
	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	if len(left) > 0 {
		result = append(result, left...)
	} else if len(right) > 0 {
		result = append(result, right...)
	}
	return result
}
