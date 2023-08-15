/*
Package _485_pivotInteger
@Time : 2023/6/26 9:38
@Author : sunxy
@File : pivotInteger
@description: https://leetcode.cn/problems/find-the-pivot-integer/
*/
package _485_pivotInteger

/*
*
1.
*/
func pivotInteger(n int) int {
	if n == 1 {
		return 1
	}
	c := 2
	l := 3
	r := 0
	for i := c; i <= n; i++ {
		r += i
	}
	for {
		if l == r {
			return c
		}
		c++
		if c > n {
			return -1
		}
		l += c
		r = r - c + 1
	}
}
