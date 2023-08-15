/*
Package _344_reverseString
@Time : 2023/8/7 9:01
@Author : sunxy
@File : reverseString
@description:
*/
package _344_reverseString

func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
