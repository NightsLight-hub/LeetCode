/*
@Time : 2021/7/1 11:59
@Author : sunxy
@File : solution
@description:
*/
package __Longest_Substring_Without_Repeating_Characters

func lengthOfLongestSubstring(s string) int {
	bs := []byte(s)
	bsl := len(bs)
	if bsl <= 1 {
		return bsl
	}
	i := 0
	cursor, res := 1, 1
	for i < bsl && i != bsl-1 {
		for ; ; cursor++ {
			ind := bsIndex(bs[i:cursor], bs[cursor])
			if ind >= 0 {
				res = max(cursor-i, res)
				i = ind + i + 1
				break
			}
			if cursor == bsl-1 {
				res = max(cursor-i+1, res)
				return res
			}
		}
	}
	return res
}

func bsIndex(bs []byte, b byte) int {
	for j, i := range bs {
		if i == b {
			return j
		}
	}
	return -1
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
