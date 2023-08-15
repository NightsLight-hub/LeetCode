/*
Package _61_repeatedNum
@Time : 2022/5/21 10:16
@Author : sunxy
@File : solution_test.go
@description:
*/
package _61_repeatedNum

import "testing"

func Test_repeatedNTimes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				nums: []int{1, 2, 3, 3},
			},
			want: 3,
		},
		{
			name: "test2",
			args: args{
				nums: []int{2, 1, 2, 5, 3, 2},
			},
			want: 2,
		},
		{
			name: "test2",
			args: args{
				nums: []int{5, 1, 5, 2, 5, 3, 5, 4},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatedNTimes(tt.args.nums); got != tt.want {
				t.Errorf("repeatedNTimes() = %v, want %v", got, tt.want)
			}
		})
	}
}
