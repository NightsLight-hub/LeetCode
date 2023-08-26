/*
Package _228_summaryRanges
@Time : 2023/8/26 9:25
@Author : sunxy
@File : summaryRanges_test.go
@description:
*/
package _228_summaryRanges

import (
	"testing"
)

func Test_summaryRanges(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{
				nums: []int{0, 1, 2, 4, 5, 7},
			},
		},
		{
			name: "t2",
			args: args{
				nums: []int{0, 2, 3, 4, 6, 8, 9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := summaryRanges(tt.args.nums)
			t.Logf("%#v", got)
		})
	}
}
