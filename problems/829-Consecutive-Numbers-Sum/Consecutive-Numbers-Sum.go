/*
Package _29__Consecutive_Numbers_Sum
@Time : 2022/6/3 11:32
@Author : sunxy
@File : Consecutive-Numbers-Sum
@description:
*/
package _829__Consecutive_Numbers_Sum

func isKConsecutive(n, k int) bool {
	if k%2 == 1 {
		return n%k == 0
	}
	return n%k != 0 && 2*n%k == 0
}

func consecutiveNumbersSum(n int) (ans int) {
	for k := 1; k*(k+1) <= n*2; k++ {
		if isKConsecutive(n, k) {
			ans++
		}
	}
	return
}

func consecutiveNumbersSum(n int) int {
	mid := n/2 + 1
	res := 1
	for i := 1; i < mid; i++ {
		for j := i + 1; j < mid+1; j++ {
			var ij = consecutiveSum(i, j)
			if ij == n {
				res += 1
			}
		}
	}
	return res
}

func consecutiveSum(l, r int) int {
	if r < l {
		return 0
	}
	n := r - l + 1
	return n * (l + r) / 2
}
