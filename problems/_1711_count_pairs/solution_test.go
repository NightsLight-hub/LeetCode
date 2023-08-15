/*
@Time : 2021/7/7 10:16
@Author : sunxy
@File : solution_test.go
@description:
*/
package _1711_count_pairs

import (
	"testing"
)

func Test_countPairs(t *testing.T) {
	deliciousness := []int{1, 3, 5, 7, 9}
	t.Log(countPairs(deliciousness))
}
