/*
Package _7_insert
@Time : 2023/8/28 17:35
@Author : sunxy
@File : insert
@description: https://leetcode.cn/problems/insert-interval/
*/
package _7_insert

func insert(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0)
	for i := range intervals {
		if newInterval == nil {
			res = append(res, intervals[i])
			continue
		}
		if intervals[i][1] < newInterval[0] {
			res = append(res, intervals[i])
		} else if intervals[i][1] >= newInterval[0] {
			// right place
			if newInterval[1] < intervals[i][0] {
				res = append(res, newInterval)
				// already insert newInterval, set it to nil
				newInterval = nil
				res = append(res, intervals[i])
			} else {
				// overlap
				left := min(intervals[i][0], newInterval[0])
				right := max(intervals[i][1], newInterval[1])
				newInterval = []int{left, right}
			}
		}
	}
	// if newInterval is not nil, then it should be at the end of the intervals
	if newInterval != nil {
		res = append(res, newInterval)
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
