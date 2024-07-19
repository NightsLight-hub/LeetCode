/*
Package dynamic_programing
@Time : 2024/7/11 10:43
@Author : sunxy
@File : upstair
@description:

三步问题。有个小孩正在上楼梯，楼梯有n阶台阶，小孩一次可以上1阶、2阶或3阶。
实现一种方法，计算小孩有多少种上楼梯的方式。（提示：#152，#178，#217，#237，#262，#359）

假设有n阶台阶，小孩一次可以上1阶、2阶或3阶，那么小孩的上楼梯方式有f(n)种，那么
f(n) = f(n-1) + f(n-2) + f(n-3).
f(0) = 1
f(1) = 1
f(2) = 2
f(3) = 4
f(4) = 7
*/
package dynamic_programing

var dpCache []int

func dp(n int) int {
	if dpCache[n] != 0 {
		return dpCache[n]
	} else {
		res := (dp(n-1) + dp(n-2) + dp(n-3)) % 1000000007
		dpCache[n] = res
		return res
	}
}

func waysToStep(n int) int {
	dpCache = make([]int, n+3)
	dpCache[0] = 1
	dpCache[1] = 1
	dpCache[2] = 2
	return dp(n)
}
