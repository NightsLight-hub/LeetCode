/*
Package _490_isCircularSentence
@Time : 2023/6/30 9:32
@Author : sunxy
@File : isCircularSentence_test.go
@description:
*/
package _490_isCircularSentence

import "testing"

func Test_isCircularSentence(t *testing.T) {
	type args struct {
		sentence string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{sentence: "leetcode exercises sound delightful"},
			want: true,
		},
		{
			name: "t2",
			args: args{sentence: "eetcode"},
			want: true,
		},
		{
			name: "t1",
			args: args{sentence: "Leetcode is cool"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isCircularSentence(tt.args.sentence); got != tt.want {
				t.Errorf("isCircularSentence() = %v, want %v", got, tt.want)
			}
		})
	}
}
