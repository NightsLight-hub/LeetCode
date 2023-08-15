/*
Package _170_CompareStringsbyFrequencyoftheSmallestCharacter
@Time : 2023/6/10 9:54
@Author : sunxy
@File : Solutionn_test.go
@description:
*/
package _170_CompareStringsbyFrequencyoftheSmallestCharacter

import (
	"testing"
)

func Test_numSmallerByFrequency(t *testing.T) {
	quires := []string{"cbbd"}
	words := []string{"zaaaz"}
	t.Log(numSmallerByFrequency(quires, words))

	quires = []string{"bbb", "cc"}
	words = []string{"a", "aa", "aaa", "aaaa"}
	t.Log(numSmallerByFrequency(quires, words))
}
