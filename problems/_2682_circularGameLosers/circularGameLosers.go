/*
Package _2682_circularGameLosers
@Time : 2023/8/16 9:42
@Author : sunxy
@File : circularGameLosers
@description:
*/
package _2682_circularGameLosers

func circularGameLosers(n int, k int) []int {
	var slot = make([]int, n)
	i := 0
	slot[i] = 1
	round := 1
	for {
		i = (i + k*round) % n
		if slot[i] == 1 {
			break
		} else {
			slot[i] = 1
			round++
		}
	}
	var res []int
	for index, v := range slot {
		if v == 0 {
			res = append(res, index+1)
		}
	}
	return res
}
