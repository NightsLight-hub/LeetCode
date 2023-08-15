/*
Package _2569_handleQuery
@Time : 2023/7/26 9:58
@Author : sunxy
@File : handleQuery_test.go
@description:
*/
package _2569_handleQuery

import (
	"reflect"
	"testing"
)

func Test_handleQuery(t *testing.T) {
	type args struct {
		nums1   []int
		nums2   []int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{
				nums1:   []int{1, 0, 1},
				nums2:   []int{0, 0, 0},
				queries: [][]int{{1, 1, 1}, {2, 1, 0}, {3, 0, 0}},
			},
			want: []int64{3},
		},
		// TODO: Add test cases.
		{
			name: "t2",
			args: args{
				nums1:   []int{1},
				nums2:   []int{5},
				queries: [][]int{{2, 0, 0}, {3, 0, 0}},
			},
			want: []int64{5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handleQuery(tt.args.nums1, tt.args.nums2, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handleQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
