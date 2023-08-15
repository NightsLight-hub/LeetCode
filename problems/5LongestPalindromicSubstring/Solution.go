/*
Package _LongestPalindromicSubstring
@Time : 2023/6/5 11:13
@Author : sunxy
@File : Solution
@description:
*/
package _LongestPalindromicSubstring

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	if len(s) == 2 {
		if s[0] == s[1] {
			return s
		} else {
			return s[0:1]
		}
	}
	var l = len(s)
	var dp = make([][]bool, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]bool, l)
		dp[i][i] = true
	}
	maxl := 0
	left, right := 0, 0
	for i := 0; i < l-1; i++ {
		dp[i][i+1] = s[i] == s[i+1]
		if dp[i][i+1] {
			maxl = 2
			left, right = i, i+1
		}
	}
	for window := 3; window <= l; window++ {
		for wl := 0; wl+window-1 < l; wl++ {
			dp[wl][wl+window-1] = dp[wl+1][wl+window-2] && s[wl] == s[wl+window-1]
			if dp[wl][wl+window-1] && window > maxl {
				maxl = window
				left = wl
				right = wl + window - 1
			}
		}
	}

	return s[left : right+1]
}

func Reverse(s string) string {
	a := []rune(s)
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return string(a)
}
