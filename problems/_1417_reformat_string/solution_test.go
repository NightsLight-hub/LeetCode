/*
Package _1417_reformat_string
@Time : 2022/8/11 9:24
@Author : sunxy
@File : solution_test.go
@description:
*/
package _1417_reformat_string

import "testing"

func Test_reformat(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "t1",
			args: args{s: "a0b1c2"},
			want: "0a1b2c",
		},
		{
			name: "t2",
			args: args{s: "leetcode"},
			want: "",
		},
		{
			name: "t3",
			args: args{s: "covid2019"},
			want: "c2o0v1i9d",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reformat(tt.args.s); got != tt.want {
				t.Errorf("reformat() = %v, want %v", got, tt.want)
			}
		})
	}
}
