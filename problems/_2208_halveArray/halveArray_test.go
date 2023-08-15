/*
Package _2208_halveArray
@Time : 2023/7/25 10:00
@Author : sunxy
@File : halveArray_test.go
@description:
*/
package _2208_halveArray

import "testing"

func Test_halveArray(t *testing.T) {
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
			args: args{nums: []int{5, 19, 8, 1}},
			want: 3,
		},
		{
			name: "t2",
			args: args{nums: []int{3, 8, 20}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := halveArray(tt.args.nums); got != tt.want {
				t.Errorf("halveArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
