/*
Package _240_tilingRectangle
@Time : 2023/6/8 14:19
@Author : sunxy
@File : solution_test.go
@description:
*/
package _240_tilingRectangle

import "testing"

func Test_tilingRectangle(t *testing.T) {
	type args struct {
		n int
		m int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "t1",
			args: args{
				n: 2,
				m: 3,
			},
			want: 3,
		},
		{
			name: "t2",
			args: args{
				n: 5,
				m: 8,
			},
			want: 5,
		},
		{
			name: "t3",
			args: args{
				n: 11,
				m: 13,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tilingRectangle(tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("tilingRectangle() = %v, want %v", got, tt.want)
			}
		})
	}
}
