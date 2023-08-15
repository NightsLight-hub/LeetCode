/*
Package _352_EqualRowAndColumnPairs
@Time : 2023/6/6 20:21
@Author : sunxy
@File : Solution
@description: https://leetcode.cn/problems/equal-row-and-column-pairs/
*/
package _352_EqualRowAndColumnPairs

import (
	"strconv"
	"strings"
)

func equalPairs(grid [][]int) int {
	// res := 0
	// var m = len(grid)
	// for i := 0; i < m; i++ {
	// 	for j := 0; j < m; j++ {
	// 		flag := true
	// 		for k := 0; k < m; k++ {
	// 			if grid[i][k] != grid[k][j] {
	// 				flag = false
	// 				break
	// 			}
	// 		}
	// 		if flag {
	// 			res++
	// 		}
	// 	}
	// }
	// return res
	var m = len(grid)
	var res = 0
	var rowMap = make(map[string]int)
	for i := 0; i < m; i++ {
		s := strings.Builder{}
		for j := 0; j < m; j++ {
			s.Write([]byte(strconv.Itoa(grid[i][j])))
			s.Write([]byte(","))
		}
		rowMap[(s.String())]++
	}
	var colMap = make(map[string]int)
	for i := 0; i < m; i++ {
		s := strings.Builder{}
		for j := 0; j < m; j++ {
			s.Write([]byte(strconv.Itoa(grid[j][i])))
			s.Write([]byte(","))
		}
		colMap[(s.String())]++
	}
	for k, v := range rowMap {
		for k2, v2 := range colMap {
			if k == k2 {
				res += v * v2
			}
		}
	}
	return res
}
