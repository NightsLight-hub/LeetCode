/*
Package _0_1_Fibonacci
@Time : 2022/5/21 19:49
@Author : sunxy
@File : fibonacci
@description:
*/
package _0_1_Fibonacci

// 时间复杂度 O(n)   空间复杂度O(1)
func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	// fib1 is fib[n - 1]    fib2 is  fib[n - 2]
	fib1, fib2 := 1, 0
	for i := 2; i <= n; i++ {
		fib := (fib1 + fib2) % 1000000007
		fib2, fib1 = fib1, fib
	}
	return fib1
}
