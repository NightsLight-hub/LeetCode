/*
Package _2544_alternateDigitSum
@Time : 2023/7/12 15:17
@Author : sunxy
@File : alternateDigitSum_test.go
@description:
*/
package _2544_alternateDigitSum

import "testing"

func Test_alternateDigitSum(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{n: 521},
			want: 4,
		},
		{
			name: "t1",
			args: args{n: 111},
			want: 1,
		},
		{
			name: "t1",
			args: args{n: 886996},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := alternateDigitSum(tt.args.n); got != tt.want {
				t.Errorf("alternateDigitSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
