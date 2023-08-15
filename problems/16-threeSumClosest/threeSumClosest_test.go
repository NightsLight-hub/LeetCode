/*
Package _16_threeSumClosest
@Time : 2023/7/10 11:14
@Author : sunxy
@File : threeSumClosest_test.go
@description:
*/
package _16_threeSumClosest

import "testing"

func Test_threeSumClosest(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{
				nums:   []int{-1, 2, 1, -4},
				target: 1,
			},
			want: 2,
		},
		{
			name: "t2",
			args: args{
				nums:   []int{0, 0, 0},
				target: 1,
			},
			want: 0,
		},
		{
			name: "t3",
			args: args{
				nums:   []int{4, 0, 5, -5, 3, 3, 0, -4, -5},
				target: -2,
			},
			want: -2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSumClosest(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("threeSumClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
