/*
Package _473_Matchsticks_to_Square
@Time : 2022/6/1 18:38
@Author : sunxy
@File : solution
@description: https://leetcode.cn/problems/matchsticks-to-square/
*/
package _473_Matchsticks_to_Square

import (
	"sort"
)

// 回溯法

func makesquare(matchsticks []int) bool {
	// return solution1(matchsticks)
	return solution2(matchsticks)
}

func solution2(matchsticks []int) bool {
	sum := 0
	for _, n := range matchsticks {
		sum += n
	}
	if sum%4 != 0 {
		return false
	}
	edgeL := sum / 4
	sort.Ints(matchsticks)
	if matchsticks[len(matchsticks)-1] > edgeL {
		return false
	}
	var search func(int, int, *[]int) bool
	// 可能会选取一个不当组合， 由于它让结果失败，最好用个map 记录所有数字是否被选择
	search = func(index, sum int, resultIndexs *[]int) bool {
		for i := index; i < len(matchsticks); i++ {
			if matchsticks[i] < 0 {
				// 已经使用的数字
				continue
			}
			sum += matchsticks[i]
			// 如果 这个边用index索引的值 大于边长，由于数列增序，
			// 后面数字更大，不可能有结果，因此直接return false
			if sum > edgeL {
				return false
			} else if sum < edgeL {
				// 如果用了这个数字，小于变长，可以递归下去，搜索
				res := search(index+1, sum, resultIndexs)
				if res {
					// 子结果成功，记录下当前的数字索引
					*resultIndexs = append(*resultIndexs, i)
					return true
				} else {
					// 递归失败，说明这个数字不能用，减掉数字，走for循环换别的数字
					sum -= matchsticks[i]
				}
			} else {
				// sum == edgeL，成功了
				*resultIndexs = append(*resultIndexs, i)
				return true
			}
		}
		return false
	}
	resIndexs := make([]int, 0)
	for i := 0; i < 4; i++ {
		if search(0, 0, &resIndexs) {
			for _, j := range resIndexs {
				matchsticks[j] = -1
			}
			resIndexs = make([]int, 0)
		} else {
			return false
		}
	}
	return true
}

func solution1(matchsticks []int) bool {
	sum := 0
	for _, n := range matchsticks {
		sum += n
	}
	if sum%4 != 0 {
		return false
	}
	edgeL := sum / 4
	sort.Ints(matchsticks)
	sort.Sort(sort.Reverse(sort.IntSlice(matchsticks)))
	if matchsticks[len(matchsticks)-1] > edgeL {
		return false
	}
	var edges [4]int
	var dfs func(int) bool
	dfs = func(index int) bool {
		if index == len(matchsticks) {
			return true
		}
		for i := range edges {
			edges[i] += matchsticks[index]
			if edges[i] <= edgeL && dfs(index+1) {
				return true
			}
			edges[i] -= matchsticks[index]
		}
		return false
	}
	return dfs(0)
}
