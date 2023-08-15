/*
Package _485_pivotInteger
@Time : 2023/6/26 9:43
@Author : sunxy
@File : pivotInteger_test.go
@description:
*/
package _485_pivotInteger

import "testing"

func Test_pivotInteger(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "t1",
			args: args{n: 8},
			want: 6,
		},
		{
			name: "t2",
			args: args{n: 1},
			want: 1,
		},
		{
			name: "t3",
			args: args{n: 4},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pivotInteger(tt.args.n); got != tt.want {
				t.Errorf("pivotInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}
