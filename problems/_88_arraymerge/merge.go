/*
Package _88_arraymerge
@Time : 2023/8/13 9:09
@Author : sunxy
@File : merge.go
@description:
*/
package _88_arraymerge

// 逆向双指针，可以不用临时空间
// https://leetcode.cn/problems/merge-sorted-array/description/

func merge(nums1 []int, m int, nums2 []int, n int) {
	temp := make([]int, m+n)
	i, i1, i2 := 0, 0, 0
	for i1 != m && i2 != n {
		if nums1[i1] < nums2[i2] {
			temp[i] = nums1[i1]
			i1++
		} else {
			temp[i] = nums2[i2]
			i2++
		}
		i++
	}
	if i2 != n {
		copy(temp[i:], nums2[i2:])
	}
	if i1 != m {
		copy(temp[i:], nums1[i1:])
	}
	copy(nums1, temp)
}
