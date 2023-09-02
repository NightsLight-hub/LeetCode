/*
Package _2511_captureForts
@Time : 2023/9/2 9:19
@Author : sunxy
@File : captureForts
@description: https://leetcode.cn/problems/maximum-enemy-forts-that-can-be-captured/?envType=daily-question&envId=2023-09-02
*/
package _2511_captureForts

func captureForts(forts []int) int {
	if len(forts) == 1 {
		return 0
	}
	res := 0
	i, j := 0, 0
	for i < len(forts) {
		if forts[i] == 0 {
			i++
		} else {
			break
		}
	}
	j = i + 1
	for i < len(forts) && j < len(forts) {
		if forts[j] == 0 {
			j++
		} else if forts[j] == -forts[i] {
			res = max(res, abs(i-j)-1)
			i, j = j, j+1
		} else {
			// j = 1
			i, j = j, j+1
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
