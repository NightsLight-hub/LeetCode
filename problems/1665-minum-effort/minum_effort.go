/*
@Time : 2021/6/1 20:18
@Author : sunxy
@File : minum_effort
@description:
*/
package _665_minum_effort

import "sort"

/*
*
https://leetcode-cn.com/problems/minimum-initial-energy-to-finish-tasks/

[[1,2],[2,4],[4,8]]

r >4
r > 2 + 2 || r>1+4

[a,b]  [c,d]
r > b && r>d

r > a +d || r > b + c

处理的顺序应该先处理性价比最高的，也就是门槛与实际消耗 差值最大的
*/
func sortByDiff(tasks [][]int) {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i][1]-tasks[i][0] > tasks[j][1]-tasks[j][0]
	})
}

func minimumEffort(tasks [][]int) int {
	l := len(tasks)
	if l == 1 {
		a, _ := maxAndMin(tasks[0][0], tasks[0][1])
		return a
	}
	sortByDiff(tasks)

	minumEffort, remnant := getMinumEffortAndRemnant(tasks[0], tasks[1])
	tuple := []int{minumEffort - remnant, minumEffort}
	for i := 2; i < l; i++ {
		minumEffort, remnant = getMinumEffortAndRemnant(tasks[i], tuple)
		tuple = []int{minumEffort - remnant, minumEffort}
	}
	return minumEffort
}
func getMinumEffortAndRemnant(task1, task2 []int) (int, int) {
	r, _ := maxAndMin(task1[1], task2[1])
	_, t := maxAndMin(task1[0]+task2[1], task2[0]+task1[1])
	minumEffort, _ := maxAndMin(r, t)
	return minumEffort, minumEffort - task1[0] - task2[0]
}
func maxAndMin(a, b int) (int, int) {
	if a > b {
		return a, b
	} else {
		return b, a
	}
}
