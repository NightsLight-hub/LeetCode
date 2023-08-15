/*
Package _829__Consecutive_Numbers_Sum
@Time : 2022/6/3 11:36
@Author : sunxy
@File : Consecutive-Numbers-Sum_test.go
@description:
*/
package _829__Consecutive_Numbers_Sum

import "testing"

func Test_consecutiveNumbersSum(t *testing.T) {
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
			args: args{n: 5},
			want: 2,
		},
		{
			name: "t2",
			args: args{n: 9},
			want: 3,
		},
		{
			name: "15",
			args: args{n: 15},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := consecutiveNumbersSum(tt.args.n); got != tt.want {
				t.Errorf("consecutiveNumbersSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
