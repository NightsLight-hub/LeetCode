/*
@Time : 2021/7/11 10:37
@Author : sunxy
@File : solution
@description:
*/
package _274_hIndex

import "sort"

func hIndex(citations []int) int {
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})
	for i := 0; i < len(citations); i++ {
		if citations[i] < i+1 {
			return i
		}
	}
	return len(citations)
}
