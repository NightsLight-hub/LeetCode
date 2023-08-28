/*
Package _56_merge
@Time : 2023/8/27 13:18
@Author : sunxy
@File : merge
@description:
*/
package _56_merge

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := make([][]int, 0)
	i, j := 0, 1
	currentRightBoundary := intervals[i][1]
	for j < len(intervals) {
		if intervals[j][0] <= currentRightBoundary {
			currentRightBoundary = max(currentRightBoundary, intervals[j][1])
		} else {
			res = append(res, []int{intervals[i][0], currentRightBoundary})
			i = j
			currentRightBoundary = intervals[i][1]
		}
		j++
	}
	res = append(res, []int{intervals[i][0], currentRightBoundary})
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
