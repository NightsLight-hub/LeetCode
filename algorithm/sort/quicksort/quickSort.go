/*
Package quicksort
@Time : 2023/8/9 9:34
@Author : sunxy
@File : quickSort
@description:
*/
package quicksort

func quickSort(nums []int) {
	if len(nums) < 2 {
		return
	}
	midI := 0
	leftCursor, rightCursor := 1, len(nums)-1
	for leftCursor <= rightCursor {
		if midI < rightCursor {
			if nums[rightCursor] < nums[midI] {
				swap(nums, rightCursor, midI)
				midI = rightCursor
			}
			rightCursor--
		} else {
			if nums[leftCursor] > nums[midI] {
				swap(nums, leftCursor, midI)
				midI = leftCursor
			}
			leftCursor++
		}
	}
	quickSort(nums[:midI])
	quickSort(nums[midI+1:])
}

func swap(nums []int, a, b int) {
	nums[a], nums[b] = nums[b], nums[a]
}
