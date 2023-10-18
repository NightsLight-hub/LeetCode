/*
Package _2560_minCapability
@Time : 2023/9/19 9:57
@Author : sunxy
@File : minCapability_test.go
@description:
*/
package _2560_minCapability

import "testing"

func Test_minCapability(t *testing.T) {
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
				nums: []int{2, 7, 9, 3, 1},
				k:    2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCapability(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("minCapability() = %v, want %v", got, tt.want)
			}
		})
	}
}
