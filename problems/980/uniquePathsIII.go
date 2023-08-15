/*
Package _80
@Time : 2023/8/4 10:01
@Author : sunxy
@File : uniquePathsIII
@description:
*/
package _980

func uniquePathsIII(grid [][]int) int {
	length, width := len(grid), len(grid[0])
	var startX, startY, nEmpty int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				startX, startY = i, j
			} else if grid[i][j] == 0 {
				nEmpty++
			}
		}
	}
	var dfs func(x, y int)
	res := 0
	walkDistance := 0
	dfs = func(x, y int) {
		if x < 0 || x >= length || y < 0 || y >= width || grid[x][y] == -1 {
			return
		}
		oldV := grid[x][y]
		if grid[x][y] == 2 {
			// 走过了所有empty格子
			if walkDistance == nEmpty {
				res++
			}
			return
		}
		if grid[x][y] == 0 {
			grid[x][y] = -1
			walkDistance++
		} else if grid[x][y] == 1 {
			grid[x][y] = -1
		}
		dfs(x+1, y)
		dfs(x-1, y)
		dfs(x, y+1)
		dfs(x, y-1)
		// 还原格子的值
		if oldV == 0 {
			walkDistance--
		}
		grid[x][y] = oldV

	}

	dfs(startX, startY)
	return res
}
