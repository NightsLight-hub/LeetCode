/*
Package _61_repeatedNum
@Time : 2022/5/21 10:15
@Author : sunxy
@File : solution.go
@description:
*/
package _61_repeatedNum

import "math/rand"

// func repeatedNTimes(nums []int) int {
//     m := make(map[int]int)
//     for _, n := range nums {
//         if _, ok := m[n]; ok {
//             return n
//         } else {
//             m[n] = 1
//         }
//     }
//     return -1
// }

func repeatedNTimes(nums []int) int {
	n := len(nums)
	for {
		x, y := rand.Intn(n), rand.Intn(n)
		if x != y && nums[x] == nums[y] {
			return nums[x]
		}
	}
}
