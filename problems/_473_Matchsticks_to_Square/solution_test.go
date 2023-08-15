/*
Package _473_Matchsticks_to_Square
@Time : 2022/6/1 18:41
@Author : sunxy
@File : solution_test.go
@description:
*/
package _473_Matchsticks_to_Square

import "testing"

func Test_makesquare(t *testing.T) {
	type args struct {
		matchsticks []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{matchsticks: []int{1, 1, 2, 2, 2}},
			want: true,
		},
		{
			name: "t2",
			args: args{matchsticks: []int{3, 3, 3, 3, 4}},
			want: false,
		},
		{
			name: "t3",
			args: args{matchsticks: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}},
			want: true,
		},
		{
			name: "t4",
			args: args{matchsticks: []int{1, 1, 2, 3, 3, 2, 4}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makesquare(tt.args.matchsticks); got != tt.want {
				t.Errorf("makesquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
