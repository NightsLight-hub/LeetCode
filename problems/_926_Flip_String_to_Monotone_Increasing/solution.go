/*
Package _926_Flip_String_to_Monotone_Increasing
@Time : 2022/6/12 9:41
@Author : sunxy
@File : solution
@description:
*/
package _926_Flip_String_to_Monotone_Increasing

import "math"

// https://leetcode.cn/problems/flip-string-to-monotone-increasing/solution/
// 使用回溯法，超时。。。参考官方的 动态规划吧

func minFlipsMonoIncr(s string) int {
	var res int = math.MaxInt32
	sb := []byte(s)
	var check func(int, byte, int)
	// 判断当前位要不要变， 可变的值范围是由前一位的值约束的
	check = func(index int, previousValue byte, count int) {
		if index == len(sb) {
			res = min(res, count)
			return
		}
		for i := previousValue; i <= '1'; i++ {
			if sb[index] != i {
				check(index+1, i, count+1)
			} else {
				check(index+1, i, count)
			}
		}
	}
	if sb[0] == '0' {
		check(1, '0', 0)
		check(1, '1', 1)
	} else {
		check(1, '0', 1)
		check(1, '1', 0)
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
