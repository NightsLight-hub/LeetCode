// Package strange_printer /*
/**
https://leetcode-cn.com/problems/strange-printer/

*/
package strange_printer

import "math"

func strangePrinter(s string) int {
	n := len(s)
	f := make([][]int, n)
	for i := 0; i < n; i++ {
		f[i] = make([]int, n)
	}
	for i := n - 1; i >= 0; i-- {
		f[i][i] = 1
		for j := i + 1; j < n; j++ {
			f[i][j] = math.MaxInt32
			if s[i] == s[j] {
				f[i][j] = f[i][j-1]
			} else {
				for k := i; k+1 <= j; k++ {
					f[i][j] = min(f[i][j], f[i][k]+f[k+1][j])
				}
			}
		}
	}
	return f[0][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
