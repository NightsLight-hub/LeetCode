/*
Package _240_tilingRectangle
@Time : 2023/6/8 10:11
@Author : sunxy
@File : solution
@description:
*/
package _240_tilingRectangle

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func tilingRectangle(n int, m int) int {
	ans := max(n, m)
	rect := make([][]bool, m)
	for i := 0; i < m; i++ {
		rect[i] = make([]bool, n)
	}

	isAvailable := func(x, y, k int) bool {
		for i := 0; i < k; i++ {
			for j := 0; j < k; j++ {
				if rect[x+i][y+j] {
					return false
				}
			}
		}
		return true
	}

	fillUp := func(x, y, k int, val bool) {
		for i := 0; i < k; i++ {
			for j := 0; j < k; j++ {
				rect[x+i][y+j] = val
			}
		}
	}
	var dfs func(int, int, int)
	dfs = func(x int, y int, count int) {
		if count >= ans {
			return
		}
		// 填满了
		if y >= n {
			ans = count
			return
		}
		if x >= m {
			dfs(0, y+1, count)
			return
		}
		// x,y被填充了
		if rect[x][y] {
			dfs(x+1, y, count)
			return
		}
		for l := min(m-x, n-y); isAvailable(x, y, l) && l >= 1; l-- {
			fillUp(x, y, l, true)
			dfs(x+l, y, count+1)
			fillUp(x, y, l, false)
		}
	}
	dfs(0, 0, 0)
	return ans
}
