/*
Package _464_can_i_win
@Time : 2022/5/22 10:39
@Author : sunxy
@File : caniwin_test.go
@description:
*/
package _464_can_i_win

import "testing"

func Test_canIWin(t *testing.T) {
	type args struct {
		maxChoosableInteger int
		desiredTotal        int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{
				maxChoosableInteger: 10,
				desiredTotal:        11,
			},
			want: false,
		},
		{
			name: "t2",
			args: args{
				maxChoosableInteger: 10,
				desiredTotal:        0,
			},
			want: true,
		},
		{
			name: "t3",
			args: args{
				maxChoosableInteger: 10,
				desiredTotal:        1,
			},
			want: true,
		},
		{
			name: "t4",
			args: args{
				maxChoosableInteger: 10,
				desiredTotal:        40,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canIWin(tt.args.maxChoosableInteger, tt.args.desiredTotal); got != tt.want {
				t.Errorf("canIWin() = %v, want %v", got, tt.want)
			}
		})
	}
}
