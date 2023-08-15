/*
Package _375_numTimesAllBlue
@Time : 2023/6/14 17:31
@Author : sunxy
@File : Solution_test.go
@description:
*/
package _375_numTimesAllBlue

import "testing"

func Test_numTimesAllBlue(t *testing.T) {
	type args struct {
		flips []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "t1",
			args: args{flips: []int{3, 2, 4, 1, 5}},
			want: 2,
		},
		{
			name: "t2",
			args: args{flips: []int{4, 1, 2, 3}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numTimesAllBlue(tt.args.flips); got != tt.want {
				t.Errorf("numTimesAllBlue() = %v, want %v", got, tt.want)
			}
		})
	}
}
