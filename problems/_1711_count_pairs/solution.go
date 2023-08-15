/*
@Time : 2021/7/7 10:14
@Author : sunxy
@File : solution
@description:
*/
package _1711_count_pairs

func countPairs(deliciousness []int) int {
	var res int
	max := deliciousness[0]
	for _, d := range deliciousness[1:] {
		if max < d {
			max = d
		}
	}

	cnt := make(map[int]int)

	for _, n := range deliciousness {
		for sum := 1; sum <= (max * 2); sum <<= 1 {
			res += cnt[sum-n]
		}
		cnt[n]++
	}
	return res % (1e9 + 7)
}
