/*
Package _1388_3npizza
@Time : 2023/8/18 10:16
@Author : sunxy
@File : maxSizeSlices_test.go
@description:
*/
package _1388_3npizza

import "testing"

func Test_maxSizeSlices(t *testing.T) {
	type args struct {
		slices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{slices: []int{1, 2, 3, 4, 5, 6}},
			want: 10,
		},
		{
			name: "t2",
			args: args{slices: []int{8, 9, 8, 6, 1, 1}},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSizeSlices(tt.args.slices); got != tt.want {
				t.Errorf("maxSizeSlices() = %v, want %v", got, tt.want)
			}
		})
	}
}
