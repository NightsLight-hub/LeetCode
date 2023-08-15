/*
@Time : 2021/5/27 15:44
@Author : sunxy
@File : hamming-distance
@description:
*/
package Hamming_distance

import (
	"fmt"
	"strings"
)

// 字符串比较，终究是落了下乘。。。
func hammingDistance(x int, y int) int {
	r := x ^ y
	rs := fmt.Sprintf("%b", r)
	return strings.Count(rs, "1")
}

// 优秀的做法
//func hammingDistance(x int, y int) int {
//	count := 0
//	temp := x ^ y
//	for temp != 0 {
//		if temp&1 == 1 {
//			count++
//		}
//		temp >>= 1
//	}
//	return count
//}
