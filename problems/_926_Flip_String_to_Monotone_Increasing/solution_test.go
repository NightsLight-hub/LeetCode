/*
Package _926_Flip_String_to_Monotone_Increasing
@Time : 2022/6/12 9:41
@Author : sunxy
@File : solution_test.go
@description:
*/
package _926_Flip_String_to_Monotone_Increasing

import "testing"

func Test_minFlipsMonoIncr(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{s: "00110"},
			want: 1,
		},
		{
			name: "t2",
			args: args{s: "010110"},
			want: 2,
		},
		{
			name: "t3",
			args: args{s: "0"},
			want: 0,
		},
		{
			name: "t3",
			args: args{s: "01001110110"},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minFlipsMonoIncr(tt.args.s); got != tt.want {
				t.Errorf("minFlipsMonoIncr() = %v, want %v", got, tt.want)
			}
		})
	}
}
