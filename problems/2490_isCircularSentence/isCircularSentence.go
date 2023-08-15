/*
Package _490_isCircularSentence
@Time : 2023/6/30 9:29
@Author : sunxy
@File : isCircularSentence
@description:
*/
package _490_isCircularSentence

func isCircularSentence(sentence string) bool {
	l := len(sentence)
	if sentence[0] != sentence[l-1] {
		return false
	}
	for i := 0; i < l; i++ {
		if sentence[i] == ' ' {
			if sentence[i-1] != sentence[i+1] {
				return false
			}
		}
	}
	return true
}
