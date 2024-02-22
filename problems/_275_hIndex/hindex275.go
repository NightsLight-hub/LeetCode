/*
Package _275_hIndex
@Time : 2023/10/30 10:44
@Author : sunxy
@File : hindex275
@description:
*/
package _275_hIndex

func hIndex(citations []int) int {
	n := len(citations)
	l, r := 0, n-1
	for l <= r {
		mid := (r + l) >> 1
		if citations[mid] >= n-mid {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return n - l
}
