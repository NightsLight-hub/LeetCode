/*
Package _1089_Duplicate_Zeros
@Time : 2022/6/17 17:10
@Author : sunxy
@File : solution
@description:
*/
package _1089_Duplicate_Zeros

func duplicateZeros(arr []int) {
	res := make([]int, 0)
	i := 0
	for len(res) < len(arr) {
		res = append(res, arr[i])
		if arr[i] == 0 {
			res = append(res, 0)
		}
		i++
	}
	for i := range arr {
		arr[i] = res[i]
	}
}
