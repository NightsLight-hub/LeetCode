/*
Package _2569_handleQuery
@Time : 2023/7/26 9:44
@Author : sunxy
@File : handleQuery
@description:
*/
package _2569_handleQuery

func handleQuery(nums1 []int, nums2 []int, queries [][]int) []int64 {
	var lastSum int64
	var have2 bool = true
	var result = make([]int64, 0)
	var h1 = func(l, r int) {
		for i := l; i <= r; i++ {
			if nums1[i] == 0 {
				nums1[i] = 1
			} else {
				nums1[i] = 0
			}
		}
	}
	var h2 = func(p int) {
		for i, _ := range nums2 {
			nums2[i] = nums2[i] + p*nums1[i]
		}
	}
	var h3 = func() {
		if have2 {
			var s int64
			for _, v := range nums2 {
				s += int64(v)
			}
			result = append(result, s)
		} else {
			result = append(result, lastSum)
		}
	}

	for _, q := range queries {
		switch q[0] {
		case 1:
			h1(q[1], q[2])
		case 2:
			h2(q[1])
		case 3:
			h3()
		}
	}
	return result
}
