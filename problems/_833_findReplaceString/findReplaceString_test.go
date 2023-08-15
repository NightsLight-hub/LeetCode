/*
Package _833_findReplaceString
@Time : 2023/8/15 9:10
@Author : sunxy
@File : findReplaceString_test.go
@description:
*/
package _833_findReplaceString

import "testing"

func Test_findReplaceString(t *testing.T) {
	type args struct {
		s       string
		indices []int
		sources []string
		targets []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "t1",
			args: args{
				s:       "abcd",
				indices: []int{0, 2},
				sources: []string{"a", "cd"},
				targets: []string{"eee", "ffff"},
			},
		},
		{
			name: "t2",
			args: args{
				s:       "abcd",
				indices: []int{0, 2},
				sources: []string{"ab", "ec"},
				targets: []string{"eee", "ffff"},
			},
		},
	}
	for _, tt := range tests {
		got := findReplaceString(tt.args.s, tt.args.indices, tt.args.sources, tt.args.targets)
		t.Logf("%#v", got)
	}
}
