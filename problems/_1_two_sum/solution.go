/*
@Time : 2021/7/6 17:39
@Author : sunxy
@File : solution
@description:
*/
package _1_two_sum

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, n := range nums {
		m[n] = i
	}
	for i, n := range nums {
		v, ok := m[target-n]
		if ok {
			if i == v {
				continue
			} else {
				return []int{i, v}
			}
		}
	}
	return []int{0, 1}
}
