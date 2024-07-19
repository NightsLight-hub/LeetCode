/*
Package __reverse
@Time : 2024/7/9 12:30
@Author : sunxy
@File : reverse
@description: https://leetcode.cn/problems/reverse-integer/
*/
package __reverse

func reverse(x int) int {
	// 反转x
	var res int
	for x != 0 {
		res = res*10 + x%10
		x = x / 10
	}
	if res > 0x7fffffff || res < -0x80000000 {
		return 0
	}
	return res
}
