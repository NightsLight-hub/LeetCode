/*
Package _300_LengthOfLis
@Time : 2023/12/4 10:29
@Author : sunxy
@File : lenghthOfLis_test.go
@description:
*/
package _300_LengthOfLis

import "testing"

func Test_lengthOfLIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{nums: []int{10, 9, 2, 5, 3, 7, 101, 18}},
			want: 4,
		},
		{
			name: "t2",
			args: args{nums: []int{10}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS(tt.args.nums); got != tt.want {
				t.Errorf("lengthOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
