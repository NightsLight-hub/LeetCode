/*
Package _636_ExclusiveTimeofFunctions
@Time : 2022/8/7 9:33
@Author : sunxy
@File : solution_test.go
@description:
*/
package _636_ExclusiveTimeofFunctions

import (
	"reflect"
	"testing"
)

func Test_exclusiveTime(t *testing.T) {
	type args struct {
		n    int
		logs []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "t0",
			args: args{
				n:    2,
				logs: []string{"0:start:0", "1:start:2", "1:end:5", "0:end:6"},
			},
			want: []int{3, 4},
		},
		{
			name: "t1",
			args: args{
				n:    2,
				logs: []string{"0:start:0", "0:start:2", "0:end:5", "0:start:6", "0:end:6", "0:end:7"},
			},
			want: []int{8},
		},
		{
			name: "t2",
			args: args{
				n:    2,
				logs: []string{"0:start:0", "0:start:2", "0:end:5", "1:start:6", "1:end:6", "0:end:7"},
			},
			want: []int{7, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exclusiveTime(tt.args.n, tt.args.logs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("exclusiveTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
