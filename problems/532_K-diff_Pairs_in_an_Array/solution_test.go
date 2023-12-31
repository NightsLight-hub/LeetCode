/*
Package _32_K_diff_Pairs_in_an_Array
@Time : 2022/6/17 17:52
@Author : sunxy
@File : solution_test.go
@description:
*/
package _32_K_diff_Pairs_in_an_Array

import "testing"

func Test_findPairs(t *testing.T) {
	type args struct {
		nums []int
		k    int
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
				nums: []int{3, 1, 4, 1, 5},
				k:    2,
			},
			want: 2,
		},
		{
			name: "t2",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
				k:    1,
			},
			want: 4,
		},
		{
			name: "t3",
			args: args{
				nums: []int{1, 3, 1, 5, 4},
				k:    0,
			},
			want: 1,
		},
		{
			name: "t4",
			args: args{
				nums: []int{1, 2, 4, 4, 3, 3, 0, 9, 2, 3},
				k:    3,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPairs(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
