/*
Package _611_miceAndCheese
@Time : 2023/6/7 10:00
@Author : sunxy
@File : Solution
@description:https://leetcode.cn/problems/mice-and-cheese/
*/
package _611_miceAndCheese

import "sort"

func miceAndCheese(reward1 []int, reward2 []int, k int) int {
	type cheese struct {
		Index    int
		getScore int
	}
	var cheeses = make([]*cheese, len(reward1))
	for i, _ := range reward1 {
		cheeses[i] = &cheese{
			Index:    i,
			getScore: reward1[i] - reward2[i],
		}
	}
	sort.SliceStable(cheeses, func(i, j int) bool {
		return cheeses[i].getScore > cheeses[j].getScore
	})
	res := 0
	for i := 0; i < k; i++ {
		res += reward1[cheeses[i].Index]
	}
	for i := k; i < len(reward1); i++ {
		res += reward2[cheeses[i].Index]
	}
	return res
}
