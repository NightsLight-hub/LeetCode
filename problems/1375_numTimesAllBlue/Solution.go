/*
Package _375_numTimesAllBlue
@Time : 2023/6/14 17:23
@Author : sunxy
@File : Solution
@description:
*/
package _375_numTimesAllBlue

// 利用前缀和来简化计算，O(n)
func numTimesAllBlue(flips []int) int {
	res := 0
	count := 0
	countBoundary := 0
	f := make([]int, len(flips)+1)
	s := make([]int, len(flips)+1)
	copy(f[1:], flips)
	for i := 1; i < len(f); i++ {
		s[f[i]] = 1
		if f[i] <= countBoundary {
			count++
		}
		countBoundary = i
		if s[countBoundary] == 1 {
			count++
		}
		if count == i {
			res++
		}
	}
	return res
}
