/*
Package _17_11_find_closest_LCCI
@Time : 2022/5/27 8:39
@Author : sunxy
@File : solution_test.go
@description:
*/
package _17_11_find_closest_LCCI

import "testing"

func Test_findClosest(t *testing.T) {
	type args struct {
		words []string
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "t1",
			args: args{
				words: []string{"I", "am", "a", "student", "from", "a", "university", "in", "a", "city"},
				word1: "a",
				word2: "student",
			},
			want: 1,
		},
		{
			name: "t2",
			args: args{
				words: []string{"I", "am", "a", "student", "from", "a", "university", "in", "a", "city"},
				word1: "I",
				word2: "university",
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findClosest(tt.args.words, tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("findClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
