/*
Package _1089_Duplicate_Zeros
@Time : 2022/6/17 17:10
@Author : sunxy
@File : solution_test.go
@description:
*/
package _1089_Duplicate_Zeros

import "testing"

func Test_duplicateZeros(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{arr: []int{1, 0, 2, 3, 0, 4, 5, 0}},
			want: []int{1, 0, 0, 2, 3, 0, 0, 4},
		},
		{
			name: "t2",
			args: args{arr: []int{1, 2, 3}},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			duplicateZeros(tt.args.arr)
			if !check(tt.args.arr, tt.want) {
				t.Fail()
			}
		})
	}
}

func check(a1, a2 []int) bool {
	if len(a1) != len(a2) {
		return false
	}
	for i, _ := range a1 {
		if a1[i] != a2[i] {
			return false
		}
	}
	return true
}
