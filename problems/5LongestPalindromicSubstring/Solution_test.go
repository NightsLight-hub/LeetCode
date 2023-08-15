/*
Package _LongestPalindromicSubstring
@Time : 2023/6/5 11:20
@Author : sunxy
@File : Solution_test.go
@description:
*/
package _LongestPalindromicSubstring

import (
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
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
			args: args{s: "bb"},
			want: "bb",
		},
		{
			name: "t2",
			args: args{s: "bacab"},
			want: "bacab",
		},
		{
			name: "t2",
			args: args{s: "ddbacabddd"},
			want: "ddbacabdd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
