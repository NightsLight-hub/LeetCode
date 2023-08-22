/*
Package _849_maxDistToClosest
@Time : 2023/8/22 9:35
@Author : sunxy
@File : maxDistToClosest
@description: https://leetcode.cn/problems/maximize-distance-to-closest-person/
*/
package _849_maxDistToClosest

import "math"

func maxDistToClosest(seats []int) int {
	res := math.MinInt32
	i := 0
	for ; i < len(seats); i++ {
		if seats[i] == 1 {
			res = max(res, i)
			break
		}
	}
	l := i
	i++
	for ; i < len(seats); i++ {
		if seats[i] == 1 {
			res = max(res, (i-l)/2)
			l = i
		}
	}
	res = max(res, len(seats)-l-1)
	return res
}

// O(m*n) ,  m ones and n zeros
// func maxDistToClosest(seats []int) int {
// 	res := math.MinInt32
// 	var index1 []int
// 	for i, s := range seats {
// 		if s == 1 {
// 			index1 = append(index1, i)
// 		}
// 	}
// 	for i, s := range seats {
// 		if s == 0 {
// 			i1 := 0
// 			dist := 0
// 			for ; i1 < len(index1); i1++ {
// 				if index1[i1] > i {
// 					break
// 				}
// 			}
// 			if i1 == len(index1) {
// 				// 这个0 右侧没有1
// 				dist = i - index1[i1-1]
// 			} else if i1 == 0 {
// 				// 0 左侧没有 1
// 				dist = index1[i1] - i
// 			} else {
// 				// 0在两个1 中间
// 				dist = min(i-index1[i1-1], index1[i1]-i)
// 			}
// 			res = max(res, dist)
// 		}
// 	}
// 	return res
// }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
