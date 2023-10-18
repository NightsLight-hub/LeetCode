/*
Package _2530_maxKelements
@Time : 2023/10/18 9:49
@Author : sunxy
@File : maxkelements_test.go
@description:
*/
package _2530_maxKelements

import "testing"

func Test_maxKelements(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{
				nums: []int{10, 10, 10, 10, 10},
				k:    5,
			},
			want: 50,
		},
		{
			name: "t2",
			args: args{
				nums: []int{1, 10, 3, 3, 3},
				k:    3,
			},
			want: 17,
		},
		// {
		// 	name: "t1",
		// 	args: args{
		// 		nums: []int{10, 10, 10, 10, 10},
		// 		k:    5,
		// 	},
		// 	want: 50,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxKelements(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maxKelements() = %v, want %v", got, tt.want)
			}
		})
	}
}
