/*
Package sortandsearch
@Time : 2023/8/23 12:17
@Author : sunxy
@File : test1.go
@description:
*/
package sortandsearch

import (
	"sort"
	"testing"
)

func Test_1(t *testing.T) {
	//         0  1  2  3  4  5  6  7  8  9 10 11
	s := []int{1, 2, 3, 4, 5, 5, 5, 5, 5, 6, 7, 8}
	sort.Ints(s)

	t.Log(binarySearchLeft(s, 8) == 11)
	t.Log(binarySearchLeft(s, 7) == 10)
	t.Log(binarySearchLeft(s, 6) == 9)
	t.Log(binarySearchLeft(s, 5) == 4)
	t.Log(binarySearchLeft(s, 1) == 0)

	t.Log(binarySearchRight(s, 8) == 11)
	t.Log(binarySearchRight(s, 7) == 10)
	t.Log(binarySearchRight(s, 6) == 9)
	t.Log(binarySearchRight(s, 5) == 8)
	t.Log(binarySearchRight(s, 1) == 0)

	s2 := []int{1, 2}
	t.Log(binarySearchRight(s2, 2) == 1)
}

func binarySearchRight(s []int, a int) int {
	i, j := 0, len(s)-1
	for i < j {
		mid := int(uint(i+j)>>1 + 1)
		if s[mid] <= a {
			i = mid
		} else {
			j = mid - 1
		}
	}
	if s[i] != a {
		return -1
	}
	return i
}

func binarySearchLeft(s []int, a int) int {
	i, j := 0, len(s)-1
	for i < j {
		mid := int(uint(i+j) >> 1)
		if s[mid] >= a {
			j = mid
		} else {
			i = mid + 1
		}
	}
	if s[i] != a {
		return -1
	}
	return i
}
