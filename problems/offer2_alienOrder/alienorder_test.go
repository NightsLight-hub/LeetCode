/*
Package offer2_alienOrder
@Time : 2022/5/31 10:08
@Author : sunxy
@File : alienorder_test.go
@description:
*/
package offer2_alienOrder

import "testing"

func Test_alienOrder(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "t1",
			args: args{words: []string{"wrt", "wrf", "er", "ett", "rftt"}},
			want: "wertf",
		},
		{
			name: "t2",
			args: args{words: []string{"z", "x"}},
			want: "zx",
		},
		{
			name: "t3",
			args: args{words: []string{"z", "z"}},
			want: "z",
		},
		{
			name: "t4",
			args: args{words: []string{"zy", "zx"}},
			want: "yxz",
		},
		{
			name: "t5",
			args: args{words: []string{"wrt", "wrtkj"}},
			want: "wrtkj",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := alienOrder(tt.args.words); got != tt.want {
				t.Errorf("alienOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
