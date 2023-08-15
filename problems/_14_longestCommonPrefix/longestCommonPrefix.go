/*
Package _14_longestCommonPrefix
@Time : 2023/8/4 17:34
@Author : sunxy
@File : longestCommonPrefix
@description:
*/
package _14_longestCommonPrefix

func longestCommonPrefix(strs []string) string {
	lm := make(map[string]int)
	for _, s := range strs {
		lm[s] = len(s)
	}
	for i := 0; ; i++ {
		for _, s := range strs {
			if i >= lm[s] || strs[0][i] != s[i] {
				return strs[0][:i]
			}
		}
	}
}
