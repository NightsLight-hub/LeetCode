/*
Package _56_merge
@Time : 2023/8/27 13:41
@Author : sunxy
@File : merge_test.go
@description:
*/
package _56_merge

import (
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := merge(tt.args.intervals)
			t.Logf("%#v", got)
		})
	}
}
