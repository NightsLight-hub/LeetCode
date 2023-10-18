/*
Package _652_sumOfMultiples
@Time : 2023/10/17 10:11
@Author : sunxy
@File : sumOfMultiples
@description:
*/
package _652_sumOfMultiples

func sumOfMultiples(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		if i%7 == 0 {
			res += i
		} else if i%5 == 0 {
			res += i
		} else if i%3 == 0 {
			res += i
		}
	}
	return res
}
