/*
Package _2511_captureForts
@Time : 2023/9/2 9:27
@Author : sunxy
@File : captureForts_test.go
@description:
*/
package _2511_captureForts

import "testing"

func Test_captureForts(t *testing.T) {
	type args struct {
		forts []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{forts: []int{1, 0, 0, -1, 0, 0, 0, 0, 1}},
			want: 4,
		},
		{
			name: "t2",
			args: args{forts: []int{0, 0, 1, -1}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := captureForts(tt.args.forts); got != tt.want {
				t.Errorf("captureForts() = %v, want %v", got, tt.want)
			}
		})
	}
}
