/*
Package robot_in_a_grid
@Time : 2024/7/11 11:03
@Author : sunxy
@File : robotinagrid
@description:

设想有个机器人坐在一个网格的左上角，网格 r 行 c 列。机器人只能向下或向右移动，
但不能走到一些被禁止的网格（有障碍物）。设计一种算法，寻找机器人从左上角移动到右下角的路径。
*/
package robot_in_a_grid

var path = make([][]int, 0)

func pathWithObstacles(obstacleGrid [][]int) [][]int {
	walk([]int{0, 0}, obstacleGrid)
	return path
}

func walk(currentPos []int, obstacleGrid [][]int) bool {
	if currentPos[0] >= len(obstacleGrid) || currentPos[1] >= len(obstacleGrid[0]) {
		return false
	}
	if obstacleGrid[currentPos[0]][currentPos[1]] == 1 {
		return false
	}
	path = append(path, currentPos)
	// 判断是否终点
	if currentPos[0] == len(obstacleGrid)-1 && currentPos[1] == len(obstacleGrid[0])-1 {
		return true
	}
	// right
	newPos := []int{currentPos[0], currentPos[1] + 1}
	if !walk(newPos, obstacleGrid) {
		newPos = []int{currentPos[0] + 1, currentPos[1]}
		if !walk(newPos, obstacleGrid) {
			path = path[:len(path)-1]
			return false
		}
	}
	return true
}
