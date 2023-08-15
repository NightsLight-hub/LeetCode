/*
@Time : 2021/7/4 7:55
@Author : sunxy
@File : solution
@description:
*/
package _45_Set_Mismatch

// O(n)
func findErrorNums(nums []int) []int {
	nu := make([]int, len(nums)+1)
	wrongNum := 0
	for _, n := range nums {
		if nu[n] == 1 {
			wrongNum = n
		} else {
			nu[n] = 1
		}
	}
	correctNum := 0
	for i, n := range nu {
		if n == 0 {
			correctNum = i
		}
	}
	return []int{wrongNum, correctNum}
}
