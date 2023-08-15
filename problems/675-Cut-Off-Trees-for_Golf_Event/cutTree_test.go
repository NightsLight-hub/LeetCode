/*
Package _75_Cut_Off_Trees_for_Golf_Event
@Time : 2022/5/23 21:17
@Author : sunxy
@File : cutTree_test.go
@description:
*/
package _75_Cut_Off_Trees_for_Golf_Event

import "testing"

func Test_cutOffTree(t *testing.T) {
	type args struct {
		forest [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{forest: [][]int{{1, 2, 3}, {0, 0, 4}, {7, 6, 5}}},
			want: 6,
		},
		{
			name: "t2",
			args: args{forest: [][]int{{1, 2, 3}, {0, 0, 0}, {7, 6, 5}}},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cutOffTree(tt.args.forest); got != tt.want {
				t.Errorf("cutOffTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
