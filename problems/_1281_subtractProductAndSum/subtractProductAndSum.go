/*
Package _1281_subtractProductAndSum
@Time : 2023/8/9 8:57
@Author : sunxy
@File : subtractProductAndSum
@description:
*/
package _1281_subtractProductAndSum

func subtractProductAndSum(n int) int {
	sum, product := 0, 1
	for n > 0 {
		sum += n % 10
		product *= n % 10
		n /= 10
	}
	return product - sum
}
