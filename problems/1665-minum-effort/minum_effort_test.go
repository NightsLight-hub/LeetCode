/*
@Time : 2021/6/1 21:01
@Author : sunxy
@File : minum_effort_test.go
@description:
*/
package _665_minum_effort

import "testing"

func Test_minimumEffort(t *testing.T) {
	tasks := [][]int{{3, 7}, {3, 8}, {3, 9}, {4, 10}, {5, 11}, {6, 13}, {3, 7}, {3, 8}, {3, 9}, {4, 10}, {5, 11}, {6, 13}, {3, 7}, {3, 8}, {3, 9}, {4, 10}, {5, 11}, {6, 13}}
	t.Logf("%d", minimumEffort(tasks))
}
