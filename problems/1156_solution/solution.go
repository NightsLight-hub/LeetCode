/*
Package _156_solution
@Time : 2023/6/3 7:33
@Author : sunxy
@File : solution
@description:
https://leetcode.cn/problems/swap-for-longest-repeated-character-substring/
*/
package _156_solution

func maxRepOpt1(text string) int {
	ccount := make(map[rune]int)
	for _, c := range text {
		ccount[c]++
	}
	i := 0
	res := 1
	for {
		j := i
		for j < len(text) && text[j] == text[i] {
			j++
		}
		k := j + 1
		res = max(res, j-i)
		for k < len(text) && text[k] == text[i] {
			k++
		}
		res = max(res, min(k-i, ccount[rune(text[i])]))
		i = j
		if i >= len(text)-1 {
			break
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
