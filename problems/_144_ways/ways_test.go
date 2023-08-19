/*
Package _144_ways
@Time : 2023/8/17 9:30
@Author : sunxy
@File : ways_test.go
@description:
*/
package _144_ways

import "testing"

func Test_ways(t *testing.T) {
	type args struct {
		pizza []string
		k     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{
				pizza: []string{"A..", "AAA", "..."},
				k:     3,
			},
			want: 3,
		},
		{
			name: "t2",
			args: args{
				pizza: []string{"A..", "AA.", "..."},
				k:     3,
			},
			want: 1,
		},
		{
			name: "t3",
			args: args{
				pizza: []string{"A..", "A..", "..."},
				k:     1,
			},
			want: 1,
		},
		{
			name: "t3",
			args: args{
				pizza: []string{".A..A", "A.A..", "A.AA.", "AAAA.", "A.AA."},
				k:     5,
			},
			want: 153,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ways(tt.args.pizza, tt.args.k); got != tt.want {
				t.Errorf("ways() = %v, want %v", got, tt.want)
			}
		})
	}
}
