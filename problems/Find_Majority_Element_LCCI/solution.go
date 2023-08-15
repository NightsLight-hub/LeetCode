/*
@Time : 2021/7/9 10:22
@Author : sunxy
@File : solution
@description:

\text{Boyer-Moore}Boyer-Moore  算法
*/
package Find_Majority_Element_LCCI

import "math"

func majorityElement(nums []int) int {
	candidate := math.MaxInt32
	c := 0
	for _, n := range nums {
		if c == 0 {
			candidate = n
			c++
		} else {
			if candidate == n {
				c++
			} else {
				c--
			}
		}
	}
	c = 0
	for _, n := range nums {
		if candidate == n {
			c++
		}
	}
	if c >= len(nums)/2+1 {
		return candidate
	} else {
		return -1
	}
}
