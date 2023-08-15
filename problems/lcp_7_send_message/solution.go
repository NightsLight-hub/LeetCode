/*
@Time : 2021/7/1 12:54
@Author : sunxy
@File : solution
@description:
https://leetcode-cn.com/problems/chuan-di-xin-xi/
*/
package lcp_7_send_message

type result struct {
	m     map[int][]int
	res   int
	start int
}

func numWays(n int, relation [][]int, k int) int {
	res := new(result)
	res.m = makeInitialMap(n, relation)

	findWay(n-1, k, res)
	return res.res
}

func findWay(end int, rounds int, res *result) {
	ps := res.m[end]
	if rounds == 1 {
		for _, p := range ps {
			if p == res.start {
				res.res++
			}
		}
	} else {
		for _, p := range ps {
			findWay(p, rounds-1, res)
		}
	}
}

func makeInitialMap(n int, relation [][]int) map[int][]int {
	m := make(map[int][]int)
	for i := 0; i < n; i++ {
		m[i] = make([]int, 0)
	}
	for _, r := range relation {
		m[r[1]] = append(m[r[1]], r[0])
	}
	return m
}
