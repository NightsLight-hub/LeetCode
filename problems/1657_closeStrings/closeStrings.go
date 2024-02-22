/*
Package _1657_closeStrings
@Time : 2023/11/30 9:47
@Author : sunxy
@File : closeStrings.go
@description: https://leetcode.cn/problems/determine-if-two-strings-are-close/
*/
package _1657_closeStrings

import "sort"

func closeStrings(word1 string, word2 string) bool {
	// sort word1  and word2  => sw1, sw2
	// if sw1 == sw2  then true
	// count every letter of sw1, and then sort the countsArray, if countsArray1 == countsArray2  then true

	wbs1, wbs2 := []byte(word1), []byte(word2)
	sort.Slice(wbs1, func(i, j int) bool { return wbs1[i] < wbs1[j] })
	sort.Slice(wbs2, func(i, j int) bool { return wbs1[i] < wbs1[j] })
	if string(wbs1) == string(wbs2) {
		return true
	}
	letterCounts1 := make([]int, 0)
	letterCounts2 := make([]int, 0)
	var c int = 1
	var lb byte = 0
	for i, _ := range wbs1 {

	}
}
