/*
@Time : 2021/6/3 9:27
@Author : sunxy
@File : number_of_closed_islands
@description:
有一个二维矩阵 grid ，每个位置要么是陆地（记号为 0 ）要么是水域（记号为 1 ）。

我们从一块陆地出发，每次可以往上下左右 4 个方向相邻区域走，能走到的所有陆地区域，我们将其称为一座「岛屿」。

如果一座岛屿 完全 由水域包围，即陆地边缘上下左右所有相邻区域都是水域，那么我们将其称为 「封闭岛屿」。

请返回封闭岛屿的数目。



来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-closed-islands
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package _1254_Number_of_Closed_Islands

import "container/list"

const (
	land    = 0
	water   = 1
	visited = 2
)

type result struct {
	value bool
}

var (
	landList   = list.New()
	vertical   = [4]int{-1, 1, 0, 0}
	horizontal = [4]int{0, 0, -1, 1}
)

func closedIsland(grid [][]int) int {
	c := 0
	height := len(grid)
	width := len(grid[0])
	y, x := findLand(grid)
	for y != -1 {
		flag := result{value: true}
		dsf(y, x, grid, &flag, height, width)
		if flag.value {
			c++
		}
		y, x = findLand(grid)
	}
	return c
}

func findLand(grid [][]int) (int, int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == land {
				return i, j
			}
		}
	}
	return -1, -1
}

// y,x  startpoint  coordinate
func dsf(y, x int, grid [][]int, flag *result, gridHeight, gridWidth int) {
	if y == gridHeight || y < 0 || x == gridWidth || x < 0 {
		return
	}
	if grid[y][x] == land {
		grid[y][x] = visited
	} else if grid[y][x] == water || grid[y][x] == visited {
		return
	}
	//判断是否边界
	if y == gridHeight-1 || y == 0 || x == gridWidth-1 || x == 0 {
		flag.value = false
	}

	for i := 0; i < 4; i++ {
		dsf(y+vertical[i], x+horizontal[i], grid, flag, gridHeight, gridWidth)
	}
}
