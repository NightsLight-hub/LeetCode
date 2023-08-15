/*
@Time : 2021/5/28 11:48
@Author : sunxy
@File : a.go
@description:
*/
package Longest_Substring_Without_Repeating_Characters

/*
*
checkRepeat 的方法可以改用map 代替， map记录每个字符是否出现过，考虑到map的开销，实际优化程度未知

cursor 是子串的头位置，i 是字串的最后位置，只需要判断末尾是否与前面字符重复
如果重复，cursor走到重复位置的下一位。原因是出现以重复字符左侧字串开头的任何字串，长度不可能超过刚刚失效的那个字符串

deabac

	| |

以d开头的子串，到s[4] 重复，重复字符a，则直接从a第一次出现位置2之后的位置开始重新查找可能的最长子串即可
*/
func lengthOfLongestSubstring(s string) int {
	bs := []byte(s)
	l := len(bs)
	var repeatIndex, cursor, length int
	for i := 1; i < l; i++ {
		repeatIndex = checkRepeat(bs[cursor:i], bs[i])
		if repeatIndex >= 0 {
			length = max(length, i-cursor)
			cursor += repeatIndex + 1
		}
	}
	length = max(length, l-cursor)
	return length
}

func checkRepeat(bytes []byte, b byte) int {
	for i, c := range bytes {
		if c == b {
			return i
		}
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
