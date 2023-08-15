/*
Package _17_11_find_closest_LCCI
@Time : 2022/5/27 8:38
@Author : sunxy
@File : solution
@description:
*/
package _17_11_find_closest_LCCI

// todo  待优化
/**
有个内含单词的超大文本文件，给定任意两个不同的单词，找出在这个文件中这两个单词的最短距离(相隔单词数)。如果寻找过程在这个文件中会重复多次，而每次寻找的单词不同，你能对此优化吗?
*/
func findClosest(words []string, word1 string, word2 string) int {
	// hold all index of nums
	mi := make(map[string][]int)
	for i, w := range words {
		if _, ok := mi[w]; !ok {
			mi[w] = make([]int, 0)
		}
		mi[w] = append(mi[w], i)
	}
	return findTwoOrderArrayMinDifference(mi[word1], mi[word2])
}

func findTwoOrderArrayMinDifference(a, b []int) int {
	var difference int = 1<<31 - 1
	ai, bi := 0, 0
	for ai < len(a) && bi < len(b) {
		difference = min(difference, absD(a[ai], b[bi]))
		if a[ai] < b[bi] {
			ai++
		} else {
			bi++
		}
	}
	return difference
}

func absD(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
