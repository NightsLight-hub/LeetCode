/*
Package _1282_group_people
@Time : 2022/8/12 9:08
@Author : sunxy
@File : solution_test.go
@description:
*/
package _1282_group_people

import (
	"reflect"
	"testing"
)

func Test_groupThePeople(t *testing.T) {
	type args struct {
		groupSizes []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "t1",
			args: args{groupSizes: []int{3, 3, 3, 3, 3, 1, 3}},
			want: [][]int{{5}, {0, 1, 2}, {3, 4, 6}},
		},
		{
			name: "t2",
			args: args{groupSizes: []int{2, 1, 3, 3, 3, 2}},
			want: [][]int{{1}, {0, 5}, {2, 3, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupThePeople(tt.args.groupSizes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupThePeople() = %v, want %v", got, tt.want)
			}
		})
	}
}
