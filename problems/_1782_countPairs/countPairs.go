/*
Package _1782_countPairs
@Time : 2023/8/23 10:16
@Author : sunxy
@File : countPairs
@description: https://leetcode.cn/problems/count-pairs-of-nodes/description/
*/
package _1782_countPairs

import "sort"

func countPairs(n int, edges [][]int, queries []int) []int {
	degree := make([]int, n)
	cnt := map[int]int{}
	for _, edge := range edges {
		x, y := edge[0]-1, edge[1]-1
		if x > y {
			x, y = y, x
		}
		degree[x]++
		degree[y]++
		cnt[x*n+y]++
	}

	arr := make([]int, n)
	copy(arr, degree)
	sort.Ints(arr)
	ans := []int{}
	for _, bound := range queries {
		total := 0
		for i := 0; i < n; i++ {
			j := sort.SearchInts(arr, bound-arr[i]+1)
			total += n - max(i+1, j)
		}
		for val, freq := range cnt {
			x, y := val/n, val%n
			if degree[x]+degree[y] > bound && degree[x]+degree[y]-freq <= bound {
				total--
			}
		}
		ans = append(ans, total)
	}
	return ans
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
