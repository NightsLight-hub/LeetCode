/*
Package robot_in_a_grid
@Time : 2024/7/11 11:24
@Author : sunxy
@File : robotinagrid_test.go
@description:
*/
package robot_in_a_grid

import (
	"testing"
)

func Test_pathWithObstacles(t *testing.T) {
	// grid := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	// res := pathWithObstacles(grid)
	// for i, _ := range res {
	// 	t.Log(res[i])
	// }

	grid := [][]int{{0}}
	res := pathWithObstacles(grid)
	for i, _ := range res {
		t.Log(res[i])
	}
}
