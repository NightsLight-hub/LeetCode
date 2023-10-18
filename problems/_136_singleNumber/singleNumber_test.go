/*
Package _136_singleNumber
@Time : 2023/10/14 9:13
@Author : sunxy
@File : singleNumber_test.go
@description:
*/
package _136_singleNumber

import "testing"

func Test_singleNumber(t *testing.T) {
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
			args: args{nums: []int{2, 2, 1}},
			want: 1,
		},
		{
			name: "t2",
			args: args{nums: []int{4, 1, 2, 1, 2}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber(tt.args.nums); got != tt.want {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
