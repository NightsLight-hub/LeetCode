/*
@Time : 2021/7/1 13:13
@Author : sunxy
@File : solution_test.go
@description:
*/
package lcp_7_send_message

import "testing"

func Test_numWays(t *testing.T) {
	n := 3
	relation := [][]int{{0, 2}, {2, 1}}
	k := 2
	t.Log(numWays(n, relation, k))
}
