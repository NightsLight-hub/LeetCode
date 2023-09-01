/*
Package _2240_waysToBuyPensPencils
@Time : 2023/9/1 17:45
@Author : sunxy
@File : waysToBuyPensPencils_test.go
@description:
*/
package _2240_waysToBuyPensPencils

import "testing"

func Test_waysToBuyPensPencils(t *testing.T) {
	type args struct {
		total int
		cost1 int
		cost2 int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{
				total: 20,
				cost1: 10,
				cost2: 5,
			},
			want: 9,
		},
		{
			name: "t2",
			args: args{
				total: 5,
				cost1: 10,
				cost2: 10,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := waysToBuyPensPencils(tt.args.total, tt.args.cost1, tt.args.cost2); got != tt.want {
				t.Errorf("waysToBuyPensPencils() = %v, want %v", got, tt.want)
			}
		})
	}
}
