/*
Package _849_maxDistToClosest
@Time : 2023/8/22 9:44
@Author : sunxy
@File : maxDistToClosest_test.go
@description:
*/
package _849_maxDistToClosest

import "testing"

func Test_maxDistToClosest(t *testing.T) {
	type args struct {
		seats []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{seats: []int{1, 0, 0, 0, 1, 0, 1}},
			want: 2,
		},
		{
			name: "t2",
			args: args{seats: []int{1, 0, 0, 0}},
			want: 3,
		},
		{
			name: "t3",
			args: args{seats: []int{1, 1, 0, 1, 1}},
			want: 1,
		},
		{
			name: "t4",
			args: args{seats: []int{0, 1, 0, 1, 0}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDistToClosest(tt.args.seats); got != tt.want {
				t.Errorf("maxDistToClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
