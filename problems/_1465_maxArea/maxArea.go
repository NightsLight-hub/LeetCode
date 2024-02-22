/*
Package _1465_maxArea
@Time : 2023/10/27 17:23
@Author : sunxy
@File : maxArea
@description:
*/
package _1465_maxArea

import "sort"

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)
	maxH, maxV := -1, -1
	for i := range horizontalCuts {
		if i == 0 {
			maxH = horizontalCuts[i]
		} else {
			maxH = max(maxH, horizontalCuts[i]-horizontalCuts[i-1])
		}
	}
	maxH = max(maxH, h-horizontalCuts[len(horizontalCuts)-1])

	for i := range verticalCuts {
		if i == 0 {
			maxV = verticalCuts[i]
		} else {
			maxV = max(maxV, verticalCuts[i]-verticalCuts[i-1])
		}
	}
	maxV = max(maxV, w-verticalCuts[len(verticalCuts)-1])
	return int(int64(maxH) * int64(maxV) % 1000000007)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
