/*
Package _1851_minInterval
@Time : 2023/7/18 10:32
@Author : sunxy
@File : minInterval
@description:
*/
package _1851_minInterval

import "sort"

type Interval struct {
	left  int
	right int
	val   int
}

func minInterval(intervals [][]int, queries []int) []int {
	var is []Interval
	for _, v := range intervals {
		is = append(is, Interval{
			left:  v[0],
			right: v[1],
			val:   v[1] - v[0] + 1,
		})
	}
	sort.Slice(is, func(i, j int) bool {
		return is[i].val < is[j].val
	})
	var res []int = make([]int, len(queries))
	for i := range queries {
		f := -1
		for j := range is {
			if queries[i] >= is[j].left && queries[i] <= is[j].right {
				f = j
				break
			}
		}
		if f != -1 {
			res[i] = is[f].val
		} else {
			res[i] = -1
		}
	}
	return res
}
