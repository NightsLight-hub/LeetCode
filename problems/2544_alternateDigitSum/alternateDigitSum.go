/*
Package _544_alternateDigitSum
@Time : 2023/7/12 15:09
@Author : sunxy
@File : alternateDigitSum
@description: https://leetcode.cn/problems/alternating-digit-sum/
*/
package _2544_alternateDigitSum

func alternateDigitSum(n int) int {
	l, c := 0, n
	for c != 0 {
		l++
		c /= 10
	}
	flag := true
	if l%2 == 0 {
		flag = false
	}
	res := 0
	for n != 0 {
		if flag {
			res += n % 10
		} else {
			res -= n % 10
		}
		n /= 10
		flag = !flag
	}
	return res
}
