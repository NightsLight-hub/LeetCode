/*
Package _890_Find_and_Replace_Pattern
@Time : 2022/6/12 9:15
@Author : sunxy
@File : solution
@description:  https://leetcode.cn/problems/find-and-replace-pattern/
*/
package _890_Find_and_Replace_Pattern

// time O(nm)  space O(m)
func findAndReplacePattern(words []string, pattern string) []string {
	p := mapString(pattern)
	result := make([]string, 0)
	for _, w := range words {
		if mapString(w) == p {
			result = append(result, w)
		}
	}
	return result
}

func mapString(origin string) string {
	resp := make([]rune, 0)
	originRunes := []rune(origin)
	runeMap := make(map[rune]rune)
	var ph rune = 'a'
	for _, r := range originRunes {
		if v, ok := runeMap[r]; ok {
			resp = append(resp, v)
		} else {
			resp = append(resp, ph)
			runeMap[r] = ph
			ph = ph + 1
		}
	}
	return string(resp)
}
