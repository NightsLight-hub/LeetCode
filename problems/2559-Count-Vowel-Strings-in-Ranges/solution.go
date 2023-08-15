/*
Package _559_Count_Vowel_Strings_in_Ranges
@Time : 2023/6/2 20:20
@Author : sunxy
@File : solution
@description:
*/
package _559_Count_Vowel_Strings_in_Ranges

func vowelStrings(words []string, queries [][]int) []int {
	// 用数组 表示元音
	wi := make([]int, 0)
	for i, word := range words {
		if isVowel(word) {
			wi = append(wi, i)
		}
	}
	res := make([]int, len(queries))
	// for 统计
	for i, q := range queries {
		for _, w := range wi {
			if w >= q[0] && w <= q[1] {
				res[i]++
			}
		}
	}
	return res
}

func isVowel(word string) bool {
	bs := []byte(word)
	if bs[0] == 'a' || bs[0] == 'e' ||
		bs[0] == 'i' || bs[0] == 'o' || bs[0] == 'u' {
		if bs[len(word)-1] == 'a' || bs[len(word)-1] == 'e' || bs[len(word)-1] == 'i' || bs[len(word)-1] == 'o' || bs[len(word)-1] == 'u' {
			return true
		}
	}
	return false
}
