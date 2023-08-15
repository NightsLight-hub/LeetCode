/*
@Time : 2021/7/1 11:19
@Author : sunxy
@File : kthLargestElement
@description:
*/
package _15_kth_largest_element

func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, k)
}

func quickSelect(nums []int, index int) int {
	start := 0
	end := len(nums)
	for start != index {
		for i := start + 1; i < end; i++ {
			if nums[start] < nums[i] {
				nums[start], nums[i] = nums[i], nums[start]
			}
		}
		start++
	}
	return nums[index-1]
}

//func quickSelect(a []int, l, r, index int) int {
//	q := randomPartition(a, l, r)
//	if q == index {
//		return a[q]
//	} else if q < index {
//		return quickSelect(a, q + 1, r, index)
//	}
//	return quickSelect(a, l, q - 1, index)
//}
//
//func randomPartition(a []int, l, r int) int {
//	i := rand.Int() % (r - l + 1) + l
//	a[i], a[r] = a[r], a[i]
//	return partition(a, l, r)
//}
//
//func partition(a []int, l, r int) int {
//	x := a[r]
//	i := l - 1
//	for j := l; j < r; j++ {
//		if a[j] <= x {
//			i++
//			a[i], a[j] = a[j], a[i]
//		}
//	}
//	a[i+1], a[r] = a[r], a[i+1]
//	return i + 1
//}
