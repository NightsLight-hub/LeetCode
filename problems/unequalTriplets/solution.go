/*
Package unequalTriplets
@Time : 2023/6/13 10:58
@Author : sunxy
@File : solution
@description:
*/
package unequalTriplets

func unequalTriplets(nums []int) int {
	var counts = make(map[int]int)
	for _, n := range nums {
		counts[n]++
	}

	t := 0 // 当前元素之前的元素数量
	l := len(nums)
	res := 0
	for _, v := range counts {
		// v 是当前元素的数量
		res += t * v * (l - t - v)
		t += v
	}
	return res
}
