/*
Package _1021_Remove_Outermost_Parentheses
@Time : 2022/5/28 9:46
@Author : sunxy
@File : remove_outermost_parentheses_test.go
@description:
*/
package _1021_Remove_Outermost_Parentheses

import "testing"

func Test_removeOuterParentheses(t *testing.T) {
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
			args: args{s: "(()())(())"},
			want: "()()()",
		},
		{
			name: "t2",
			args: args{s: "(()())(())(()(()))"},
			want: "()()()()(())",
		},
		{
			name: "t3",
			args: args{s: "()()"},
			want: "",
		},
		{
			name: "t4",
			args: args{s: "((()()))"},
			want: "(()())",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeOuterParentheses(tt.args.s); got != tt.want {
				t.Errorf("removeOuterParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
