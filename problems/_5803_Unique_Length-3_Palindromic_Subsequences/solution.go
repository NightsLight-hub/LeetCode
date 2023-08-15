/*
@Time : 2021/7/11 11:14
@Author : sunxy
@File : solution
@description:
*/
package _5803_Unique_Length_3_Palindromic_Subsequences

func countPalindromicSubsequence(s string) int {
	bs := []byte(s)
	l := len(bs)
	ans := 0
	firstIndexMap := make(map[byte]int)
	lastIndexMap := make(map[byte]int)

	for i, b := range bs {
		if i, ok := firstIndexMap[b]; !ok {
			firstIndexMap[b] = i
		}
		lastIndexMap[b] = i
	}

	tmp := make(map[byte]struct{})
	for i := 0; i < l; i++ {
		if _, ok := tmp[bs[i]]; ok {
			continue
		}
		if lastIndexMap[bs[i]]-i > 1 {
			ans += getUniqueCharCount(bs, i, lastIndexMap[bs[i]])
		}
		tmp[bs[i]] = struct{}{}
	}
	return ans
}

func getUniqueCharCount(bs []byte, start, end int) int {
	res := 0
	tmp := make(map[byte]struct{})
	for i := start + 1; i < end; i++ {
		if _, ok := tmp[bs[i]]; ok {
			continue
		} else {
			res++
			tmp[bs[i]] = struct{}{}
		}
	}
	return res
}
