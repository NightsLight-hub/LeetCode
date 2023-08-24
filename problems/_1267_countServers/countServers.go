/*
Package _1267_countServers
@Time : 2023/8/24 13:59
@Author : sunxy
@File : countServers
@description:
*/
package _1267_countServers

// O(n*m)
func countServers(grid [][]int) int {
	// map[rowIndex]count
	// var rm = make(map[int]int)
	var rm = make([]int, len(grid[0]))
	// map[colIndex]count
	// var cm = make(map[int]int)
	var cm = make([]int, len(grid))
	res := 0
	var servers [][2]int
	for i, row := range grid {
		for j, v := range row {
			if v == 1 {
				rm[j]++
				cm[i]++
				if rm[j] == 1 && cm[i] == 1 {
					// 行和列上只有一个的主机，记录下坐标，其他的不必记录了，肯定可以联通
					servers = append(servers, [2]int{i, j})
				} else {
					res++
				}
			}
		}
	}
	for _, server := range servers {
		if rm[server[1]] > 1 || cm[server[0]] > 1 {
			res++
		}
	}
	return res
}
