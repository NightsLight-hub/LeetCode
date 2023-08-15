/*
Package _890_Find_and_Replace_Pattern
@Time : 2022/6/12 9:18
@Author : sunxy
@File : solution_test.go
@description:
*/
package _890_Find_and_Replace_Pattern

import (
	"reflect"
	"testing"
)

func Test_findAndReplacePattern(t *testing.T) {
	type args struct {
		words   []string
		pattern string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{
				words:   []string{"abc", "deq", "mee", "aqq", "dkd", "ccc"},
				pattern: "abb",
			},
			want: []string{"mee", "aqq"},
		},
		{
			name: "t2",
			args: args{
				words:   []string{"a", "b", "c"},
				pattern: "a",
			},
			want: []string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAndReplacePattern(tt.args.words, tt.args.pattern); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAndReplacePattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
