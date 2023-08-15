/*
Package _156_solution
@Time : 2023/6/3 7:48
@Author : sunxy
@File : solution_test.go
@description:
*/
package _156_solution

import "testing"

func Test_maxRepOpt1(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{text: "aaacaaabba"},
			want: 7,
		},
		{
			name: "test2",
			args: args{text: "aaacaaaba"},
			want: 7,
		},
		{
			name: "test3",
			args: args{text: "aaacdaaaba"},
			want: 5,
		},
		{
			name: "test4",
			args: args{text: "aaacd"},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxRepOpt1(tt.args.text); got != tt.want {
				t.Errorf("maxRepOpt1() = %v, want %v", got, tt.want)
			}
		})
	}
}
